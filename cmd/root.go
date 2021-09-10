/*
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
package cmd

import (
	"fmt"
	"os"

	"github.com/TejaBeta/kubectl-istiolog/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
)

var (
	flagVersion bool
	flagDebug   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-istiolog",
	Short: "A Kubectl plugin to debug Istio logs",
	Long: `kubectl istiolog is a small plugin to change envoy log
	and also provides a functionality to watch:

Happy Debugging!!!⛵️⛵️⛵️`,

	Run: func(cmd *cobra.Command, args []string) {
		pkg.KubectlIstioLog()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if flagVersion {
			fmt.Println("version:", version)
			fmt.Println("commit:", commit)
		}
		if flagDebug {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}
	})
	rootCmd.Flags().BoolVarP(&flagDebug, "debug", "d", false, "Set debug mode on")
	rootCmd.Flags().BoolVarP(&flagVersion, "version", "v", false, "Get version info")
}
