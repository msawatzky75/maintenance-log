package model

type User struct {
	Base
	Email        string `json:"email" pg:",unique"`
	PreferenceID string `json:"preference"`
}

type UserPreference struct {
	Distance  *DistanceUnit `json:"distance"`
	Fluid     *FluidUnit    `json:"fluid"`
	Money     *MoneyUnit    `json:"money"`
	VehicleID string        `json:"vehicle"`
}
