# DeployApp #
### 简介 ###
该部署脚本，适用于go-kratos v2框架。构建时根据询问输入对应的kratos目录名。kratos new XXX，则认为XXX是项目名称。cmd目录下的一级子目录名，则认为是服务名称

### 准备工作 ###
在需要构建go应用的宿主机上执行
```
docker build -t my-go-build:1.20.2 -f ./build .
```

### 操作步骤 ###
>1. sh ./build.sh。可构建deploy-platform和app
>2. sh ./run.sh。可构建运行镜像
>3. kubectl create -f deploy-app.yaml，可通过sh ./deploy.sh动态部署发布
>4. kubectl create -f ingress-app.yaml

### 未来 ###
>1. 区分env
>2. 区分tag/version
>3. 界面平台化
>4. etc.