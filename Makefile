#!/usr/bin/env bash

set -e

#source /root/.bashrc

version=`git describe --abbrev=0 --tags`
echo "Build Version: ${version}"
#go_version=`go version`
#commit=`git log --pretty=format:"%h" -1`
#date=`date +'%Y-%m-%d %H:%M:%S'`
#branch=`git rev-parse --abbrev-ref HEAD`
#
#go build -mod=vendor -o bin/order-api -ldflags "-X 'main.version=${version}' -X 'main.goVersion=${go_version}' -X 'main.commit=${commit}' -X 'main.date=${date}'" cmd/main.go

make

elderImageId=`docker images | grep 'order-api' | grep 'latest'| awk '{print $3}'`
docker build -t hub.ucloudadmin.com/upangu/order-api:latest .

# stop and rm container
#isRunningContainer=`docker ps | grep -c 'order-api'`
isRunningContainer=`docker ps | grep 'order-api' | wc -l`
if [[ ${isRunningContainer} -gt 0 ]]; then
    containerId=`docker ps -a | grep 'order-api' | awk '{print $1}'`
    echo "Stopping container ${containerId} ... "
    docker stop order-api
    echo "Removing container ${containerId} ... "
    docker rm order-api
fi

currentImageId=`docker images | grep 'order-api' | grep 'latest'| awk '{print $3}'`
if [[ "$elderImageId" != "$currentImageId" ]]; then
    echo "Removing previous image ${elderImageId} ..."
    docker rmi ${elderImageId}
fi

docker run \
    --name order-api \
    --detach \
    --restart=always \
    --publish 15103:15103 \
    hub.ucloudadmin.com/upangu/order-api:latest

if [[ ${DEVELOPMENTENV} == host ]]; then
    docker tag hub.ucloudadmin.com/upangu/order-api:latest hub.ucloudadmin.com/upangu/order-api:${version}
    docker push hub.ucloudadmin.com/upangu/order-api:${version}
fi

echo "Application was integrated and deployed SUCCESSFULLY! - ${DEVELOPMENTENV}"
