package main

import (
	"final-project/helpers"
	_ "final-project/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	helpers.ConnectDatabase()

	beego.Run()

}

// func GetConnectDb() *gorm.DB {
// 	dbHost, _ := beego.AppConfig.String("Database::dbHost")
// 	dbUser, _ := beego.AppConfig.String("Database::dbUser")
// 	dbPass, _ := beego.AppConfig.String("Database::dbPass")
// 	dbName, _ := beego.AppConfig.String("Database::dbName")
// 	dbPort, _ := beego.AppConfig.String("Database::dbPort")
// 	conn, err := helpers.ConnectDatabase(dbHost, dbUser, dbPass, dbName, dbPort)
// 	sqlDb, err := conn.DB()
// 	sqlDb.SetConnMaxIdleTime(10)
// 	sqlDb.SetConnMaxLifetime(100)
// 	if err != nil {
// 		fmt.Println("cannot connect to db")
// 	}
// 	return conn
// }
