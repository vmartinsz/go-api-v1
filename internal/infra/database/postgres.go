package database

import (
	"fmt"
	"go-api-v1/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type GORMImpl struct {
	c *gorm.DB
}

func NewPGGORMImpl(config config.Config) (GORMImpl, error) {
	var err error

	gormimpl := GORMImpl{}

	DBURL := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Name,
		"disable",
		config.Database.Password,
	)

	gormimpl.c, err = gorm.Open("postgres", DBURL)

	if err != nil {
		log.Println(err)
		log.Fatal(err)

		return gormimpl, err
	}

	if gormimpl.c.DB().Ping() == nil {
		return gormimpl, nil
	}

	fatalErrorAboutDB := fmt.Errorf("can't start the database %s", err)

	log.Fatal(fatalErrorAboutDB)

	return gormimpl, fatalErrorAboutDB
}

func (pg GORMImpl) Client() *gorm.DB {
	return pg.c
}

func (pg GORMImpl) Close() error {
	return pg.Client().Close()
}
