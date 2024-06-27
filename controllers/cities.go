package controllers

import (
	"cities/models"
	"cities/pb/cities"
	"context"
	"database/sql"
)

// City struct
type City struct {
	DB *sql.DB
}

// FUNCTION
// function GetCity
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}

// function Create
func (s *City) Create(ctx context.Context, in *cities.CityInput) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Create(ctx, s.DB, in)
	return &cityModel.Pb, err
}
