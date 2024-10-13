package main

import "fmt"

type (
	Database interface {
		Insert() error
		Update() error
	}

	PostgresDb struct{} // Real database

	MockDb struct{}
)

func NewPostgres() Database {
	return &PostgresDb{}
}

func NewMockDb() Database {
	return &MockDb{}
}

func (p *PostgresDb) Insert() error {
	fmt.Println("Real Insert!!")
	return nil
}

func (p *PostgresDb) Update() error {
	fmt.Println("Real update!!")
	return nil
}

func (p *MockDb) Insert() error {
	fmt.Println("Mock Insert!!")
	return nil
}

func (p *MockDb) Update() error {
	fmt.Println("Mock update!!")
	return nil
}

func InsertPlayerItem(db Database) error {
	return db.Insert()
}

func UpdatePlayerItem(db Database) error {
	return db.Update()
}

func main() {
	p := NewPostgres()
	mock := NewMockDb()

	InsertPlayerItem(p)
	InsertPlayerItem(mock)

	UpdatePlayerItem(p)
	UpdatePlayerItem(mock)
}
