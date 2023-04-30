package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Age       int       `gorm:"not null"`
	UpdatedAt time.Time `gorm:"updated_at"`
	CreatedAt time.Time `gorm:"created_at"`
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=elastic port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		user := User{
			Name:  faker.Name(),
			Email: faker.Email(),
			Age:   rand.Intn(80) + 20,
		}
		db.Create(&user)
	}

	fmt.Println("Data created successfully!")
}
