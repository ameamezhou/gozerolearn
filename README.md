# gozerolearn

go-zero go-zero整体上做为一个稍重的微服务框架，提供了微服务框架需要具备的通用能力，同时也只带一部分的强约束，例如针对web和rpc服务需要按照其定义的DSL的协议格式进行定义，日志配置、服务配置、apm配置等都要按照框架定义的最佳实践来走。 社区建设： go-zero已经是CNCF项目，做为一个后起的微服务框架，不得不说在国内社区生态建设和维护上，完美适配国内开源的现状，在微信群、公众号、各种大会等多渠道进行推广，社区也时常有文章指导实践。

greet start: go run greet.go -f etc/greet-api.yaml

greet 验证: curl -i http://localhost:8888/from/you //此处接口是预定义接口，作为验证使用。


