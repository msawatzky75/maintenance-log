package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type FuelLog struct {
	BaseWithTimestamps
	Date       *time.Time     `json:"date"`
	VehicleID  *uuid.UUID     `json:"vehicle"`
	Notes      string         `json:"notes"`
	Trip       *DistanceValue `json:"trip"`
	Grade      *int           `json:"grade"`
	FuelAmount *FluidValue    `json:"fuelAmount"`
	FuelPrice  *MoneyValue    `json:"fuelPrice"`
}

// IsLog satisfies Log interface
func (FuelLog) IsLog() {}

type MaintenanceLog struct {
	BaseWithTimestamps
	Date      *time.Time `json:"date"`
	VehicleID *uuid.UUID `json:"vehicle"`
	Notes     string     `json:"notes"`
}

// IsLog satisfies Log interface
func (MaintenanceLog) IsLog() {}

type OilChangeLog struct {
	BaseWithTimestamps
	Date      *time.Time `json:"date"`
	VehicleID *uuid.UUID `json:"vehicle"`
	Notes     string     `json:"notes"`
}

// IsLog satisfies Log interface
func (OilChangeLog) IsLog() {}
