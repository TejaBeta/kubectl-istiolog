// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1

import (
	v1alpha1 "istio.io/api/meta/v1alpha1"
	securityv1 "istio.io/api/security/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthorizationPolicy enables access control on workloads.
//
// <!-- crd generation tags
// +cue-gen:AuthorizationPolicy:groupName:security.istio.io
// +cue-gen:AuthorizationPolicy:version:v1
// +cue-gen:AuthorizationPolicy:annotations:helm.sh/resource-policy=keep
// +cue-gen:AuthorizationPolicy:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:AuthorizationPolicy:subresource:status
// +cue-gen:AuthorizationPolicy:scope:Namespaced
// +cue-gen:AuthorizationPolicy:resource:categories=istio-io,security-istio-io,plural=authorizationpolicies
// +cue-gen:AuthorizationPolicy:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:security/v1beta1/authorization_policy.proto
// -->
type AuthorizationPolicy struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec securityv1.AuthorizationPolicy `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthorizationPolicyList is a collection of AuthorizationPolicies.
type AuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []*AuthorizationPolicy `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RequestAuthentication defines what request authentication methods are supported by a workload.
// It will reject a request if the request contains invalid authentication information, based on the
// configured authentication rules. A request that does not contain any authentication credentials
// will be accepted but will not have any authenticated identity. To restrict access to authenticated
// requests only, this should be accompanied by an authorization rule.
// Examples:
//
// - Require JWT for all request for workloads that have label `app:httpbin`
//
// {{<tabset category-name="example">}}
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	jwtRules:
//	- issuer: "issuer-foo"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ```
// {{</tab>}}
//
// {{<tab name="v1" category-value="v1">}}
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	jwtRules:
//	- issuer: "issuer-foo"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ```
// {{</tab>}}
// {{</tabset>}}
//
// - A policy in the root namespace ("istio-system" by default) applies to workloads in all namespaces
// in a mesh. The following policy makes all workloads only accept requests that contain a
// valid JWT token.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//
//	name: req-authn-for-all
//	namespace: istio-system
//
// spec:
//
//	jwtRules:
//	- issuer: "issuer-foo"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//
//	name: require-jwt-for-all
//	namespace: istio-system
//
// spec:
//
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ```
// {{</tab>}}
//
// {{<tab name="v1" category-value="v1">}}
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: req-authn-for-all
//	namespace: istio-system
//
// spec:
//
//	jwtRules:
//	- issuer: "issuer-foo"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: require-jwt-for-all
//	namespace: istio-system
//
// spec:
//
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ```
// {{</tab>}}
// {{</tabset>}}
//
// - The next example shows how to set a different JWT requirement for a different `host`. The `RequestAuthentication`
// declares it can accept JWTs issued by either `issuer-foo` or `issuer-bar` (the public key set is implicitly
// set from the OpenID Connect spec).
//
// {{<tabset category-name="example">}}
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	jwtRules:
//	- issuer: "issuer-foo"
//	- issuer: "issuer-bar"
//
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["issuer-foo/*"]
//	  to:
//	  - operation:
//	      hosts: ["example.com"]
//	- from:
//	  - source:
//	      requestPrincipals: ["issuer-bar/*"]
//	  to:
//	  - operation:
//	      hosts: ["another-host.com"]
//
// ```
//
// {{</tab>}}
//
// {{<tab name="v1" category-value="v1">}}
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	jwtRules:
//	- issuer: "issuer-foo"
//	- issuer: "issuer-bar"
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["issuer-foo/*"]
//	  to:
//	  - operation:
//	      hosts: ["example.com"]
//	- from:
//	  - source:
//	      requestPrincipals: ["issuer-bar/*"]
//	  to:
//	  - operation:
//	      hosts: ["another-host.com"]
//
// ```
// {{</tab>}}
// {{</tabset>}}
//
// - You can fine tune the authorization policy to set different requirement per path. For example,
// to require JWT on all paths, except /healthz, the same `RequestAuthentication` can be used, but the
// authorization policy could be:
//
// {{<tabset category-name="example">}}
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//	- to:
//	  - operation:
//	      paths: ["/healthz"]
//
// ```
// {{</tab>}}
//
// {{<tab name="v1" category-value="v1">}}
// ```yaml
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//	- to:
//	  - operation:
//	      paths: ["/healthz"]
//
// ```
// {{</tab>}}
// {{</tabset>}}
//
// [Experimental] Routing based on derived [metadata](https://istio.io/latest/docs/reference/config/security/conditions/)
// is now supported. A prefix '@' is used to denote a match against internal metadata instead of the headers in the request.
// Currently this feature is only supported for the following metadata:
//
// - `request.auth.claims.{claim-name}[.{sub-claim}]*` which are extracted from validated JWT tokens. The claim name
// currently does not support the `.` character. Examples: `request.auth.claims.sub` and `request.auth.claims.name.givenName`.
//
// The use of matches against JWT claim metadata is only supported in Gateways. The following example shows:
//
// - RequestAuthentication to decode and validate a JWT. This also makes the `@request.auth.claims` available for use in the VirtualService.
// - AuthorizationPolicy to check for valid principals in the request. This makes the JWT required for the request.
// - VirtualService to route the request based on the "sub" claim.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//
//	name: jwt-on-ingress
//	namespace: istio-system
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: istio-ingressgateway
//	jwtRules:
//	- issuer: "example.com"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//
//	name: require-jwt
//	namespace: istio-system
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: istio-ingressgateway
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ---
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//
//	name: route-jwt
//
// spec:
//
//	hosts:
//	- foo.prod.svc.cluster.local
//	gateways:
//	- istio-ingressgateway
//	http:
//	- name: "v2"
//	  match:
//	  - headers:
//	      "@request.auth.claims.sub":
//	        exact: "dev"
//	  route:
//	  - destination:
//	      host: foo.prod.svc.cluster.local
//	      subset: v2
//	- name: "default"
//	  route:
//	  - destination:
//	      host: foo.prod.svc.cluster.local
//	      subset: v1
//
// ```
// {{</tab>}}
//
// {{<tab name="v1" category-value="v1">}
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: jwt-on-ingress
//	namespace: istio-system
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: istio-ingressgateway
//	jwtRules:
//	- issuer: "example.com"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: require-jwt
//	namespace: istio-system
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: istio-ingressgateway
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ---
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//
//	name: route-jwt
//
// spec:
//
//	hosts:
//	- foo.prod.svc.cluster.local
//	gateways:
//	- istio-ingressgateway
//	http:
//	- name: "v2"
//	  match:
//	  - headers:
//	      "@request.auth.claims.sub":
//	        exact: "dev"
//	  route:
//	  - destination:
//	      host: foo.prod.svc.cluster.local
//	      subset: v2
//	- name: "default"
//	  route:
//	  - destination:
//	      host: foo.prod.svc.cluster.local
//	      subset: v1
//
// ```
// {{</tab>}}
// {{</tabset>}}
//
// <!-- crd generation tags
// +cue-gen:RequestAuthentication:groupName:security.istio.io
// +cue-gen:RequestAuthentication:version:v1
// +cue-gen:RequestAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:RequestAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:RequestAuthentication:subresource:status
// +cue-gen:RequestAuthentication:scope:Namespaced
// +cue-gen:RequestAuthentication:resource:categories=istio-io,security-istio-io,shortNames=ra
// +cue-gen:RequestAuthentication:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-from:security/v1beta1/request_authentication.proto
// -->
type RequestAuthentication struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec securityv1.RequestAuthentication `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RequestAuthenticationList is a collection of RequestAuthentications.
type RequestAuthenticationList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []*RequestAuthentication `json:"items" protobuf:"bytes,2,rep,name=items"`
}
