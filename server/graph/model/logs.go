package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MileToKilometerRate is the constant rate of conversion between Miles and Kilometres
const MileToKilometerRate = 1.609344

func (d *DistanceValue) ConvertTo(t DistanceUnit) *DistanceValue {
	switch t {
	case DistanceUnitKilometre:
		if *d.Type != DistanceUnitKilometre {
			var converted = *d.Value * MileToKilometerRate
			var convertedType = DistanceUnitKilometre

			d.Value = &converted
			d.Type = &convertedType
		}
	case DistanceUnitMile:
		if *d.Type != DistanceUnitMile {
			var converted = *d.Value / MileToKilometerRate
			var convertedType = DistanceUnitMile

			d.Value = &converted
			d.Type = &convertedType
		}
	}

	return d
}

const GallonToLitreRate = 3.785411784

func (f *FluidValue) ConvertTo(t FluidUnit) *FluidValue {
	switch t {
	case FluidUnitLitre:
		if *f.Type != FluidUnitLitre {
			var converted = *f.Value * GallonToLitreRate
			var convertedType = FluidUnitLitre

			f.Value = &converted
			f.Type = &convertedType
		}
	case FluidUnitGallon:
		if *f.Type != FluidUnitGallon {
			var converted = *f.Value / GallonToLitreRate
			var convertedType = FluidUnitGallon

			f.Value = &converted
			f.Type = &convertedType
		}
	}

	return f
}

func (m *MoneyValue) ConvertTo(t MoneyUnit) *MoneyValue {
	return m // TODO this at some point

	switch t {
	case MoneyUnitCad:
		if *m.Type != MoneyUnitCad {
			var converted = *m.Value * 1
			var convertedType = MoneyUnitCad

			m.Value = &converted
			m.Type = &convertedType
		}
	case MoneyUnitUsd:
		if *m.Type != MoneyUnitUsd {
			var converted = *m.Value / 1
			var convertedType = MoneyUnitUsd

			m.Value = &converted
			m.Type = &convertedType
		}
	}

	return m
}

type Log interface {
	IsLog()
	GetDate() *time.Time
}

type FuelLog struct {
	BaseWithTimestamps
	Date       *time.Time     `json:"date"`
	VehicleID  *uuid.UUID     `json:"vehicle"`
	Odometer   *DistanceValue `json:"odometer" gorm:"embedded;embedded_prefix:odometer_"`
	Notes      *string        `json:"notes"`
	Trip       *DistanceValue `json:"trip" gorm:"embedded;embedded_prefix:trip_"`
	Grade      *int           `json:"grade"`
	FuelAmount *FluidValue    `json:"fuelAmount" gorm:"embedded;embedded_prefix:fuel_amount_"`
	FuelPrice  *MoneyValue    `json:"fuelPrice" gorm:"embedded;embedded_prefix:fuel_price_"`
}

// IsLog satisfies Log interface
func (FuelLog) IsLog() {}

// GetDate returns the log date
func (l FuelLog) GetDate() *time.Time { return l.Date }

type MaintenanceLog struct {
	BaseWithTimestamps
	Date      *time.Time     `json:"date"`
	VehicleID *uuid.UUID     `json:"vehicle"`
	Odometer  *DistanceValue `json:"odometer" gorm:"embedded;embedded_prefix:odometer_"`
	Notes     *string        `json:"notes"`
}

// IsLog satisfies Log interface
func (MaintenanceLog) IsLog() {}

// GetDate returns the log date
func (l MaintenanceLog) GetDate() *time.Time { return l.Date }

type OilChangeLog struct {
	BaseWithTimestamps
	Date      *time.Time     `json:"date"`
	VehicleID *uuid.UUID     `json:"vehicle"`
	Odometer  *DistanceValue `json:"odometer" gorm:"embedded;embedded_prefix:odometer_"`
	Notes     *string        `json:"notes"`
}

// IsLog satisfies Log interface
func (OilChangeLog) IsLog() {}

// GetDate returns the log date
func (l OilChangeLog) GetDate() *time.Time { return l.Date }
