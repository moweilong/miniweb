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


## ca证书

生成证书

```shell
make ca
```

创建一个测试用户

```shell
curl  -XPOST -H"Content-Type: application/json" -d'{"username":"catest","password":"miniweb1234","nickname":"catest","email":"catest@qq.com","phone":"18188888xxx"}' http://127.0.0.1:8080/v1/users
```

登录测试用户

```shell
token=`curl -s -XPOST -H"Content-Type: application/json" -d'{"username":"catest","password":"miniweb1234"}' http://127.0.0.1:8080/login | jq -r .token`
```

获取用户详细信息

1. 通过 HTTPS 协议访问 miniblog。不指定根证书，无法认证服务端证书报错

```shell
curl -XGET -H"Authorization: Bearer $token" https://127.0.0.1:8443/v1/users/catest
curl: (60) SSL certificate problem: EE certificate key too weak
More details here: https://curl.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.
```

2. 读取根证书，并使用根证书认证服务端

```shell
curl -XGET --cacert _output/cert/ca.crt -H"Authorization: Bearer $token" https://127.0.0.1:8443/v1/users/catest
```

3. 忽略 HTTPS 证书参数，指定跳过 SSL 检测

```shell
curl -XGET -k -H"Authorization: Bearer $token" https://127.0.0.1:8443/v1/users/catest
{"username":"catest","nickname":"catest","email":"catest@qq.com","phone":"18188888xxx","postCount":0,"createdAt":"2024-11-03 17:33:20","updatedAt":"2024-11-03 17:33:20"}
```