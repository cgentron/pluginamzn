#!bin/bash

mkdir -p gen
wget -qO- https://github.com/cgentron/api/archive/e2c0aa4cdf4264674bf813dcbd2ac91670c70cb0.zip | bsdtar --strip-components=1 -C gen -xvf-
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/annotations.proto
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/methods.proto
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/fields.proto
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/empty.proto
rm -rf gen