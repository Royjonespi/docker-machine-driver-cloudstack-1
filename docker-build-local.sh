#!/bin/sh

CI_PIPELINE_ID=$1
container="buildcontainer"
image="buildcontainerimage"
docker build -f "Dockerfile.multistage" -t $image .
docker run --name=$container -d $image
mkdir -p dist/
docker cp $container:/test/dm-cs-amd64-$CI_PIPELINE_ID.deb dist/
ls -alh dist
docker rm -f $container

docker rmi $image
