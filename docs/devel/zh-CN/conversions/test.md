# 测试API

## 创建用户

```shell
 curl -XPOST -H"Content-Type: application/json" -d'{"username":"root","password":"miniweb1234","nickname":"root","email":"kalandramo@gmail.com","phone":"18188889999"}' http://127.0.0.1:8080/v1/users
```

## 登录

创建测试用户

```shell
curl -XPOST -H"Content-Type: application/json" -d'{"username":"authntest","password":"authntest1234","nickname":"authntest","email":"authntest@qq.com","phone":"1818888xxxx"}' http://127.0.0.1:8080/v1/users
```

登录

```shell
curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"authntest","password":"authntest1234"}' http://127.0.0.1:8080/login
```

## 修改密码

```shell
curl -XPUT -H"Content-Type: application/json" -d'{"oldPassword":"authntest1234","newPassword":"authntest12345"}' http://127.0.0.1:8080/v1/users/authntest/change-password
```

验证密码

```shell
curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"authntest","password":"authntest1234"}' http://127.0.0.1:8080/login

curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"authntest","password":"authntest12345"}' http://127.0.0.1:8080/login
```

## 认证和授权

```shell
# 创建用户 belma
curl -XPOST -H"Content-Type: application/json" -d'{"username":"belma","password":"miniweb1234","nickname":"belma","email":"nosbelma@qq.com","phone":"18188888xxx"}' http://127.0.0.1:8080/v1/users

# 创建用户 belmb
curl -XPOST -H"Content-Type: application/json" -d'{"username":"belmb","password":"miniweb1234","nickname":"belmb","email":"nosbelmb@qq.com","phone":"18188888xxx"}' http://127.0.0.1:8080/v1/users

# belma 用户登录 miniweb 平台
token=`curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"belma","password":"miniweb1234"}' http://127.0.0.1:8080/login | jq -r .token`

# belma 获取 belma 的详细信息
curl -XGET -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/users/belma
{"username":"belma","nickname":"belma","email":"nosbelma@qq.com","phone":"18188888xxx","postCount":0,"createdAt":"2024-11-03 16:45:45","updatedAt":"2024-11-03 16:45:45"}

# belma 获取 belmb 的详细信息
curl -XGET -H"Authorization: Bearer $token" http://127.0.0.1:8080/v1/users/belmb
{"code":"AuthFailure.Unauthorized","message":"Unauthorized."}
```