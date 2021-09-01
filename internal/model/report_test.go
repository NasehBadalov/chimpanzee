package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDBReportResultToReportResult(t *testing.T) {
	assert := require.New(t)
	type args struct {
		r DBReportResult
	}
	tests := []struct {
		name string
		args args
		want ReportResult
	}{
		{
			name: "Basic",
			args: args{
				r: DBReportResult{
					ID:        1,
					SurveyID:  2,
					Type:      "3",
					Question:  "4",
					Position:  5,
					Scale:     &[]int{6}[0],
					TextField: &[]string{"7"}[0],
					Checkbox:  []string{"8"},
					Radio:     &[]string{"9"}[0],
				},
			},
			want: ReportResult{
				ID:        1,
				SurveyID:  2,
				Type:      "3",
				Question:  "4",
				Position:  5,
				Scale:     &[]int{6}[0],
				TextField: &[]string{"7"}[0],
				Checkbox:  []string{"8"},
				Radio:     &[]string{"9"}[0],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DBReportResultToReportResult(tt.args.r)
			assert.Equal(tt.want, got)
		})
	}
}

func TestReportAnswers_Value(t *testing.T) {
	assert := require.New(t)
	ra := []ReportAnswer{{
		Type:     SurveyTypeScale,
		Position: 1,
		Scale:    2,
	}}
	raJSON, err := json.Marshal(ra)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		r       ReportAnswers
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "json",
			r:       ra,
			want:    raJSON,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Value()
			if tt.wantErr == false {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}

func TestReportAnswers_Scan(t *testing.T) {
	assert := require.New(t)
	ra := []ReportAnswer{{
		Type:     SurveyTypeScale,
		Position: 1,
		Scale:    2,
	}}
	raJSON, err := json.Marshal(ra)
	if err != nil {
		panic(err)
	}

	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		r       ReportAnswers
		args    args
		wantErr bool
	}{
		{
			name:    "json",
			r:       ra,
			args:    args{
				value: raJSON,
			},
			wantErr: false,
		},
		{
			name:    "malformed json",
			r:       ra,
			args:    args{
				value: []uint8{'s'},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.r.Scan(tt.args.value)
			if tt.wantErr == false {
				assert.NoError(err)
			}
		})
	}
}
