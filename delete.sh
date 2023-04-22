#!/bin/bash

set -e

echo 'start building'
echo -n '项目名称：'
read project_name
echo -n '服务名称：'
read service_name

kubectl delete deployment feature-flag-admin -n app
kubectl delete svc feature-flag-admin-svc -n app