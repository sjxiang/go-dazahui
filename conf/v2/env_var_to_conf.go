package conf_v2

import (
	"log"

	"github.com/joho/godotenv"  
)


// 加载环境变量
func Init() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal(err)
	}
}


/*

注意事项
	.env 配置文件要和 main.main 在同一个目录下

*/