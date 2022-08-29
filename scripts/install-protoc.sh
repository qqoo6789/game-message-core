#!/bin/bash
set -e

echo "------------------ install protoc tool --------------"
sudo mkdir /usr/local/protobuf/
sudo chmod -R 777 /usr/local/protobuf/
cd /usr/local/protobuf/
curl 'https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protobuf-all-3.20.1.tar.gz' -o protobuf-all-3.20.1.tar.gz -v -L
tar -zxvf protobuf-all-3.20.1.tar.gz

cd ./protobuf-3.20.1/

sudo ./configure
sudo make
sudo make check
sudo make install 

echo "------------------ print protoc version  --------------"
protoc --version 


