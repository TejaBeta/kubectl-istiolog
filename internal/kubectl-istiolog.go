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

package internal

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"istio.io/istio/pkg/kube"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	kubeClient    = newKubeClient
	kubeconfig    string
	configContext string
)

type Level int

const (
	defaultLoggerName  = "level"
	defaultOutputLevel = WarningLevel
)

var allLoggers = []string{
	"admin",
	"aws",
	"assert",
	"backtrace",
	"client",
	"config",
	"connection",
	"conn_handler", // Added through https://github.com/envoyproxy/envoy/pull/8263
	"dubbo",
	"file",
	"filter",
	"forward_proxy",
	"grpc",
	"hc",
	"health_checker",
	"http",
	"http2",
	"hystrix",
	"init",
	"io",
	"jwt",
	"kafka",
	"lua",
	"main",
	"misc",
	"mongo",
	"quic",
	"pool",
	"rbac",
	"redis",
	"router",
	"runtime",
	"stats",
	"secret",
	"tap",
	"testing",
	"thrift",
	"tracing",
	"upstream",
	"udp",
	"wasm",
}

var levelToString = map[Level]string{
	TraceLevel:    "trace",
	DebugLevel:    "debug",
	InfoLevel:     "info",
	WarningLevel:  "warning",
	ErrorLevel:    "error",
	CriticalLevel: "critical",
	OffLevel:      "off",
}

var stringToLevel = map[string]Level{
	"trace":    TraceLevel,
	"debug":    DebugLevel,
	"info":     InfoLevel,
	"warning":  WarningLevel,
	"error":    ErrorLevel,
	"critical": CriticalLevel,
	"off":      OffLevel,
}

const (
	istioContainer = "istio-proxy"
	// OffLevel disables logging
	OffLevel Level = iota
	// CriticalLevel enables critical level logging
	CriticalLevel
	// ErrorLevel enables error level logging
	ErrorLevel
	// WarningLevel enables warning level logging
	WarningLevel
	// InfoLevel enables info level logging
	InfoLevel
	// DebugLevel enables debug level logging
	DebugLevel
	// TraceLevel enables trace level logging
	TraceLevel
)

func (opts *options) isPodExists(podName string) error {
	result, err := opts.clientset.CoreV1().Pods(opts.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, pod := range result.Items {
		if pod.Name == podName {
			return nil
		}
	}

	return fmt.Errorf("%v Pod doesn't exist", podName)
}

func newKubeClientWithRevision(kubeconfig, configContext string, revision string) (kube.CLIClient, error) {
	return kube.NewCLIClient(kube.BuildClientCmd(kubeconfig, configContext), revision)
}

func newKubeClient(kubeconfig, configContext string) (kube.CLIClient, error) {
	return newKubeClientWithRevision(kubeconfig, configContext, "")
}

func setupEnvoyLog(param, pod, namespace string) (string, error) {
	kubeClient, err := kubeClient(kubeconfig, configContext)
	if err != nil {
		return "", fmt.Errorf("failed to create Kubernetes client: %v", err)
	}
	path := "logging"
	if param != "" {
		path = path + "?" + param
	}
	result, err := kubeClient.EnvoyDo(context.TODO(), pod, namespace, "POST", path)
	if err != nil {
		return "", fmt.Errorf("failed to execute command on Envoy: %v", err)
	}
	return string(result), nil
}

func handleLog(logLevel string, pod string, namespace string) error {
	destLoggerLevels := map[string]Level{}

	levels := strings.Split(logLevel, ",")
	for _, ol := range levels {
		if !strings.Contains(ol, ":") && !strings.Contains(ol, "=") {
			level, ok := stringToLevel[ol]
			if ok {
				destLoggerLevels = map[string]Level{
					defaultLoggerName: level,
				}
			} else {
				return fmt.Errorf("unrecognized logging level: %v", ol)
			}
		} else {
			invalidLogName := true
			loggerLevel := regexp.MustCompile(`[:=]`).Split(ol, 2)

			for _, logName := range allLoggers {
				if logName == loggerLevel[0] {
					invalidLogName = false
					break
				}
			}

			if invalidLogName {
				return fmt.Errorf("unrecognized logger name: %v", loggerLevel[0])
			}

			level, ok := stringToLevel[loggerLevel[1]]
			if !ok {
				return fmt.Errorf("unrecognized logging level: %v", loggerLevel[1])
			}
			destLoggerLevels[loggerLevel[0]] = level
		}
	}

	var resp string
	var err error

	if len(destLoggerLevels) == 0 {
		resp, err = setupEnvoyLog("", pod, namespace)
	} else {
		if ll, ok := destLoggerLevels[defaultLoggerName]; ok {
			// update levels of all loggers first
			resp, err = setupEnvoyLog(defaultLoggerName+"="+levelToString[ll], pod, namespace)
			delete(destLoggerLevels, defaultLoggerName)
		}
		for lg, ll := range destLoggerLevels {
			resp, err = setupEnvoyLog(lg+"="+levelToString[ll], pod, namespace)
		}
	}
	if err != nil {
		return err
	}
	fmt.Print(resp)
	return nil
}

func (opts *options) streamLogs(podName string, containerName string) error {
	count := int64(1)
	podLogOptions := corev1.PodLogOptions{
		Container: containerName,
		Follow:    true,
		TailLines: &count,
	}

	req := opts.clientset.CoreV1().Pods(opts.namespace).GetLogs(podName, &podLogOptions)
	stream, err := req.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer stream.Close()
	for {
		buf := make([]byte, 2000)
		numBytes, err := stream.Read(buf)
		if numBytes == 0 {
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		message := string(buf[:numBytes])
		fmt.Print(message)
	}
	return nil
}

func (options *options) KubectlIstioLog(pod string, logLevel string, follow bool) error {
	err := options.isPodExists(pod)
	if err != nil {
		return err
	}

	err = handleLog(logLevel, pod, options.namespace)
	if err != nil {
		return err
	}

	if follow {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-c
			err := handleLog(levelToString[defaultOutputLevel], pod, options.namespace)
			if err != nil {
				fmt.Print(err)
			}
			os.Exit(0)
		}()

		err := options.streamLogs(pod, istioContainer)
		if err != nil {
			return err
		}
	}

	return nil
}
