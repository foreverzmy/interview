package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

// DB Gorm database engine
var DB *gorm.DB

func init() {
	var err error
	// 初始化数据库连接
	DB, err = gorm.Open("mysql", "root:zmyisno1@tcp(0.0.0.0:3306)/interview_qu?charset=utf8mb4&parseTime=true")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
