package config

import (
	"context"
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDbUtil struct {
	dsn    string
	dbName string
	ctx    context.Context
}

func NewPostgresDbUtil() *PostgresDbUtil {
	return &PostgresDbUtil{
		dsn:    "host=localhost user=gis password=rahasia2023 dbname=sdb port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		dbName: "sdb",
		ctx:    nil,
	}
}

func (d PostgresDbUtil) Connect() (client *gorm.DB, err error) {
	fmt.Println(aurora.Cyan(d.dsn))
	db, err := gorm.Open(postgres.Open(d.dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
