package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQuestions_Value(t *testing.T) {
	assert := require.New(t)
	q := []Question{{
		Type:      SurveyTypeTextField,
		Position:  1,
		TextField: "123",
	}}
	qJSON, err := json.Marshal(q)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		q       Questions
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "json",
			q:       q,
			want:    qJSON,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Value()
			if tt.wantErr == false {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}

func TestQuestions_Scan(t *testing.T) {
	assert := require.New(t)
	q := []Question{{
		Type:      SurveyTypeTextField,
		Position:  1,
		TextField: "123",
	}}
	qJSON, err := json.Marshal(q)
	if err != nil {
		panic(err)
	}

	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		q       Questions
		args    args
		wantErr bool
	}{
		{
			name: "json",
			q:    q,
			args: args{
				value: qJSON,
			},
			wantErr: false,
		},
		{
			name: "malformed json",
			q:    q,
			args: args{
				value: []uint8{'s'},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.q.Scan(tt.args.value)
			if tt.wantErr == false {
				assert.NoError(err)
			}
		})
	}
}
