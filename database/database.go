package database

import (
	"fmt"
	"os"

	//"gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	databaseUrl := os.Getenv("DATABASE_URL")

	//	USING MYSQL
	//const MYSQL = "root:@tcp(127.0.0.1:3306)/go_fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := MYSQL
	//DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//	USING POSTGRESQL
	//const POSTGRESQL = "postgresql://postgres:@localhost:5432/go_fiber_gorm?sslmode=disable&TimeZone=Asia/Jakarta"
	if databaseUrl == "" {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbname := os.Getenv("DB_NAME")
		if err != nil {
			panic(err)
		}
		if username == "" {
			username = "postgres"
		}
		if password == "" {
			password = ""
		}
		if host == "" {
			host = "localhost"
		}
		if port == "" {
			port = "5432"
		}
		if dbname == "" {
			dbname = "miniproject"
		}
		databaseUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Jakarta", username, password, host, port, dbname)
	}

	dsn := databaseUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")
}
