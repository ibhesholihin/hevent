package db

import (
	"fmt"
	"log"
	"os"
	"time"

	md "github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Database(cfg *config.Config) (DB *gorm.DB, err error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})

	if err != nil {
		return
	}
	fmt.Println("Database Connected")

	return DB, err
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&md.Admin{},

		&md.Event{},
		&md.EventCategory{},
		&md.EventPriceTipe{},
		&md.EventSponsor{},

		&md.User{},

		&md.UserProfile{},
		&md.UserAddress{},

		//&md.CartSession{},
		//&md.CartItem{},
		//&md.Order{},
		//&md.OrderItem{},
		//&md.PaymentDetails{},
	)
}
