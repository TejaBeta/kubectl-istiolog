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
	"regexp"
	"testing"

	appv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestGetPod_A001(t *testing.T) {
	cs := testclient.NewSimpleClientset()
	input := &appv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-pod"}}

	options := options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Pods(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.isPodExists("unit-test-pod")
	if err != nil {
		t.Errorf("Error while looking for pod")
	}
}

func TestGetPod_A002(t *testing.T) {
	cs := testclient.NewSimpleClientset()
	input := &appv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-pod"}}

	options := options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Pods(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.isPodExists("unit-test-pod1")
	if err == nil {
		t.Errorf("Error during reterving pod that doesn't exist")
	}
}

func TestIstioLogGetPod_A001(t *testing.T) {
	cs := testclient.NewSimpleClientset()
	input := &appv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-pod"}}

	options := options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Pods(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.KubectlIstioLog("unit-test-pod1", "debug", false)
	if err == nil {
		t.Errorf("Error during reterving pod that doesn't exist")
	}
}

func TestHandleLoggerName_A001(t *testing.T) {
	err := handleLog("hello", "hello-pod", "hello-ns")
	if err == nil {
		t.Errorf("Error while using illegal loggerName")
	}
}

func TestHandleLoggerLevel_A001(t *testing.T) {
	err := handleLog("debug:hello", "hello-pod", "hello-ns")
	if err == nil {
		t.Errorf("Error while using illegal loggerLevel")
	}
}

func TestIstioLoggerName_A001(t *testing.T) {
	cs := testclient.NewSimpleClientset()
	input := &appv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-pod"}}

	options := options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Pods(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.KubectlIstioLog("unit-test-pod", "hello", false)
	if err == nil {
		t.Errorf("Error while using illegal loggerName")
	}
}

// We successfully parsed logger name & level and we failed on envoy call
func successfullyParsedLoggerNameAndLevel(err error) bool {
	errString := err.Error()
	return err != nil && regexp.MustCompile("failed to execute command on Envoy").MatchString(errString)
}

func TestIstioLoggerName_A002(t *testing.T) {
	cs := testclient.NewSimpleClientset()
	input := &appv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-pod"}}

	options := options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Pods(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.KubectlIstioLog("unit-test-pod", "debug", false)

	if !successfullyParsedLoggerNameAndLevel(err) {
		t.Errorf("Error while using illegal loggerLevel")
	}
}

func TestIstioLoggerLevel_A001(t *testing.T) {
	cs := testclient.NewSimpleClientset()
	input := &appv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-pod"}}

	options := options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Pods(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.KubectlIstioLog("unit-test-pod", "debug:hello", false)
	if err == nil {
		t.Errorf("Error while using illegal loggerLevel")
	}
}

func TestIstioLoggerLevel_A002(t *testing.T) {
	cs := testclient.NewSimpleClientset()
	input := &appv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-pod"}}

	options := options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().Pods(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.KubectlIstioLog("unit-test-pod", "http:debug", false)
	if !successfullyParsedLoggerNameAndLevel(err) {
		t.Errorf("Error while using illegal loggerName")
	}
}
