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
echo -n '实例数量：'
read pod_count

if [ "$version" == "" ]; then
  version=`ls -lcr out/$project_name-$service_name | awk '{if(NR>1) print $NF}' | head -n 1`
fi

./deploy-platform -tpl_dir=./tpl -data='{"project_name":"'${project_name}'","service_name":"'${service_name}'","git_version_tag":"'${version}'","pod_count":'${pod_count}}