# Stack-RPC 示例集合

## 目录

[老目录（Go-Micro）](./deprecated)

- [Client](./client)
  - [Retry](./client/retry) 重试
- [Logger](./logger) 日志
  - [Logrus](./logger/logrus) logrus 插件
    - 持久化、压缩归档、清理
- [Config](./config)
  - [Apollo](./config/apollo) Apollo 配置中心
  - [Stack-Plugin](./config/stack) [TODO] Stack 配置中心插件
- [Service](./service)
  - [GetIP](./service/getip) 获取服务与调用者ip
  - [Metadata](./service/metadata) 上下文信息传递
  - [NoProto](./service/noproto) 免proto文件接口，不想写proto文件可参考
  - [RPC](./service/rpc) RPC服务
  - [Stream](./service/stream) 流服务
  - [Web](./service/web) Web HTTP服务
    - 编写RPC服务  
- [micro-wrapper](./wrapper) 
  - [基础用法](./wrapper/basic)
  - [WrapAPI](./wrapper/handler) todo
  - [WrapCall](./wrapper/call)
  - [WrapClient](./wrapper/client)