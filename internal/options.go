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
	"errors"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type options struct {
	clientset kubernetes.Interface
	namespace string
}

func GetOpts(context *rest.Config, ns string) (*options, error) {
	cs, err := kubernetes.NewForConfig(context)
	if err != nil {
		return nil, err
	}

	return &options{
		clientset: cs,
		namespace: ns,
	}, nil
}

func GetContext() (*rest.Config, error) {
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
