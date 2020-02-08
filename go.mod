// This is a generated file. Do not edit directly.

module k8s.io/sample-apiserver

go 1.13

require (
	github.com/go-openapi/spec v0.19.3
	github.com/google/gofuzz v1.1.0
	github.com/spf13/cobra v0.0.5
	k8s.io/apimachinery v0.0.0-20200207025655-52a338251bb2
	k8s.io/apiserver v0.0.0-20200208150703-8a2af90395f3
	k8s.io/client-go v0.0.0-20200208144352-4b7e8bfcc145
	k8s.io/code-generator v0.0.0-20200208144352-dd255cc51571
	k8s.io/component-base v0.0.0-20200207030544-616550b070ba
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20200207025841-85a41f27a10c
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200207025655-52a338251bb2
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20200208150703-8a2af90395f3
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200208144352-4b7e8bfcc145
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200208144352-dd255cc51571
	k8s.io/component-base => k8s.io/component-base v0.0.0-20200207030544-616550b070ba
)
