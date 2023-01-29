

# GORM 的基本使用


GORM 的约定（默认）


# 建表相关

1. 使用结构体`蛇形复数`作为表名
    
    避免方法
    - 实现 TableName()
    - gorm 命名策略更改

```go
import (
    "gorm.io/gorm/schema"
)

_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: false,
	},
})

```



2. 字段名的`蛇形`作为列名
    避免方法
    - GROM Tag `gorm:"column:xxx"`


3. 一般多 4 个字段
    - 使用名为 ID 的字段作为主键
    - 使用 CreatedAt、UpdatedAt、DeletedAt 字段作为创建、更新、删除时间



