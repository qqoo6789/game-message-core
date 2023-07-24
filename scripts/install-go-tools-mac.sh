#!/bin/bash
set -o errexit

readonly GO_VERSION="go1.20.4"

echo "------------------ install golang version manager 'gvm' --------------"
rm -rf $HOME/.gvm
${SHELL} < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source $HOME/.gvm/scripts/gvm    


echo "------------------ install golang base version  --------------"
brew install go  

echo "------------------ install golang version ${GO_VERSION} --------------"
gvm install ${GO_VERSION}
gvm use ${GO_VERSION} --default 
go version
go env -w GO111MODULE=auto
go env -w GOPROXY=https://goproxy.cn,direct

echo "------------------ remove golang base version --------------"
brew uninstall go


echo "------------------ install golang protobuf tool 'protoc-gen-go' --------------"
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


