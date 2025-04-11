package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Movie struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model // Adiciona os campos CreatedAt, UpdatedAt e DeletedAt
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// O migrate serve para criar a tabela no banco de dados caso ela não exista
	// e para atualizar a tabela caso ela já exista
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Category{})

	// Comandos para inserir dados no banco de dados

	// db.Create(&Movie{Name: "Inception", Price: 10.99})
	// db.CreateInBatches([]Movie{
	// 	{Name: "The Matrix", Price: 8.99},
	// 	{Name: "Interstellar", Price: 12.99},
	// 	{Name: "Lilo e Stitch", Price: 9.99},
	// }, 2)

	// Comandos para buscar dados no banco de dados

	// var movie Movie
	// First neste caso é o mesmo que Select * from movies where id = 1
	// db.First(&movie, 6)
	// fmt.Println(movie)
	// db.First(&movie, "name = ?", "Lilo e Stitch")
	// fmt.Println(movie)

	// var movies []Movie
	// db.Find(&movies)
	// db.Limit(2).Find(&movies)
	// db.Limit(2).Offset(2).Find(&movies)
	// db.Where("ID > ?", 2).Find(&movies)
	// db.Where("name LIKE ?", "%cep%").Find(&movies)
	// for _, movie := range movies {
	// 	fmt.Println(movie)
	// }

	// Comandos para atualizar dados no banco de dados

	// // ao mudar o nome do filme, o gorm atualiza o ID do filme
	// var movie2 Movie
	// db.First(&movie2, 2)
	// movie2.Name = "Up!"
	// db.Save(&movie2)

	// comandos para deletar dados no banco de dados

	// o gorm não deleta o registro do banco de dados, ele apenas marca como deletado (soft delete)
	// var movie3 Movie
	// db.First(&movie3, 1)
	// fmt.Println(movie3.Name)
	// db.Delete(&movie3)

	// Relacionamento entre tabelas

	// category := Category{Name: "Aventura"}
	// db.Create(&category)
	// product := Movie{Name: "Uncharted", Price: 10.99, CategoryID: category.ID}
	// db.Create(&product)

	var movies []Movie
	// Preload é usado para carregar os dados relacionados de uma tabela
	db.Preload("Category").Find(&movies)
	for _, movie := range movies {
		println(movie.Name, movie.Category.Name)
	}


}
