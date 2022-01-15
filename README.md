## 概览

支持REST API访问的简易对象存储服务

## 代码结构

- server.go：管理服务生命周期
- api.go：路由请求
- model.go：请求体，响应体

## API

- 桶操作
  - GET / 查询桶列表
  - GET /BucketName 查询桶中对象列表
  - POST /BucketName 创建桶
  - DELETE /BucketName 删除桶

- 对象操作
  - GET /BucketName/ObjectName 查询对象
  - PUT /BucketName/ObjectName 上传对象
  - DELETE /BucketName/ObjectName 删除对象

## 编译

启用go module，在src目录下直接执行go build即可编译生成二进制文件


## 运行参数

- 服务端口：9000
- 数据存储路径：~/object_storage_service/data
