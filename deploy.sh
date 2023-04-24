#!/bin/bash

set -e

echo 'start building'
echo -n '项目名称：'
read project_name
echo -n '服务名称：'
read service_name
echo -n '版本号：'
read version
echo -n '实例数量：'
read pod_count

./deploy-platform -tpl_dir=./tpl -data='{"project_name":"'${project_name}'","service_name":"'${service_name}'","version":"'${version}'","pod_count":'${pod_count}}