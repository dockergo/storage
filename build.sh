#!/usr/bin/env bash

make

docker build -t 192.168.20.4:5000/agent:latest

docker push 192.168.20.4:5000/agent:latest