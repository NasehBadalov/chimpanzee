package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Questions implements driver.Valuer and sql.Scanner for database scanning and inputting
type Questions []Question

// Value is json based implementation of driver.Valuer
func (q Questions) Value() (driver.Value, error) {
	return json.Marshal(q)
}

// Scan is json based implementation of sql.Scanner
func (q *Questions) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &q)
}

// Survey is struct that represents survey held in datastore
type Survey struct {
	ID        uint      `db:"id"`
	Title     string    `db:"title"`
	Questions Questions `db:"questions,omitempty"`
}

// Question is a single entry of survey
// One of Scale, Radio, Checkbox or TextField will contain value
// Check Type to determine which one has non-zero value
type Question struct {
	Type      string   `db:"type"`
	Title     string   `db:"title"`
	Position  uint     `db:"position"`
	Scale     []int    `db:"scale"`
	Checkbox  []string `db:"checkbox"`
	Radio     []string `db:"radio"`
	TextField string   `db:"text_field"`
}
