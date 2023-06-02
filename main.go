package main

import (
	"dbtest/db"
	"fmt"
)

func main() {
	db.MigrateModels()
	fmt.Println("success")
}
