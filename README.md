# gozerolearn

go-zero go-zero整体上做为一个稍重的微服务框架，提供了微服务框架需要具备的通用能力，同时也只带一部分的强约束，例如针对web和rpc服务需要按照其定义的DSL的协议格式进行定义，日志配置、服务配置、apm配置等都要按照框架定义的最佳实践来走。 社区建设： go-zero已经是CNCF项目，做为一个后起的微服务框架，不得不说在国内社区生态建设和维护上，完美适配国内开源的现状，在微信群、公众号、各种大会等多渠道进行推广，社区也时常有文章指导实践。

greet start: go run greet.go -f etc/greet-api.yaml

greet 验证: curl -i http://localhost:8888/from/you //此处接口是预定义接口，作为验证使用。

go-zero 是一个集成了各种工程实践的web和rpc框架，通过弹性设计保障了大并发服务器的稳定性，经受了充分的实战检验。

官网：go-zero.dev

## 架构图
![img.png](img文件/架构.png)

# 什么是微服务？

顾名思义，其实就是微小的服务。

最早之前的业务系统都是单体项目，像之前的web项目都属于单体项目。

单体项目的弊端：例如：后端要改一个很小的地方，那么需要整个项目重新构建，然后停止整个项目然后重启项目。所以企业项目发布都选在深夜发布就是这个原因。

那么，如果是微服务呢？

可以将大系统按照功能或者是产品进行服务拆分，形成一个独立的服务。

以一个博客为例，将博客项目拆分为用户服务和文章服务，修改文章服务就只需要修改文章服务的部分，修改用户服务就改用户服务的部分

## 需要准备的工具：
1. etcd
2. mysql
3. protoc (转rpc代码 环境准备中有准备好的压缩包  解压放在go/bin目录下就行了)
4. goctl (https://zhuanlan.zhihu.com/p/624597859)

