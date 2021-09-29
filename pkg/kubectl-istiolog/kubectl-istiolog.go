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

package pkg

import (
	"context"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

func (kOpts *options) isPodExists(podName string) bool {
	result, err := kOpts.clientset.CoreV1().Pods(kOpts.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		for _, pod := range result.Items {
			if pod.Name == podName {
				return true
			}
		}
	}
	return false
}

func KubectlIstioLog(pod string, namespace string, logLevel string, follow bool) {
	log.Println(pod, namespace, logLevel, follow)
}
