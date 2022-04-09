#!/bin/bash

helm uninstall twitch
./build.sh
helm install twitch ./helm/twitch
