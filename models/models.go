package models

import (
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego/orm"
	// mysql driver used by beego orm
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB reference to ORM instance
var DB orm.Ormer

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	orm.RegisterModel(new(User), new(SocialAccount), new(XrplAccount), new(APIKey))

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/xrplns?charset=utf8", dbUser, dbPass), 30)

	DB = orm.NewOrm()
}
