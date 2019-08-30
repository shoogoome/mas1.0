# liuma
由Helm管理的开箱即用的分布式文件服务器  
已实现功能: 分布式部署、数据校验去重、数据冗余和即时修复、断点续存、数据压缩

# 业务流程

1. 通过系统token在后端客户端发起请求获取上传或下载令牌，
2. 前端通过令牌进行上传或下载

# 安装说明

**前置条件: 系统配置好Kubernetes分布式环境**

- 本地版  
1. clone下当前资源仓库
2. 创建nfs服务并自行修改values.yaml配置
3. 执行helm install ./liuma 自行按需添加其他参数

- 线上版  
1. helm repo add '自定义仓库名' https://docker.hub.shoogoome.com/chartrepo/liuma
2. helm repo update
3. 创建nfs服务
3. helm install '自定义仓库名'/liuma 自行按需添加其他参数（按照values文件格式修改nfs配置）

**启动系统后需手动初始化mongo环境，数据库: 'liuma'**

# 接口说明
**ps: 需附带systemToken的接口应为后端访问接口，前端与系统交互应使用后端获取的临时token**
```
/server/token [put] 修改系统token
headers: systemToken 附带系统token
return: status success
```
```
/server/active [get] 获取活跃信号
headers: systemToken 附带系统token
return status 'ip列表'
```
```
/upload/token [get] 生成上传令牌
headers: systemToken 附带系统token
url参数: hash 文件hash
return token token
```
```
/download/token [get] 生成下载令牌
headers: systemToken 附带系统token
url参数: hash 文件hash
return token token
```
```
/upload/single [post] 单文件上传
headers: token 上传临时token
form-data: file 文件
return status success
```
```
/upload/chuck [post] 文件分块上传
headers: token 上传临时token
url参数: chuck 当前分片数
form-data: file 分片文件
return status success
```
```
/upload/finish [get] 完成上传(仅断点续传模式需要)
headers: token 上传临时token  
         systemToken 附带系统token
url参数: file_name 文件名称
return status success
```
```
/download [get] 下载文件
headers: token 下载临时token
return file
```

# 使用实践
个人在使用过程中，由于系统将存储服务独立抽象，不依赖业务系统，所以难免出现存储权限问题。  
目前我的解决方案是: 开放获取上传token接口，在业务后端添加完成上传接口，此接口无论单文件上传还是分片上传，前端在上传后均需要访问，以便后端进行权限控制。
此外，存储服务的完成上传接口不由前端访问，将由后端的完成上传接口进行触发，同样为了兼容权限控制。
使用中可根据业务需求让业务提供存储凭证以达到权限控制的目的; 而提供下载的token由业务在特定情况获取并传递给前端，并不直接暴露获取下载token功能

# PS
1.0版本存在许多问题，将在后续版本持续更新改善  
这个ps就当作2.0改版需求文档..
1. 服务发现功能，当前是在系统启动时统一初始化服务数量相关参数，无法做到使用过程中动态水平扩容。考虑在后续采用诸如RabbitMQ等手段改善服务发现功能
2. 当前版本没有实现资源版本控制功能
3. 当前使用的中间件 redis、mongo仍处于单机形式，后续将部署为集群可扩容形式
4. 与主系统之间识别手段单一简陋
5. 忘记支持断点下载了...
6. 采用statefulset策略部署server服务，由nfs提供存储服务，但一旦nfs服务器宕机则全面停止服务。2.0将支持其他存储服务
7. 没有查询、删除文件信息接口