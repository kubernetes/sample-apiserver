# sample-apiserver

Demonstration of how to use the k8s.io/apiserver library to build a functional API server.


## Purpose

You may use this code if you want to build an Extension API Server to use with API Aggregation, or to build a stand-alone Kubernetes-style API server.


However, note that there are several other options, all of which are likely better choices than using this repository.

  * If you just want to add a resource to your kubernetes cluster, then you have two choices:
    * Extension API Server (what this repo implements)
    * Custom Resource Definition a.k.a CRDs (easier and less coding and rebasing).
    * Read about the differences between Custom Resource Definitions vs Extension API Servers [here](https://kubernetes.io/docs/concepts/api-extension/custom-resources).

If you do decide to build an Extension API server, consider using [apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder) instead of this repo, which is a complete framework for generating the both the apiserver and client libraries, and the installation program.

If you do decide to use this repository, then the recommended pattern is to fork this repository, modify it to add your types, and then periodically rebase your changes on top of this repo, to pick up improvements and bug fixes to the apiserver.


## Compatibility

HEAD of this repo will match HEAD of k8s.io/apiserver, k8s.io/apimachinery, and k8s.io/client-go.

## Where does it come from?

`sample-apiserver` is synced from https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/sample-apiserver.
Code changes are made in that location, merged into `k8s.io/kubernetes` and later synced here.

