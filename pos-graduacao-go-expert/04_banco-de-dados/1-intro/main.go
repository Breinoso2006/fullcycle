package main

import (
	"database/sql"
	"fmt"

	// o _ é para importar o pacote sem usar, apenas para registrar o driver
	// o driver é registrado para que o sql.Open possa usar o driver para conexão com o MySQL
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	// Prepara a query para ser executada, evitando SQL Injection
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executa a query com os parâmetros passados. O _ pode ser substituído por algo que você queira receber de retorno
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	// Prepara a query para ser executada, evitando SQL Injection
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executa a query com os parâmetros passados. O _ pode ser substituído por algo que você queira receber de retorno
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	// Prepara a query para ser executada, evitando SQL Injection
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	// Poderia usar querycontext para passar um contexto de timeout por exemplo
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func main() {
	// Abre a conexão com o banco de dados, porém só valida os parâmetros de conexão quando a conexão é usada
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mydb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Nintendo Switch 2", 3999.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 3299.99
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f reais", p.Name, p.Price)

}
