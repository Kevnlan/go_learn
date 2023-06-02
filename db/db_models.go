package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	EmpNo      string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	ModifiedBy time.Time `gorm:"not null"`

	PersonalInfo PersonalInformation `gorm:"foreignKey:EmpNo"`
}

type Roles struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Status string `gorm:"not null"`

	PersonalInfo []PersonalInformation `gorm:"foreignKey:RoleID"`
}

type PersonalInformation struct {
	gorm.Model
	EmpNo         string `gorm:"not null"`
	Gender        string `gorm:"not null"`
	Surname       string `gorm:"not null"`
	FirstName     string `gorm:"not null"`
	MiddleName    string
	Dob           time.Time `gorm:"not null"`
	RoleID        uint      `gorm:"not null"`
	Town          string
	Country       string
	Location      string
	PhoneNumber   string `gorm:"not null"`
	MaritalStatus bool   `gorm:"not null"`
	ImageURL      string `gorm:"not null"`
	IsInsured     bool   `gorm:"not null"`
	Status        string

	Role                  Roles                    `gorm:"foreignKey:RoleID"`
	EmployeeLeavesRequest []EmployeeLeavesRequests `gorm:"foreignKey:EmpNo"`
	EmployeeLeave         []EmployeeLeaves         `gorm:"foreignKey:EmpNo"`
}

type PersonalData struct {
	gorm.Model
	EmpNo            string    `gorm:"not null"`
	TeamID           uint      `gorm:"not null"`
	DepartmentID     uint      `gorm:"not null"`
	DateEmployed     time.Time `gorm:"not null"`
	DateLeft         time.Time `gorm:"not null"`
	ReportingManager string    `gorm:"not null"`
	JobTitle         string    `gorm:"not null"`
	Status           string    `gorm:"not null"`

	PersonalInfo PersonalInformation `gorm:"foreignKey:EmpNo"`
	Team         Teams               `gorm:"foreignKey:TeamID"`
	Department   Departments         `gorm:"foreignKey:DepartmentID"`
}

type LeaveTypes struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Code         string `gorm:"not null"`
	NumberInDays int    `gorm:"not null"`
	Description  string `gorm:"not null"`
	Status       string `gorm:"not null"`

	EmployeeLeavesRequest []EmployeeLeavesRequests `gorm:"foreignKey:LeaveTypeId"`
	EmployeeLeave         []EmployeeLeaves         `gorm:"foreignKey:LeaveTypeId"`
}

type EmployeeLeaves struct {
	gorm.Model
	EmpNo           string `gorm:"not null"`
	LeaveTypeId     uint   `gorm:"not null"`
	LeaveLimit      string `gorm:"not null"`
	AllocatedLeaves string `gorm:"not null"`
	Status          string `gorm:"not null"`

	// PersonalInfo PersonalInformation `gorm:"foreignKey:EmpNo"`
	LeaveType    LeaveTypes          `gorm:"foreignKey:LeaveTypeId"`
}

type EmployeeLeavesRequests struct {
	gorm.Model
	EmpNo         string    `gorm:"not null"`
	Reason        string    `gorm:"not null"`
	LeaveTypeId   string    `gorm:"not null"`
	DaysRemaining int       `gorm:"not null"`
	LeaveLimit    string    `gorm:"not null"`
	FromDate      time.Time `gorm:"not null"`
	ToDate        time.Time `gorm:"not null"`
	LeaveStatus   string    `gorm:"not null"`
	Status        string    `gorm:"not null"`
	DocumentUrl   string    `gorm:"not null"`

	// PersonalInfo PersonalInformation `gorm:"foreignKey:EmpNo"`
	LeaveType    LeaveTypes          `gorm:"foreignKey:LeaveTypeId"`
}

type Teams struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      string `gorm:"not null"`

	PersonData []PersonalData `gorm:"foreignKey:TeamID"`
}

type Departments struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      string `gorm:"not null"`

	PersonData []PersonalData `gorm:"foreignKey:DepartmentID"`
}

type Holidays struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	Year        time.Time `gorm:"not null"`
	Description string    `gorm:"not null"`
	Status      string    `gorm:"not null"`
}

func MigrateModels() {
	dbName := "value8"
	dbUser := "value8"
	dbPassword := ""
	dbHost := "localhost"
	dbPort := "5432"

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connStr,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&Users{},
		&PersonalData{},
		&Roles{},
		&Teams{},
		&Departments{},
		&Holidays{},
		&PersonalInformation{},
		&LeaveTypes{},
		&EmployeeLeaves{},
		&EmployeeLeavesRequests{},
	)

	if err != nil {
		fmt.Println("Error migrating the models:", err)
		return
	}

	fmt.Println("Database migration completed successfully!")
}
