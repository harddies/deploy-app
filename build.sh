#!/bin/bash

set -e

echo 'start building'
echo -n 'deploy? app?'
read build_type

if [ "$build_type" == "deploy" ]; then
  docker build -t deploy-platform -f ./build-deploy-platform --output ./ .
else
  echo '1. https://github.com/harddies/kratos-v2-demo'
  echo -n '请输入仓库地址编号：'
  read repo
  echo -n '请输入分支名称：'
  read branch_name
  echo -n '项目名称：'
  read project_name
  echo -n '服务名称：'
  read service_name

  if [ "$repo" == "1" ] || [ "$repo" == "" ]; then
    repo="https://github.com/harddies/kratos-v2-demo"
  fi

  cd /root/deploy-app
  DOCKER_BUILDKIT=1 docker build --build-arg project_name=${project_name} \
  --build-arg service_name=${service_name} \
  --build-arg repo=${repo} \
  --build-arg branch_name=${branch_name} \
  -t registry.my-cluster.co:5000/build-$project_name-$service_name \
  -f ./build-go-app --output out/$project_name-$service_name .
fi
echo 'finished building'