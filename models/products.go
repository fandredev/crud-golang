package products

import (
	"store/db"
)

type Products struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func FindAllProductsFromDatabase() []Products {
	database := db.ConnectWithDatabase()

	selectAllProductsFromDatabase, errorSelectedProducts := database.Query("select * from products order by id asc")

	if errorSelectedProducts != nil {
		panic(errorSelectedProducts.Error())
	}

	product := Products{}
	products := []Products{}

	for selectAllProductsFromDatabase.Next() {
		var id, quantity int
		var name, description string
		var price float64

		errorMappedDataFromDatabase := selectAllProductsFromDatabase.Scan(&id, &name, &description, &price, &quantity)

		if errorMappedDataFromDatabase != nil {
			panic(errorMappedDataFromDatabase.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer database.Close()

	return products
}

func FindProductById(id string) Products {
	database := db.ConnectWithDatabase()

	findProductFromBankById, errFindProduct := database.Query("select * from products where id=$1", id)

	if errFindProduct != nil {
		panic(errFindProduct.Error())
	}

	productFiltered := Products{}

	for findProductFromBankById.Next() {
		var id, quantity int
		var name, description string
		var price float64

		errorScanLinesFromDatabase := findProductFromBankById.Scan(&id, &name, &description, &price, &quantity)

		if errorScanLinesFromDatabase != nil {
			panic(errorScanLinesFromDatabase.Error())
		}

		productFiltered.Id = id
		productFiltered.Name = name
		productFiltered.Quantity = quantity
		productFiltered.Description = description
		productFiltered.Price = price
	}

	defer database.Close()
	return productFiltered
}

func CreateNewProductFromDatabase(name, description string, price float64, quantity int) {
	database := db.ConnectWithDatabase()

	insertProductInDatabase, errInsertProductInDatabase := database.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if errInsertProductInDatabase != nil {
		panic(errInsertProductInDatabase.Error())
	}

	insertProductInDatabase.Exec(name, description, price, quantity)

	defer database.Close()
}

func DeleteProductFromDatabase(id string) {
	database := db.ConnectWithDatabase()

	deleteProductFromDatabase, errorDeleteProductFromDatabase := database.Prepare("delete from products where id=$1")

	if errorDeleteProductFromDatabase != nil {
		panic(errorDeleteProductFromDatabase.Error())
	}

	deleteProductFromDatabase.Exec(id)

	defer database.Close()
}

func UpdateProduct(id, quantity int, name, description string, price float64) {
	database := db.ConnectWithDatabase()

	updateProduct, errorUpdateProduct := database.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if errorUpdateProduct != nil {
		panic(errorUpdateProduct.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer database.Close()
}
