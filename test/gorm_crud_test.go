package test

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql" // GORM 通过驱动来连接数据库
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&&multiStatements=true&loc=Local",
					"root",
					"123456",
					"172.20.0.1",
					3306,
					"crud-demo")

var (
	db *gorm.DB
	err error
)


func Init() {
		
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database, error: " + err.Error())
	}

	db.AutoMigrate(&User{})  // 自动迁移

}


// 创建数据
func Test_GORM_Create(t *testing.T) {
	Init()

	// 创建一条
	user := &User{Name: "Jie", Age: 28} 
	res := db.Create(user)

	t.Log(res.RowsAffected) 
	t.Log(res.Error)         // 获取 err
	t.Log(user.ID)           // 返回被插入数据的主键
	
	
	// 创建多条
	users := []*User{{Name: "22", Age: 15}, {Name: "33", Age: 16}}
	res = db.Create(users)

	t.Log(res.Error)
	for _, u := range users {
		t.Log(u.ID)
	}

}


// 查询数据
func Test_GORM_Read(t *testing.T) {
	Init()

	
}



// /*
// 	// 查询数据 - 读取
// 	var product Product
// 	db.First(&product, "code = ?", "101")  // 查找 code 字段值为 101 的记录
			                               
// 	/* 
	
// 		First 查 1 条记录，此外存在还有些问题
// 		Find 查多条
	
// 	*/


// 	// 更新数据
// 	db.Model(&product).Update("Price", 2)  // 将 product 的 price 更新为 2

	
// 	// 更新数据 - 多个字段
// 	db.Model(&product).Updates(Product{Code: "101", Price: 3})   // 仅更新非零值字段，不大行，int 零值为 0，字符串零值为 ""
// 	db.Model(&product).Updates(map[string]interface{}{"Code": "101", "Price": 4})  // map 可以避免零值更新问题
	
// 	/*

	// 0、''、false 等零值是不会保存到数据库
// 	*/

// 	// 删除 product
// 	db.Where("id = ?", 1).Delete(&product)


// */
