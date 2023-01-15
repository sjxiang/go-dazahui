package conf_v1

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

// Configuration 项目配置
type Configuration struct {
	// 私钥
	SecretKey      string        `json:"secret_key"`
	// 是否测试环境
	TestEnv        bool          `json:"test_env"`
	// 会话超时时间
	SessionTimeout time.Duration `json:"session_timeout"`
}


var config *Configuration
var once sync.Once


// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{
			SessionTimeout: 1,
		}
		
		fd, err := os.Open("config.dev.json")
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		defer fd.Close()

		encoder := json.NewDecoder(fd)
		err = encoder.Decode(config)
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}
	})

	return config
}

