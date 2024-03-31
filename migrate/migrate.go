package main

import (
	"github.com/viviyee/go-crud/initializers"
	"github.com/viviyee/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
