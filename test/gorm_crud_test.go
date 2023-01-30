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
		
	db, err = gorm.Open(mysql.Open(dsn), 
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,  // 关闭默认事务
			PrepareStmt: true},            // 缓存预编译语句
	)
	if err != nil {
		panic("failed to connect database, error: " + err.Error())
	}

	db.AutoMigrate(&User{}, &Email{})  // 自动迁移

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

	// 获取第一条记录（主键升序），查询不到数据，则返回 ErrRecordNotFound
	u := &User{}
	db.First(u, "age = ?", 100)  // SELECT * FROM `user` WHERE age = 100 ORDER BY `user`.`id` LIMIT 1
	

	// 查询多条数据
	users := make([]*User, 0)
	result := db.Where("age > 10").Find(&users)  // SELECT * FROM `user` WHERE age > 10

	t.Log(result.RowsAffected)  // 返回找到的记录数，相当于 `len(users)`
	t.Log(result.Error)


	// SELECT * FROM `user` WHERE `user`.`nick_name` = 'Jie'
	db.Where(&User{Name: "Jie", Age: 0}).Find(&users)
	// SELECT * FROM `user` WHERE `age` = 0 AND `nick_name` = 'Jie'
	db.Where(map[string]interface{}{"nick_name": "Jie", "age": 0}).Find(&users)
	
}

/*

	First 的使用踩坑
		使用 First 时，需要注意查询不到数据，会返回 ErrRecordNotFound；
		使用 Find 查询多条数据，查询不到数据不会返回错误。

	
	使用结构体作为查询条件
		当使用 struct 作为查询条件时，GORM 只会查询非零值字段；
		这意味着如果您的字段值为 0、""、false 或其他零值，该字段不会被用于构建查询条件。

		建议使用 map 来构建查询条件

*/


// 更新数据
func Test_GORM_Update(t *testing.T) {
	Init()

	db.Model(&User{ID: 1}).Where("nick_name = ?", "Jie").Update("age = ?", 28)

	db.Model(&User{ID: 1}).Select("age", "nick_name").Updates(map[string]interface{}{"nick_name": "Xiang", "age": 0})

}

/*

	同上，建议使用 map 更新或者使用 Select 选择字段

*/



// 更新数据
func Test_GORM_Delete(t *testing.T) {
	Init()

	// 删除一条
	db.Delete(&User{ID: 1})  
	
	/*
		UPDATE `user` SET `delete_at`='2023-01-30 17:38:39.107' \ 
						WHERE `user`.`id` = 1 AND `user`.`delete_at` IS NULL;
		软删除
		NULL，代表某个不知道的值，并不是为空 
		
	*/
	
	// 批量删除
	db.Where("age = ?", 15).Delete(&User{})

	

	users := make([]*User, 0)
	// 查询时，会忽略被软删除的记录
	db.Where("age = 15").Find(&users) 
	t.Log(len(users))

	// 不忽略
	db.Unscoped().Where("age = 15").Find(&users)
	t.Log(len(users))
}	


/*

	- 物理删除
	- 软删除
	
	GORM 提供了 gorm.DeletedAt 用于帮助用户实现软删除

	拥有软删除能力的 Model 调用 Delete 时，记录不会被从数据库中真正删除。
	但 GORM 会将 DeleteAt 置为当前时间，并且不能再通过正常的查询方法找到该记录。

	使用 Unscoped 可以查询到被软删的数据

*/


// 事务 - 手动
func Test_GORM_Transaction(t *testing.T) {
	Init()

	// 开始事务
	tx := db.Begin()  
	// 在事务中执行一些 db 操作（从这里开始，应该使用 `tx`，而不是 `db`）
	if err := tx.Create(&User{Name: "Test_1"}).Error; err != nil {
		tx.Rollback()
		// 遇到错误时，回滚事务
		return
	}
	if err := tx.Create(&User{Name: "Test_2"}).Error; err != nil {
		tx.Rollback()

		return
	}

	// 提交事务
	tx.Commit() 
}


// 事务 - 自动
func Test_GORM_Transaction_Auto(t *testing.T) {
	Init()

	if err := db.Transaction(func(tx *gorm.DB) error {  // 自动提交事务，避免漏写 Commit 、Rollback
		if err := tx.Create(&User{Name: "Test_1"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&User{Name: "Test_2"}).Error; err != nil {
			tx.Rollback()
			return err
		}
		
		return nil
	}); err != nil {
		return
	}
}

