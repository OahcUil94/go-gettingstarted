# 创建Dockerfile镜像文件

## 构建docker镜像

构建`3.3.11`版本的镜像：

`docker build -t oahcuil/etcd:3.3.11 --build-arg ETCD_VER=v3.3.11 .`

## 运行etcd容器

`docker run --rm -it -p 2379:2379 -v /usr/local/var/etcd/data/:/data --name my-etcd -d oahcuil/etcd:3.3.11`

mac上可能会遇到这样的问题：

```
$ docker run --rm -it -p 2379:2379 -v /usr/local/var/etcd/data/:/data --name my-etcd -d oahcuil/etcd:3.3.11
c470d04e3dabbb98e300688560bd5378760025daa800dc70b0f61d8516ef53ac
docker: Error response from daemon: Mounts denied:
The path /usr/local/var/etcd/data/
is not shared from OS X and is not known to Docker.
You can configure shared paths from Docker -> Preferences... -> File Sharing.
See https://docs.docker.com/docker-for-mac/osxfs/#namespaces for more info.
```

> 遇到这个问题，需要把/usr/local/var/目录添加到docker的文件共享目录中：Docker -> Preferences... -> File Sharing，然后Apply&Restart

## 使用etcdctl测试连接

`./etcdctl --endpoints=http://0.0.0.0:2379 --debug ls`

返回如下信息表示连接正常：

```
$ ./etcdctl --endpoints=http://127.0.0.1:2379 --debug ls
start to sync cluster using endpoints(http://127.0.0.1:2379)
cURL Command: curl -X GET http://127.0.0.1:2379/v2/members
got endpoints(http://0.0.0.0:4001,http://0.0.0.0:2379) after sync
Cluster-Endpoints: http://0.0.0.0:4001, http://0.0.0.0:2379
cURL Command: curl -X GET http://0.0.0.0:4001/v2/keys/?quorum=false&recursive=false&sorted=false
cURL Command: curl -X GET http://0.0.0.0:2379/v2/keys/?quorum=false&recursive=false&sorted=false
```

## 打包遇到的问题

遇到两个非常坑的问题，最终还没有解决：
- run.sh的开头添加`!#/bin/sh`，则会报错，`!#/bin/sh: not found`，但是不会停止运行
- 无论run.sh开头是否添加`!#/bin/sh`，使用ENTRYPOINT运行，都会报错：`exec user process caused "exec format error"`
