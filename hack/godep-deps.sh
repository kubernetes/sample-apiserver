#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# overall flow
# 1. make a clean gopath
# 2. get client-go
# 3. flatten client-go's vendored deps
# 4. go get extra deps we need
# 5. remove old vendoring data
# 6. vendor packages we need
# 7. copy new vendored packages and save them

set -o errexit
set -o nounset
set -o pipefail

goPath=$(mktemp -d "${TMPDIR:-/tmp/}$(basename 0).XXXXXXXXXXXX")
echo ${goPath}

export GOPATH=${goPath}

mkdir -p ${goPath}/src/k8s.io/sample-apiserver
cp -R . ${goPath}/src/k8s.io/sample-apiserver

pushd ${goPath}/src/k8s.io/sample-apiserver
rm -rf vendor || true
# this command fails because we've drifted from masters, but it gets us very close
go get -d ./... || true
popd

# now that we're close, we can stomp any dependencies that have drifted with their versions from client-go
pushd ${goPath}/src/k8s.io/client-go
godep restore
popd

# now we can mop up any we missed
go get github.com/inconshreveable/mousetrap/...
# because client-go is drifting behind (they still aren't auto-syncing well), we have to pull the latest 
# apimachinery and apiserver manually.  Hopefully client-go will build a sync script that wires together properly
rm -rf ${goPath}/src/k8s.io/apimachinery
rm -rf ${goPath}/src/k8s.io/apiserver
go get -d k8s.io/apimachinery/...
go get -d k8s.io/apiserver/...

# pin a couple unruly levels until client-go can be sync'ed with matching sets of the rest of staging
# TODO we can also pin these levels in a cloned and stripped down godep.json from the original k8s.io/kubernetes
pushd ${goPath}/src/github.com/grpc-ecosystem/grpc-gateway
git checkout f52d055dc48aec25854ed7d31862f78913cf17d1
popd
pushd ${goPath}/src/github.com/prometheus/client_golang
git checkout e51041b3fa41cece0dca035740ba6411905be473
popd
pushd ${goPath}/src/github.com/coreos/go-systemd
git checkout 4484981625c1a6a2ecb40a390fcb6a9bcfee76e3
popd

pushd ${goPath}/src/k8s.io/sample-apiserver
rm -rf vendor || true
rm -rf Godeps || true
godep save ./...
popd

# remove the vendor dir we have and move the one we just made
rm -rf vendor || true
rm -rf Godeps || true
git rm -rf vendor || true
git rm -rf Godeps || true
mv ${goPath}/src/k8s.io/sample-apiserver/vendor .
mv ${goPath}/src/k8s.io/sample-apiserver/Godeps .
git add vendor
git add Godeps
git commit -m "resync vendor folder"

