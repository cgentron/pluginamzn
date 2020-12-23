#!bin/bash

mkdir -p gen
wget -qO- https://github.com/cgentron/api/archive/fcd5132bdeb82ee940cae1c15a606cda2eb13d41.zip | bsdtar --strip-components=1 -C gen -xvf-
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/annotations.proto
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/methods.proto
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/fields.proto
protoc --go_out=plugins=grpc:. -I./gen -I. --go_opt=module=github.com/cgentron/pluginamzn cgentron/amazon/empty.proto
rm -rf gen