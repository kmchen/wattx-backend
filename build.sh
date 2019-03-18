#!/bin/bash


export GOOS=linux
export CGO_ENABLED=0

docker network create wattx-network

# stop/remove rankingservice
docker stop rankingservice
docker rm rankingservice
# Build ranking service
cd ranking;go get -d;go build -o ranking-linux-amd64;echo built `pwd`;
docker build -t kmchen/rankingservice .
docker create --name rankingservice -p=50051:50051 --network=wattx-network -p=8888:8888 kmchen/rankingservice:latest
docker start rankingservice
cd -

rankingServiceIp=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' rankingservice)

# stop/remove pricingservice
docker stop pricingservice
docker rm pricingservice
# Build pricing service
cd pricing;go build -o pricing-linux-amd64;echo built `pwd`;
docker build -t kmchen/pricingservice --build-arg serverIp=$rankingServiceIp .
docker create --name pricingservice -p=55001:55001 --network=wattx-network kmchen/pricingservice:latest
docker start pricingservice

export GOOS=darwin
