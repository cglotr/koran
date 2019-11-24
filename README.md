# koran

[![Build Status](https://travis-ci.com/cglotr/koran.svg?branch=master)](https://travis-ci.com/cglotr/koran)
[![codecov](https://codecov.io/gh/cglotr/koran/branch/master/graph/badge.svg)](https://codecov.io/gh/cglotr/koran)

## Setup & development

Steps below assume that you're connected to a Kubernetes cluster.

### One-time setup

- [Install](https://helm.sh/docs/intro/install/#from-script) _Helm_ & add the official chart repository
  - `helm repo add stable https://kubernetes-charts.storage.googleapis.com/`
- `helm install nginx-ingress stable/nginx-ingress`
- `kubectl create secret generic secret-mysql-password --from-literal=MYSQL_PASSWORD=$MYSQL_PASSWORD`
- `kubectl apply -f kubernetes`
- `./migrate.sh $MYSQL_PASSWORD`
- `./seed.sh $MYSQL_PASSWORD`
- [Install](https://skaffold.dev/docs/install/) _Skaffold_

### Development

- `skaffold dev`
