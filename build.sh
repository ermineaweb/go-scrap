#!/bin/bash

SERVICES_NAME=(
    "speedometer"
    "stream-notifier"
    "server-rest"
)

for SERVICE_NAME in "${SERVICES_NAME[@]}"
do
    docker build \
    --build-arg SERVICE_NAME="${SERVICE_NAME}" \
    -t ermineaweb/"${SERVICE_NAME}":latest .

    docker push ermineaweb/"${SERVICE_NAME}":latest
done