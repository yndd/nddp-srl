/*
Copyright 2022 NDD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"os"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	ndrv1 "github.com/yndd/ndd-core/apis/dvr/v1"

	srlv1alpha1 "github.com/yndd/nddp-srl/apis/srl/v1alpha1"
	//apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	//+kubebuilder:scaffold:imports
)

var (
	scheme = runtime.NewScheme()
	debug  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "manager",
	Short: "nddo srl intent controller",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")

	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(srlv1alpha1.AddToScheme(scheme))
	utilruntime.Must(ndrv1.AddToScheme(scheme))
	//utilruntime.Must(apiextensionsv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}
