#!/bin/bash

helm upgrade --install --force twitch ./helm/twitch 
watch kubectl get pods 
