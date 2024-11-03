# 问题

## make 编译报错

### gorm.io/plugin/dbresolver

报错信息

```shell
gorm.io/plugin/dbresolver
# gorm.io/plugin/dbresolver
../../../../pkg/mod/gorm.io/plugin/dbresolver@v1.3.0/dbresolver.go:141:5: unknown field PreparedSQL in struct literal of type gorm.PreparedStmtDB
```

解决方案

```shell
go get -u gorm.io/plugin/dbresolver
```
