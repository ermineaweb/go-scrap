#!/bin/bash

SERVICE_NAME="chat-speedometer"
docker build --build-arg SERVICE_NAME=${SERVICE_NAME} -t ${SERVICE_NAME}:latest .

SERVICE_NAME="scrap"
docker build --build-arg SERVICE_NAME=${SERVICE_NAME} -t ${SERVICE_NAME}:latest .