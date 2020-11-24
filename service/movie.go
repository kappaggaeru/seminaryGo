package service

import (
	"github.com/kappaggaeru/seminaryGo/model"
	"github.com/kappaggaeru/seminaryGo/model/database"
)

// defines the behavior of the service
type Movie interface {
	Add(*model.Movie) *model.Movie
	FindByID(ID uint) *model.Movie
	Remove(ID uint)
	Update(*model.Movie) *model.Movie
}

type service struct {
	db *database.Database
}

// NewInstance of service
func NewInstance(db *database.Database) Movie {
	return &service{db}
}

// FindById
func (s *service) FindByID(ID uint) *model.Movie {
	return s.db.FindByID(ID)
}
// Add
func (s *service) Add(m *model.Movie) *model.Movie {
	s.db.Add(m)
	return m
}
// Remove
func (s *service) Remove(ID uint) {
	s.db.Remove(ID)
}
// Update
func (s *service) Update(m *model.Movie) *model.Movie {
	s.db.Update(m)
	return m
}
