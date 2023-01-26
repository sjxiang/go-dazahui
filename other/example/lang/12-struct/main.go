package main

import (
	"fmt"
)


type user struct {
	name     string
	password string
}


func main() {
	a := user{ name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha"))   // false
	fmt.Println(checkPassword2(&a, "haha")) // false
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password          
}


/*

创建 Person 类型的值，而不是说创建对象

        1. var tom Person   // ✅ 零值
           var tom *Person  // ❌ 不推荐

        2. tom := Person{}   // ❌
           tom := &Person{}  // ❌  e.g. return &Person{}，除外 

        3. tom := &Person{"Tom", 18}               // ❌ 不推荐，结构体布局顺序易混淆
           tom := Person{"Tom", 18}                // ❌ 不推荐
           tom := &Person{Name: "Tom", Age: "18"}  // ❌ 不推荐，混用值与指针
           tom := Person{Name: "Tom", Age: "18"}   // ✅ 字面量

        4. var tom *Person = new(person)  // ✅
        
        5. tom := struct { Name string }{ Name: "Tom"}  // ❌ 不推荐，匿名临时使用 e.g. JSON 拆分麻烦了              


Json Tag 标签
    
嵌套
    如果成员变量有指针类型，实例化时，注意空指针


*/