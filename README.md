# liuma
由Helm管理的开箱即用的分布式文件服务器

# 业务流程

1. 通过系统token在后端客户端发起请求获取上传或下载令牌，
2. 前端通过令牌进行上传或下载

# 接口说明
```
/server/token [put] 修改系统token
headers: system_token 附带系统token
return: status success
```
```
/server/signal [get] 获取活跃信号
headers: system_token 附带系统token
return ip 'ip列表'
```
```
/upload/token [get] 生成上传令牌
headers: system_token 附带系统token
url参数: hash 文件hash
return token token
```
```
/download/token [get] 生成下载令牌
headers: system_token 附带系统token
url参数: hash 文件hash
return token token
```
# PS
1.0版本存在许多问题，将在后续版本持续更新改善
1. 服务发现功能，当前是在系统启动时统一初始化服务数量相关参数，无法做到使用过程中动态水平扩容。考虑在后续采用诸如RabbitMQ等手段改善服务发现功能
2. 当前版本没有实现资源版本控制功能
3. 当前使用的中间件 redis、mongo仍处于单机形式，后续将部署为集群可扩容形式
4. 与主系统之间识别手段单一简陋