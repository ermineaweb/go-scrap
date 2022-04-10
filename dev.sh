#!/bin/bash

./build.sh
helm upgrade twitch ./helm/twitch --force
watch kubectl get pods 
