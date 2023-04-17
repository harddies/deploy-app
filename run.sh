#!/bin/bash

set -e

echo 'start building'
echo -n '项目名称：'
read project_name
echo -n '服务名称：'
read service_name

cd /root/deploy-app
version=`cat out/$project_name-$service_name/version`
docker build --build-arg project_name=${project_name} \
--build-arg service_name=${service_name} \
--label version=$version \
-t 192.168.1.128:5000/$project_name-$service_name \
-f ./run-go-app .

docker push 192.168.1.128:5000/$project_name-$service_name

echo 'finished building'