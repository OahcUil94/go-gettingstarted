# 开发golang的基本流程

1. 创建用于存放项目的空目录，例如：gopath，gopath目录是开发golang项目的根目录
2. 一般下载golang的包，都会使用go get命令去下载，项目基本都是在github或者gitlab上面的，然后通过go get下载到本机环境，然后再进行开发，所以需要设置GOPATH环境变量，`export GOPATH=\`pwd\``，如果不设置的话，执行go get默认会下载到GOROOT下面，GOROOT就是安装golang程序的时候，默认的安装路径下面，我们一般是希望把项目依赖的包下载到gopath下面，这样不同的项目之间就不会相互污染
3. 在github上新建一个项目，例如：Hello，执行命令：`go get github.com/xxx/Hello`，此时gopath目录下就会有src目录，src下面的目录层级结构就是`github.com/xxx/Hello`