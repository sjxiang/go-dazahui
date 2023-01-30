package test

import (
	"errors"

	"gorm.io/gorm"
)

// 定义 GROM model
type Product struct {
	gorm.Model
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
	DeleteAt gorm.DeletedAt
}

func (u User) TableName() string {
	return "user"
}


type Email struct {
	ID    int64
	Name  string
	Email string
}

func (e Email) TableName() string {
	return "email"
}


/*
	
	Hook 钩子

		- CRUD 之前、之后 `自动调用` 的函数

		- 如果任何 Hook 返回错误，GORM 将停止后续操作，并回滚事务。

*/


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 0 {
		return errors.New("can't save invalid data")
	}

	return
} 

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Create(&Email{ID: u.ID, Email: u.Name + "@***.com"}).Error
}

