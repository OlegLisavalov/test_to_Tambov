package main

import (
	"test_Task_New_server/config"
	"test_Task_New_server/database"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)
}