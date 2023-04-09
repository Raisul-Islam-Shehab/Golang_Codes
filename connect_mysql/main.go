package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Employee struct {
	ID         int
	Name       string
	Phone      string
	Email      string
	BloodGroup string
}

func main() {
	var err error
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "employees",
		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")

	employees, err := getEmployeesByName("Shehab")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(employees)

	id, err := addEmployee(Employee{
		Name:       "Shehab",
		Phone:      "543456",
		Email:      "add@yahoo.com",
		BloodGroup: "A-",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID of the added employee is", id)

	employees, err = getEmployeesByName("Shehab")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(employees)
}

func getEmployeesByName(name string) ([]Employee, error) {
	var employees []Employee
	rows, err := db.Query("SELECT * FROM employee WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Phone, &employee.Email, &employee.BloodGroup); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return employees, nil
}

func addEmployee(employee Employee) (int, error) {
	result, err := db.Exec("INSERT INTO employee(name, phone, email, blood_group) VALUES(?, ?, ?, ?)", employee.Name, employee.Phone, employee.Email, employee.BloodGroup)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}
