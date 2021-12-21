package main

import (
	"fmt"

	"github.com/jolinGalal/jumia/internal/models"
	"github.com/jolinGalal/jumia/pkg/database"
)

func main() {
	database.New().Folder("../sample.db").GetInstance()
	var c = models.Customer{ID: 1, Name: "hello", Phone: "(212) 698054317"}
	fmt.Println(c.GetCountry())
	fmt.Println(c.GetState())

}
