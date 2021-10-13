package testing

import (
	"github.com/rest-api-market/connection"
)

//Testing Insertions
func Insert() {
	db := connection.GetConnection()

	// category1 := model.Category{Name: "Linea blanca", Description: "productos de linea blanca"}
	// category2 := model.Category{Name: "Muebleria", Description: "Productos de muebleria"}
	// result := db.Model(&model.Category{}).Create(&category1)
	// result1 := db.Model(&model.Category{}).Create(&category2)
	// log.Print(result, result1)

	// producto1 := &model.Product{Name: "Lavadora", Price: 3000, Stock: 20, CategoryID: 2}
	// producto2 := &model.Product{Name: "Secadora", Price: 2000, Stock: 20, CategoryID: 2}

	// producto3 := &model.Product{Name: "Microondas", Price: 2000, Stock: 10, CategoryID: 1}
	// producto4 := &model.Product{Name: "Plancha", Price: 2000, Stock: 20, CategoryID: 1}

	// producto5 := &model.Product{Name: "Escoba", Price: 2000, Stock: 10, CategoryID: 3}
	// producto6 := &model.Product{Name: "Franela", Price: 2000, Stock: 20, CategoryID: 3}

	// db.Model(&model.Product{}).Create(&producto1)
	// db.Model(&model.Product{}).Create(&producto2)
	// db.Model(&model.Product{}).Create(&producto3)
	// db.Model(&model.Product{}).Create(&producto4)
	// db.Model(&model.Product{}).Create(&producto5)
	// db.Model(&model.Product{}).Create(&producto6)

	// customer1 := &model.Customer{Name: "Jair", Lastname: "Vasquez"}
	// customer2 := &model.Customer{Name: "José", Lastname: "Rodriguez"}
	// customer3 := &model.Customer{Name: "David", Lastname: "Martinez"}

	// db.Model(&model.Customer{}).Create(customer1)
	// db.Model(&model.Customer{}).Create(customer2)
	// db.Model(&model.Customer{}).Create(customer3)

	// compra1 := &model.CustomerProduct{CustomerID: 1, ProductID: 1}
	// compra2 := &model.CustomerProduct{CustomerID: 1, ProductID: 2}
	// compra3 := &model.CustomerProduct{CustomerID: 1, ProductID: 3}

	// compra4 := &model.CustomerProduct{CustomerID: 2, ProductID: 1}
	// compra5 := &model.CustomerProduct{CustomerID: 2, ProductID: 3}
	// compra6 := &model.CustomerProduct{CustomerID: 3, ProductID: 3}

	// db.Model(&model.CustomerProduct{}).Create(compra1)
	// db.Model(&model.CustomerProduct{}).Create(compra2)
	// db.Model(&model.CustomerProduct{}).Create(compra3)
	// db.Model(&model.CustomerProduct{}).Create(compra4)
	// db.Model(&model.CustomerProduct{}).Create(compra5)
	// db.Model(&model.CustomerProduct{}).Create(compra6)

	defer db.Close()
}
