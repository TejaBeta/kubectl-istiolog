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
	"errors"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"istio.io/istio/pkg/kube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeClient    = newKubeClient
	kubeconfig    string
	configContext string
)

type options struct {
	clientset kubernetes.Interface
	namespace string
}

func getOpts(context *rest.Config, ns string) (*options, error) {
	cs, err := kubernetes.NewForConfig(context)
	if err != nil {
		return nil, err
	}

	return &options{
		clientset: cs,
		namespace: ns,
	}, nil
}

func (opts *options) isPodExists(podName string) bool {
	result, err := opts.clientset.CoreV1().Pods(opts.namespace).List(context.TODO(), metav1.ListOptions{})
	if err == nil {
		for _, pod := range result.Items {
			if pod.Name == podName {
				return true
			}
		}
	}
	return false
}

func getContext() (*rest.Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func configPath() (string, error) {
	if home := homeDir(); home != "" {
		return filepath.Join(home, ".kube", "config"), nil
	}
	return "", errors.New("HOME OR USERPROFILE env variables are not set")
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	return os.Getenv("USERPROFILE")
}

func newKubeClientWithRevision(kubeconfig, configContext string, revision string) (kube.ExtendedClient, error) {
	return kube.NewExtendedClient(kube.BuildClientCmd(kubeconfig, configContext), revision)
}

func newKubeClient(kubeconfig, configContext string) (kube.ExtendedClient, error) {
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

func KubectlIstioLog(pod string, namespace string, logLevel string, follow bool) {
	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}
	options, err := getOpts(context, namespace)
	if err != nil {
		log.Fatalln(err)
	}

	if options.isPodExists(pod) {
		log.Println("Pod exists")
		logNames, err := setupEnvoyLog("", pod, namespace)
		log.Println(logNames)
		if err != nil {
			log.Errorln(err)
		}
	} else {
		log.Errorln("Pod doesn't exist")
	}
}
