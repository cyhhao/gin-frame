# 简介
gin 脚手架，集成了以下依赖

-  `github.com/go-ini/ini`  
    读取 ini 配置文件，用于不同环境的参数配置，支持从环境变量读取配置
    
- `github.com/sirupsen/logrus`  
    打印 log，支持各种格式、颜色的 log 打印
    
- `github.com/jinzhu/gorm`    
    著名 ORM
    
以上几个都是各自功能类库里 star 数最多的

# 启动
```shell
# 开发环境
go run main.go --env dev
# 测试环境
go run main.go --env test
# 生产环境
go run main.go --env prod
```

# 配置文件
`conf/` 目录下默认有三个环境的配置文件：

- app.ini 为通用配置，所有环境继承此配置文件
- dev.ini 相同的配置会覆盖 app.ini 中的配置，对应 --env dev 
- test.ini 对应 --env test
- prod.ini 对应 --env prod
- 新建 xxx.ini ，需修改 `utils/config.go` 内相关代码即可

配置文件中使用 `${xxx}` 可从环境变量中读取参数

如：
```
# prod.ini
[mysql]
password = ${MYSQL_PASSWORD}

# shell
export MYSQL_PASSWORD=password
``` 

# 热更新debug
推荐使用 `github.com/cosmtrek/air`

`.air.conf`是一份默认配置，若已安装 air ，直接运行 air 即可


