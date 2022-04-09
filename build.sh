#!/bin/bash

SERVICE_NAME="chat-speedometer"
docker build --build-arg SERVICE_NAME=${SERVICE_NAME} -t ermineaweb/${SERVICE_NAME}:latest .
docker push ermineaweb/${SERVICE_NAME}:latest

SERVICE_NAME="stream-notifier"
docker build --build-arg SERVICE_NAME=${SERVICE_NAME} -t ermineaweb/${SERVICE_NAME}:latest .
docker push ermineaweb/${SERVICE_NAME}:latest