package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/lib/pq"
)

// ReportAnswers implements driver.Valuer and sql.Scanner for database scanning and inputting
type ReportAnswers []ReportAnswer

// Value is json based implementation of driver.Valuer
func (r ReportAnswers) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan is json based implementation of sql.Scanner
func (r *ReportAnswers) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &r)
}

// ReportAnswer contains fields of answer
// One of Scale, Radio, Checkbox or TextField will contain value
// Check Type to determine which one has non-zero value
type ReportAnswer struct {
	Type      string   `json:"type" db:"type"`
	Position  int      `json:"position" db:"position"`
	Scale     int      `json:"scale,omitempty" db:"scale"`
	Radio     string   `json:"radio,omitempty" db:"radio"`
	Checkbox  []string `json:"checkbox,omitempty" db:"checkbox"`
	TextField string   `json:"textfield,omitempty" db:"textfield"`
}

// DBReportResult contains fields of answer
// One of Scale, Radio, Checkbox or TextField will contain value
// Check Type to determine which one has non-zero value
// ID is a database identifier of report
type DBReportResult struct {
	ID        int            `db:"id"`
	SurveyID  int            `db:"survey_id"`
	Type      string         `db:"type"`
	Question  string         `db:"question"`
	Position  int            `db:"position"`
	Scale     *int           `db:"scale"`
	TextField *string        `db:"textfield"`
	Checkbox  pq.StringArray `db:"checkbox"`
	Radio     *string        `db:"radio"`
}

// DBReportResultToReportResult is used to convert db record typed DBReportResult to model typed ReportResult that is used in service.Service
func DBReportResultToReportResult(r DBReportResult) ReportResult {
	return ReportResult{
		ID:        r.ID,
		SurveyID:  r.SurveyID,
		Type:      r.Type,
		Question:  r.Question,
		Position:  r.Position,
		Scale:     r.Scale,
		TextField: r.TextField,
		Checkbox:  r.Checkbox,
		Radio:     r.Radio,
	}
}

// ReportResult contains fields of answer
// One of Scale, Radio, Checkbox or TextField will contain value
// Check Type to determine which one has non-zero value
// ID is the identifier of report
type ReportResult struct {
	ID        int      `json:"id"`
	SurveyID  int      `json:"survey_id"`
	Type      string   `json:"type"`
	Question  string   `json:"question"`
	Position  int      `json:"position"`
	Scale     *int     `json:"scale"`
	TextField *string  `json:"textfield"`
	Checkbox  []string `json:"checkbox"`
	Radio     *string  `json:"radio"`
}
