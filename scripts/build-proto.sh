#!/bin/bash
set -e

readonly PROTO_ROOT="proto" 
protoc --proto_path=${PROTO_ROOT} --go_out=. ${PROTO_ROOT}/*.proto


# ------------- create dotnet use proto file ----------------
readonly PROTO_NET_OUT="./messageDotNet/proto"
readonly PROTO_BUILD_NET="./messageDotNet/protoNet"
rm -rf ${PROTO_BUILD_NET}
cp -r ${PROTO_ROOT} ${PROTO_BUILD_NET}
sed -i -e '/go_package/d' ${PROTO_BUILD_NET}/*.proto 
# sed -i -e 's/package .*/package MelandGame3;/g' ${PROTO_BUILD_NET}/*.proto
protoc --proto_path=${PROTO_BUILD_NET} --csharp_out=${PROTO_NET_OUT}  ${PROTO_BUILD_NET}/*.proto
rm -rf ${PROTO_BUILD_NET}

 