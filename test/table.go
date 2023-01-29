package test

import (
	// "gorm.io/gorm"
)


// 定义 GROM model
type Product struct {
	Code  string  
	Price uint   
}
  
// 为 model 定义表名
func (p Product) TableName() string {
	return "product"
}


/*

+----+-------------------------+-------------------------+------------+------+-------+
| id | created_at              | updated_at              | deleted_at | code | price |
+----+-------------------------+-------------------------+------------+------+-------+
|  1 | 2023-01-29 16:27:17.523 | 2023-01-29 16:27:17.523 | NULL       | 101  |     0 |
+----+-------------------------+-------------------------+------------+------+-------+

*/


type User struct {
	ID   int64  `gorm:"primarykey"`
	Name string `gorm:"column:nick_name; default:galeone"`  // 通过使用 default 标签，为字段定义默认值
	Age  int64  `gorm:"column:age;       default:18"`
}

func (u User) TableName() string {
	return "user"
}

