#!/bin/bash

set -e

echo 'start building'
echo -n '项目名称：'
read project_name
echo -n '服务名称：'
read service_name

ls -lcr out/$project_name-$service_name | awk '{if(NR>1)print $NF,$(NF-3)$(NF-2),$(NF-1)}'

echo -n '版本号：[latest-version]'
read version

if [ "$version" == "" ]; then
  version=`ls -lcr out/$project_name-$service_name | awk '{if(NR>1) print $NF}' | head -n 1`
fi

cd /root/deploy-app
docker build --build-arg project_name=${project_name} \
--build-arg service_name=${service_name} \
--label version=$version \
-t registry.my-cluster.co:5000/$project_name-$service_name \
-f ./run-go-app .

docker push registry.my-cluster.co:5000/$project_name-$service_name

docker tag registry.my-cluster.co:5000/$project_name-$service_name registry.my-cluster.co:5000/$project_name-$service_name:$version

docker push registry.my-cluster.co:5000/$project_name-$service_name:$version

docker builder prune -f

echo 'finished building'