package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/17-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend development description", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "bfaeb9d4-05cf-4d55-bfa9-8c2362b74743",
	// 	Name:        "Python",
	// 	Description: sql.NullString{String: "Python development description", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// err = queries.DeleteCategory(ctx, "bfaeb9d4-05cf-4d55-bfa9-8c2362b74743")
	// if err != nil {
	// 	panic(err)
	// }

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}

}
