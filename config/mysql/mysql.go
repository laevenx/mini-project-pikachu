package mysql

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// LOC localtime godoc
const LOC = "Local"

// Config godoc
type Config struct {
	Host string
	DB   string
	User string
	Pass string
	Port int
}

var DB *gorm.DB

// Connect godoc
func (conf Config) Connect() error {

	var err error

	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.DB, LOC,
	)
	DB, err = gorm.Open("mysql", uri)
	if err != nil {
		return err
	}

	if err = DB.DB().Ping(); err != nil {
		DB.Close()
	} else {
		log.Println("Database Has Connected ..........................!")
	}

	DB.DB().SetMaxIdleConns(0)
	return nil
}
