#!bin/bash

mkdir -p gen api
wget -qO- https://github.com/cgentron/api/archive/fcd5132bdeb82ee940cae1c15a606cda2eb13d41.zip | bsdtar --strip-components=1 -C gen -xvf-
protoc --go_out=plugins=grpc:./api -I./gen -I./proto --go_opt=paths=source_relative annotations.proto
protoc --go_out=plugins=grpc:./api -I./gen -I./proto --go_opt=paths=source_relative methods.proto
protoc --go_out=plugins=grpc:./api -I./gen -I./proto --go_opt=paths=source_relative fields.proto
protoc --go_out=plugins=grpc:./api -I./gen -I./proto --go_opt=paths=source_relative empty.proto
rm -rf gen