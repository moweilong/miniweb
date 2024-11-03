# db2struct

https://github.com/Shelnutt2/db2struct


安装

```shell
go install github.com/Shelnutt2/db2struct/cmd/db2struct@latest
```

使用

```shell
db2struct --gorm --no-json -H 127.0.0.1 --mysql_port 13306 -d miniweb -t user --package model --struct UserM -u miniweb -p 'miniweb1234' --target=user.go
db2struct --gorm --no-json -H 127.0.0.1 --mysql_port 13306 -d miniweb -t post --package model --struct PostM -u miniweb -p 'miniweb1234' --target=post.go
```

```shell
mysqldump -h127.0.0.1 -P13306 -uminiweb --databases miniweb -p'miniweb1234' --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > configs/miniweb.sql
```