#!/bin/bash

set -e

echo 'start building'
echo -n '请输入分支名称：'
read branch_name
echo -n '项目名称：'
read project_name
echo -n '服务名称：'
read service_name

cd /root/deploy-app/kratos-v2-demo
git pull
git checkout ${branch_name}
version=`git log --pretty=oneline -n 1| awk '{print $1}'`
cd /root/deploy-app
DOCKER_BUILDKIT=1 docker build --build-arg project_name=${project_name} \
--build-arg service_name=${service_name} \
-t 192.168.1.128:5000/$project_name-$service_name \
-f ./build-go-app --output out/$project_name-$service_name .

touch out/$project_name-$service_name/version
echo $version > out/$project_name-$service_name/version

echo 'finished building'