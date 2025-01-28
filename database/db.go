package database

import(
	"database/sql"
	"fmt"
	"log"
	"time"

	"test_Task_New_server/config"

	_ "github.com/lib/pq"  
)

var DB *sql.DB

func ConnectDB(cfg *config.Config) {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	DB.SetMaxOpenConns(10)              
	DB.SetMaxIdleConns(5)          
	DB.SetConnMaxLifetime(30 * time.Minute) 

	
	if err := DB.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	log.Println("Успешное подключение к базе данных!")
}