package db_conn

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	"fmt"
)




func DBConnect(pg_conn string)(db *gorm.DB,err error) {
	db, err = gorm.Open(postgres.Open(pg_conn), &gorm.Config{})
	if err != nil {
		fmt.Printf("error connecting %v", err)
		return nil,err
	}
	//defer conn.Close(context.Background())
	return db,nil;
}

