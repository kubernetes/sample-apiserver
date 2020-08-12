// This is a generated file. Do not edit directly.

module k8s.io/sample-apiserver

go 1.13

require (
	github.com/go-openapi/spec v0.19.3
	github.com/google/gofuzz v1.1.0
	github.com/spf13/cobra v0.0.5
	k8s.io/apimachinery v0.0.0-20200812051649-2a282836017b
	k8s.io/apiserver v0.0.0-20200812054331-15d0d4879581
	k8s.io/client-go v0.0.0-20200812053003-c8494005d10a
	k8s.io/code-generator v0.0.0-20200708172309-f186a36abf5c
	k8s.io/component-base v0.0.0-20200812053641-bf8cd1993483
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6 // release-1.18
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20200812052025-e70c4f3b2093
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200812051649-2a282836017b
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20200812054331-15d0d4879581
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200812053003-c8494005d10a
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200708172309-f186a36abf5c
	k8s.io/component-base => k8s.io/component-base v0.0.0-20200812053641-bf8cd1993483
)
