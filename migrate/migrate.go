package main

import (
	"fmt"

	"github.com/lordofthemind/jsonCrudGorm/initializers"
	"github.com/lordofthemind/jsonCrudGorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgreSQL()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	fmt.Println("Database Migrated")
}
