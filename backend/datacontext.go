package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	//Esta función realiza la conexión con la base de datos
	//utilizando el ORM GORM.
	//La base de datos se ejecuta en el localhost:1433
	//y se llama Manhattam
	dsn := "sqlserver://localhost:1433?database=Manhattam"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	return db, err
}
