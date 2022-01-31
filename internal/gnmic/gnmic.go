package gnmic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	gutils "github.com/karimra/gnmic/utils"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/stoewer/go-strcase"
	"github.com/yndd/ndd-runtime/pkg/utils"
	"github.com/yndd/ndd-yang/pkg/yparser"
)

const (
	defaultEncoding = "JSON_IETF"
	defaultTimeout  = 30 * time.Second
	maxMsgSize      = 512 * 1024 * 1024
)

func SubName(s *string) *string {
	split := strings.Split(*s, "/")
	subName := "sub"
	for _, n := range split {
		subName += strcase.UpperCamelCase(n)
	}
	return &subName
}

func CreateSubscriptionRequest(subPaths *[]string) (*gnmi.SubscribeRequest, error) {
	// create subscription
	paths := *subPaths

	gnmiPrefix, err := gutils.CreatePrefix("", "")
	if err != nil {
		return nil, fmt.Errorf("create prefix failed")
	}
	modeVal := gnmi.SubscriptionList_Mode_value[strings.ToUpper("STREAM")]
	qos := &gnmi.QOSMarking{Marking: 21}

	subscriptions := make([]*gnmi.Subscription, len(paths))
	for i, p := range paths {
		gnmiPath, _ := gutils.ParsePath(strings.TrimSpace(p))
		subscriptions[i] = &gnmi.Subscription{Path: gnmiPath}
		switch gnmi.SubscriptionList_Mode(modeVal) {
		case gnmi.SubscriptionList_STREAM:
			mode := gnmi.SubscriptionMode_value[strings.Replace(strings.ToUpper("ON_CHANGE"), "-", "_", -1)]
			subscriptions[i].Mode = gnmi.SubscriptionMode(mode)
		}
	}

	req := &gnmi.SubscribeRequest{
		Request: &gnmi.SubscribeRequest_Subscribe{
			Subscribe: &gnmi.SubscriptionList{
				Prefix:       gnmiPrefix,
				Mode:         gnmi.SubscriptionList_Mode(modeVal),
				Encoding:     46, // "JSON_IETF_CONFIG_ONLY"
				Subscription: subscriptions,
				Qos:          qos,
			},
		},
	}
	return req, nil
}

// CreateGetRequest function creates a gnmi get request
func CreateConfigGetRequest(path []*gnmi.Path, dataType, encoding *string) (*gnmi.GetRequest, error) {
	if encoding == nil {
		encoding = utils.StringPtr(defaultEncoding)
	}
	encodingVal, ok := gnmi.Encoding_value[strings.Replace(strings.ToUpper(defaultEncoding), "-", "_", -1)]
	if !ok {
		return nil, fmt.Errorf("invalid encoding type '%s'", *encoding)
	}
	dti, ok := gnmi.GetRequest_DataType_value[strings.ToUpper(*dataType)]
	if !ok {
		return nil, fmt.Errorf("unknown data type %s", *dataType)
	}
	req := &gnmi.GetRequest{
		UseModels: make([]*gnmi.ModelData, 0),
		Path:      path,
		Encoding:  gnmi.Encoding(encodingVal),
		Type:      gnmi.GetRequest_DataType(dti),
	}

	return req, nil
}

// CreateGetRequest function creates a gnmi get request
func CreateGetRequest(path, dataType, encoding *string) (*gnmi.GetRequest, error) {
	if encoding == nil {
		encoding = utils.StringPtr(defaultEncoding)
	}
	encodingVal, ok := gnmi.Encoding_value[strings.Replace(strings.ToUpper(defaultEncoding), "-", "_", -1)]
	if !ok {
		return nil, fmt.Errorf("invalid encoding type '%s'", *encoding)
	}
	dti, ok := gnmi.GetRequest_DataType_value[strings.ToUpper(*dataType)]
	if !ok {
		return nil, fmt.Errorf("unknown data type %s", *dataType)
	}
	req := &gnmi.GetRequest{
		UseModels: make([]*gnmi.ModelData, 0),
		Path:      make([]*gnmi.Path, 0),
		Encoding:  gnmi.Encoding(encodingVal),
		Type:      gnmi.GetRequest_DataType(dti),
	}
	prefix := ""
	if prefix != "" {
		gnmiPrefix, err := gutils.ParsePath(prefix)
		if err != nil {
			return nil, fmt.Errorf("prefix parse error: %v", err)
		}
		req.Prefix = gnmiPrefix
	}

	gnmiPath, err := gutils.ParsePath(strings.TrimSpace(*path))
	if err != nil {
		return nil, fmt.Errorf("path parse error: %v", err)
	}
	req.Path = append(req.Path, gnmiPath)
	return req, nil
}

// CreateSetRequest function creates a gnmi set request
func CreateSetRequest(path *string, updateBytes []byte) (*gnmi.SetRequest, error) {
	value := new(gnmi.TypedValue)
	value.Value = &gnmi.TypedValue_JsonIetfVal{
		JsonIetfVal: bytes.Trim(updateBytes, " \r\n\t"),
	}

	gnmiPrefix, err := gutils.CreatePrefix("", "")
	if err != nil {
		return nil, fmt.Errorf("prefix parse error: %v", err)
	}

	gnmiPath, err := gutils.ParsePath(strings.TrimSpace(*path))
	if err != nil {
		return nil, fmt.Errorf("path parse error: %v", err)
	}

	req := &gnmi.SetRequest{
		Prefix:  gnmiPrefix,
		Delete:  make([]*gnmi.Path, 0),
		Replace: make([]*gnmi.Update, 0),
		Update:  make([]*gnmi.Update, 0),
	}

	req.Update = append(req.Update, &gnmi.Update{
		Path: gnmiPath,
		Val:  value,
	})

	return req, nil
}

// CreateDeleteRequest function
func CreateDeleteRequest(path *string) (*gnmi.SetRequest, error) {
	gnmiPrefix, err := gutils.CreatePrefix("", "")
	if err != nil {
		return nil, fmt.Errorf("prefix parse error: %v", err)
	}

	gnmiPath, err := gutils.ParsePath(strings.TrimSpace(*path))
	if err != nil {
		return nil, fmt.Errorf("path parse error: %v", err)
	}

	req := &gnmi.SetRequest{
		Prefix:  gnmiPrefix,
		Delete:  make([]*gnmi.Path, 0),
		Replace: make([]*gnmi.Update, 0),
		Update:  make([]*gnmi.Update, 0),
	}

	req.Delete = append(req.Delete, gnmiPath)

	return req, nil
}

// HandleGetResponse handes the response
func HandleGetResponse(response *gnmi.GetResponse) ([]Update, error) {
	for _, notif := range response.GetNotification() {

		updates := make([]Update, 0, len(notif.GetUpdate()))

		for i, upd := range notif.GetUpdate() {
			// Path element processing
			pathElems := make([]string, 0, len(upd.GetPath().GetElem()))
			for _, pElem := range upd.GetPath().GetElem() {
				pathElems = append(pathElems, pElem.GetName())
			}
			var pathElemSplit []string
			var pathElem string
			if len(pathElems) != 0 {
				if len(pathElems) > 1 {
					pathElemSplit = strings.Split(pathElems[len(pathElems)-1], ":")
				} else {
					pathElemSplit = strings.Split(pathElems[0], ":")
				}

				if len(pathElemSplit) > 1 {
					pathElem = pathElemSplit[len(pathElemSplit)-1]
				} else {
					pathElem = pathElemSplit[0]
				}
			} else {
				pathElem = ""
			}

			// Value processing
			value, err := yparser.GetValue(upd.GetVal())
			if err != nil {
				return nil, err
			}
			updates = append(updates,
				Update{
					Path:   yparser.GnmiPath2XPath(upd.GetPath(), true),
					Values: make(map[string]interface{}),
				})
			updates[i].Values[pathElem] = value

		}
		x, err := json.Marshal(updates)
		if err != nil {
			return nil, nil
		}
		sb := strings.Builder{}
		sb.Write(x)

		return updates, nil
	}
	return nil, nil
}

type Update struct {
	Path   string
	Values map[string]interface{} `json:"values,omitempty"`
}
