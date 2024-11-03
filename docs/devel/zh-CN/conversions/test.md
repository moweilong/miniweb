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