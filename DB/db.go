package DB

import (
	"fmt"
	"github.com/fahimimam/UserStore/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const (
	host           = "localhost"
	user           = "postgres"
	password       = "changeme"
	dbName         = "postgres"
	port           = 5432
	runAutoMigrate = false
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could Not connect to Database, ", err)
	}
	if runAutoMigrate {
		db.AutoMigrate(models.User{})
	}
	DB = db

	return DB
}

func Get() *gorm.DB {
	if DB == nil {
		return ConnectDB()
	}
	return DB
}

func AutoMigrateDB(c *gin.Context) {
	log.Println("Starting Migrator")
	Get().AutoMigrate(models.User{}, models.Tags{})
	c.JSON(http.StatusOK, gin.H{
		"message": "Migrated successfully",
	})
}
