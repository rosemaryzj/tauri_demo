package lite

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reminder/app/entry"
	"time"
)

func NewSqlLite() *gorm.DB {
	pool, err := gorm.Open(sqlite.Open("robot.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	_ = pool.AutoMigrate(new(entry.Robot), new(entry.User))
	if err != nil {
		panic("db init failed")
	}
	sqlDB, _ := pool.DB()
	sqlDB.SetMaxIdleConns(4)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(20 * time.Minute)
	return pool
}
