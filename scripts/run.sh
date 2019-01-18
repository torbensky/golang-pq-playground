#!/bin/bash

container_name='golang-playground-postgres'

echo "Checking for previous state"
if docker inspect "$container_name"  > /dev/null 2> /dev/null ; then
    if docker inspect -f '{{.State.Running}}' "$container_name"  > /dev/null 2> /dev/null ; then
        echo "Stopping previous container"
        docker stop $container_name
    fi
    echo "Removing previous container"
    docker rm $container_name
fi

echo "Starting container"
docker run -p 5432:5432 --name $container_name -e POSTGRES_PASSWORD=playground -d torbensky/golang-playground-postgres