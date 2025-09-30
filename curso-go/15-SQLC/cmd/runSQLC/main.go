package main

import (
	"context"
	"database/sql"

	"github.com/lucasf1/15-SQLC/internal/db"

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
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx,
	// 	db.UpdateCategoryParams{
	// 		ID:          "74f263e6-def1-4fe9-ada3-a2e891b9cf6e",
	// 		Name:        "Backend Updated",
	// 		Description: sql.NullString{String: "Backend description updated"},
	// 	},
	// )
	// if err != nil {
	// 	panic(err)
	// }

	err = queries.DeleteCategory(ctx, "74f263e6-def1-4fe9-ada3-a2e891b9cf6e")
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
