package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectionDatabase(){
	server := "DESKTOP-F8RF6F9"
	port := 1433
	user := "user"
	password := "fajarbaru789"
	database  := "kerjaku_db"

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		user, password, server, port, database)
		
		var err error
		DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Gagal koneksi ke database:", err)
		}
	
		fmt.Println("Berhasil terkoneksi ke database!")	
		// DB.Migrator().DropTable(  &models.Profile{})

		// DB.AutoMigrate( &models.Profile{})

	}