#!/bin/bash

protoc --go_out=paths=source_relative:. peerlist.proto
