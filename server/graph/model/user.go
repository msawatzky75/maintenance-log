package model

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseWithTimestamps
	Email        string     `json:"email" gorm:"unique;not null"`
	Password     string     `json:"-"`
	PreferenceID *uuid.UUID `json:"preference"`
}

// Error allows for singleton errors
type Error string

func (err Error) Error() string {
	return string(err)
}

const ErrInvalidEmail = Error("invalid email")
const ErrInvalidPassword = Error("invalid password")
const ErrEmailExists = Error("email already exists")
const ErrUnableToHashPassword = Error("unable to hash password")

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

func validEmail(e string) bool {
	r := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return r.MatchString(e)
}

func validPassword(p string) bool {
	// TODO
	return len(p) >= 8
}

func CreateUser(db *gorm.DB, email string, password string) (User, error) {
	var u User

	if !validEmail(email) {
		return u, ErrInvalidEmail
	}
	if !validPassword(password) {
		return u, ErrInvalidPassword
	}

	var c int
	db.Model(User{}).Where("email = ?", email).Count(&c)
	if c != 0 {
		return u, ErrEmailExists
	}

	cost, err := strconv.Atoi(os.Getenv("PASSWORD_HASH_COST"))
	if err != nil {
		log.Error(err)
		return u, fmt.Errorf("unable to set password hash cost")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Error(err)
		return u, ErrUnableToHashPassword
	}

	tx := db.Begin()
	// Rollback the transaction if the routine panics
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	p := UserPreference{
		Distance:  nil,
		Fluid:     nil,
		Money:     nil,
		VehicleID: nil,
	}
	err = tx.Save(&p).Error
	if err != nil {
		log.Error(err)
		return u, err
	}

	u = User{
		Email:        email,
		Password:     string(hash),
		PreferenceID: &p.ID,
	}
	err = tx.Save(&u).Error
	if err != nil {
		log.Error(err)
		return u, err
	}

	err = tx.Commit().Error
	if err != nil {
		log.Error(err)
		return u, err
	}

	return u, nil
}
