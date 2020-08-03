package model

import uuid "github.com/satori/go.uuid"

type User struct {
	BaseWithTimestamps
	Email        string     `json:"email" gorm:"unique;not null"`
	PreferenceID *uuid.UUID `json:"preference"`
}

type UserPreference struct {
	Base
	Distance  *DistanceUnit `json:"distance"`
	Fluid     *FluidUnit    `json:"fluid"`
	Money     *MoneyUnit    `json:"money"`
	VehicleID *uuid.UUID    `json:"vehicle"`
}

type Vehicle struct {
	BaseWithTimestamps
	Make     string         `json:"make"`
	Model    string         `json:"model"`
	Year     int            `json:"year"`
	UserID   *uuid.UUID     `json:"user"`
	Odometer *DistanceValue `json:"odometer" gorm:"embedded;embedded_prefix:odometer_"`
}
