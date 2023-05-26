package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/gulimall_product/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:root@(127.0.0.1:3306)/gulimall_pms?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyBasic(
		g.GenerateModel("pms_category"),
	)

	// Generate the code
	g.Execute()
}
