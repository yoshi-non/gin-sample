package config

import (
	"fmt"
	"gintut/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DB接続関数
func ConnectDB() {
	err := godotenv.Load("./.env")
	checkErr(err)

	//ここは自分の環境に合った設定で行う
	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DATABASE_ADDRESS"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	checkErr(err)

	//DBのマイグレーション
	db.AutoMigrate(&models.Memo{})

	DB = db
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
