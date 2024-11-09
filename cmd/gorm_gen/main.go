package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := "host=remote user=weedien password=031209 dbname=wespace search_path=link_single port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	gormdb, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `entity.User` following conventions
	// g.ApplyBasic(entity.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `entity.User` and `entity.Company`
	// g.ApplyInterface(func(Querier) {}, entity.User{}, entity.Company{})

	// Generate the code
	g.Execute()
}
