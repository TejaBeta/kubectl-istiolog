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

	internal "github.com/TejaBeta/kubectl-istiolog/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
)

var (
	flagVersion   bool
	flagVerbose   bool
	flagNameSpace string
	flagFollow    bool
	flagLogLevel  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-istiolog [pod] [flags]",
	Short: "A Kubectl plugin to manage and set envoy log levels",

	Run: func(cmd *cobra.Command, args []string) {
		context, err := internal.GetContext()
		if err != nil {
			log.Fatalln(err)
		}
		options, err := internal.GetOpts(context, flagNameSpace)
		if err != nil {
			log.Fatalln(err)
		}
		err = options.KubectlIstioLog(args[0], flagLogLevel, flagFollow)
		if err != nil {
			panic(err)
		}
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
		if flagVerbose {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}
	})
	rootCmd.Flags().BoolVar(&flagVerbose, "verbose", false, "Verbose mode on")
	rootCmd.Flags().StringVarP(&flagNameSpace, "namespace", "n", "default", "Namespace in current context")
	rootCmd.Flags().BoolVarP(&flagFollow, "follow", "f", false, "Specify if the logs should be streamed")
	rootCmd.Flags().StringVarP(&flagLogLevel, "level", "l", "warning", "Comma-separated minimum per-logger level of messages to output")
}
