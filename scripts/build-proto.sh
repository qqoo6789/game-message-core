#!/bin/bash
set -e

protoc --go_out=. ./proto/*.proto

