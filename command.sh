#!/bin/bash


if [ $# -ne 1 ]
then
    echo "missing value"
    exit 1
fi

currentIP=$(minikube service collatz-app-service --url)

for i in $(seq $1); do
     curl "$currentIP/collatz/$i"
     echo
done