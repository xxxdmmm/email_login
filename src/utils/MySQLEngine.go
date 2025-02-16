package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

type MySQL struct {
	Username  string
	Password  string
	IpAddress string
	Port      int
	DBName    string
	Charset   string
	Engine    *xorm.Engine
}

func (m *MySQL) connectDB() (*xorm.Engine, error) {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		m.Username,
		m.Password,
		m.IpAddress,
		m.Port,
		m.DBName,
		m.Charset)

	engine, err := xorm.NewEngine("mysql", connectInfo)

	if err == nil {
		m.Engine = engine
	}

	return engine, err
}

func (m *MySQL) Migrate(model ...interface{}) {
	err := m.Engine.Sync2(model...)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func UseDefaultDB() *MySQL {

	var MyDB *MySQL = &MySQL{
		Username:  "root",
		Password:  "",
		IpAddress: "127.0.0.1",
		Port:      3306,
		DBName:    "go",
		Charset:   "utf8mb4",
	}

	_, err := MyDB.connectDB()

	if err != nil {
		log.Fatal("Error:", err)
	}

	return MyDB
}

var DefaultEngine = UseDefaultDB()
