package main

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelfirman/go-API/internal/application"
)

func main() {
	// env
	// ...

	// config
	// - addr
	addrCfg := ":8080"
	// - mysql
	mysqlCfg := mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "go_api_db",
		ParseTime: true,
	}
	// - cfg
	cfg := application.ConfigServer{Addr: addrCfg, MySQLDSN: mysqlCfg.FormatDSN()}

	// - server
	server := application.New(cfg)
	// - run
	if err := server.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
