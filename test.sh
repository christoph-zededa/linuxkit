#!/usr/bin/bash

set -e
set -x

docker build -t sometestcontainerimage some_test_container_images/
docker build pkg/getty

echo THIS WORKS
echo THIS WORKS
echo THIS WORKS
echo THIS WORKS
echo THIS WORKS


go build -C src/cmd/linuxkit
./src/cmd/linuxkit/linuxkit build linuxkit-template.yml

echo THIS DOES NOT WORK
