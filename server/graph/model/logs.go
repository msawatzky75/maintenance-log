package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type FuelLog struct {
	BaseWithTimestamps
	Date       *time.Time     `json:"date"`
	VehicleID  *uuid.UUID     `json:"vehicle"`
	Odometer   *DistanceValue `json:"odometer" gorm:"embedded;embedded_prefix:odometer_"`
	Notes      string         `json:"notes"`
	Trip       *DistanceValue `json:"trip"`
	Grade      *int           `json:"grade"`
	FuelAmount *FluidValue    `json:"fuelAmount" gorm:"embedded;embedded_prefix:fuel_amount_"`
	FuelPrice  *MoneyValue    `json:"fuelPrice" gorm:"embedded;embedded_prefix:fuel_price_"`
}

// IsLog satisfies Log interface
func (FuelLog) IsLog() {}

type MaintenanceLog struct {
	BaseWithTimestamps
	Date      *time.Time     `json:"date"`
	VehicleID *uuid.UUID     `json:"vehicle"`
	Odometer  *DistanceValue `json:"odometer" gorm:"embedded;embedded_prefix:odometer_"`
	Notes     string         `json:"notes"`
}

// IsLog satisfies Log interface
func (MaintenanceLog) IsLog() {}

type OilChangeLog struct {
	BaseWithTimestamps
	Date      *time.Time     `json:"date"`
	VehicleID *uuid.UUID     `json:"vehicle"`
	Odometer  *DistanceValue `json:"odometer" gorm:"embedded;embedded_prefix:odometer_"`
	Notes     string         `json:"notes"`
}

// IsLog satisfies Log interface
func (OilChangeLog) IsLog() {}
