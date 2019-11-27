package main

import (
	"github.com/jinzhu/gorm"
	"github.com/noaway/godao"
	"github.com/noaway/godao/example/relation"
)

func main() {
	godao.InitORM(godao.PostgreSQLConfig{})

	// how to use orm, see relation/relation.go
	// how to test, see relation/relation_test.go

	// how to use transaction
	godao.Transact(godao.Engine, func(db *gorm.DB) error {
		err := db.Create(relation.Relation{}).Error
		if err != nil {
			return err
		}
		return db.Create(relation.Relation{}).Error
	})
}
