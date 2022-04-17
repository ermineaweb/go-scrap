#!/bin/bash

helm uninstall twitch
helm install twitch ./helm/twitch 
watch kubectl get pods 
