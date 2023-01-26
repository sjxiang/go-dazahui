package test

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


const dsn = "root:123456@tcp(172.20.0.2:3306)/crud-demo?charset=utf8&parseTime=True&loc=Local"

type Product struct {
	ID    uint    `gorm:"primarykey"`
	Code  string  `gorm:"column: code"`
	Price uint    `gorm:"column: user_id"`
}
  
func (p Product) TableName() string {
	return "product"
}


func Test_GORM_CRUD(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: ,
	})
	if err != nil {
		panic("failed to connect database")
	}

}