package database

import (
	"github.com/kappaggaeru/seminaryGo/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database
type Database struct {
	db *gorm.DB
}

// NewDatabase
func NewDatabase(dsn string) *Database { //dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Database{db}
}

// CreateSchema generates entitys
func (database *Database) CreateSchema() {
	database.db.AutoMigrate(&model.Movie{})
}

// FindAll returns all results
func (database *Database) FindAll() []model.Movie {
	var movies []model.Movie
	database.db.Find(&movies)
	return movies
}
// FindByID
func (database *Database) FindByID(ID uint) *model.Movie {
	var m model.Movie
	database.db.First(&m, ID)
	return &m
}
// Add
func (database *Database) Add(m *model.Movie) {
	database.db.Create(m)
}
// Remove
func (database *Database) Remove(i uint) {
	var m model.Movie
	database.db.Unscoped().Delete(&m, i)
}
// Update
func (database *Database) Update(m *model.Movie){
	database.db.Save(m)
}