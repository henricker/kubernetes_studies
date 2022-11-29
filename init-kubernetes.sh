#!/usr/env/bin bash

# This script is used to initialize the kubernetes cluster

# Create clusters in kind
kind create cluster --config ./k8s/kind.yml --name=fullcycle

# change context to kind cluster
kubectl cluster-info --context kind-fullcycle

# RUN ENVS
kubectl apply -f ./k8s/envs/config-map-family.yml
kubectl apply -f ./k8s/envs/config-map.yml
kubectl apply -f ./k8s/envs/secret.yml

# RUN SERVICES
kubectl apply -f ./k8s/services/loadBalance-service.yml

# RUN DEPLOYMENTS
kubectl apply -f ./k8s/deployments/deployment.yml

# port-forward
kubectl port-forward service/goserver 8000:80
