package dto

import (
	"chimpanzee/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSurvey(t *testing.T) {
	assert := require.New(t)
	type args struct {
		survey model.Survey
	}
	tests := []struct {
		name string
		args args
		want Survey
	}{
		{
			name: "Radio",
			args: args{
				survey: model.Survey{
					ID:    1,
					Title: "Heeey",
					Questions: []model.Question{{
						Type:     model.SurveyTypeRadio,
						Title:    "hello",
						Position: 4,
						Radio:    []string{"hello"},
					}},
				},
			},
			want: Survey{
				ID:    1,
				Title: "Heeey",
				Questions: &[]Question{{
					Type:     model.SurveyTypeRadio,
					Title:    "hello",
					Position: 4,
					Answer:   []string{"hello"},
				}},
			},
		},
		{
			name: "Checkbox",
			args: args{
				survey: model.Survey{
					ID:    1,
					Title: "Heeey",
					Questions: []model.Question{{
						Type:     model.SurveyTypeCheckbox,
						Title:    "hello",
						Position: 4,
						Checkbox: []string{"hello"},
					}},
				},
			},
			want: Survey{
				ID:    1,
				Title: "Heeey",
				Questions: &[]Question{{
					Type:     model.SurveyTypeCheckbox,
					Title:    "hello",
					Position: 4,
					Answer:   []string{"hello"},
				}},
			},
		},
		{
			name: "Scale",
			args: args{
				survey: model.Survey{
					ID:    1,
					Title: "Heeey",
					Questions: []model.Question{{
						Type:     model.SurveyTypeScale,
						Title:    "hello",
						Position: 4,
						Scale:    []int{1, 10},
					}},
				},
			},
			want: Survey{
				ID:    1,
				Title: "Heeey",
				Questions: &[]Question{{
					Type:     model.SurveyTypeScale,
					Title:    "hello",
					Position: 4,
					Answer:   []int{1, 10},
				}},
			},
		},
		{
			name: "TextField",
			args: args{
				survey: model.Survey{
					ID:    1,
					Title: "Heeey",
					Questions: []model.Question{{
						Type:      model.SurveyTypeTextField,
						Title:     "hello",
						Position:  4,
						TextField: "lorem ipsum",
					}},
				},
			},
			want: Survey{
				ID:    1,
				Title: "Heeey",
				Questions: &[]Question{{
					Type:     model.SurveyTypeTextField,
					Title:    "hello",
					Position: 4,
					Answer:   "lorem ipsum",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSurvey(tt.args.survey)
			assert.Equal(tt.want, got)
		})
	}
}

func TestNewSurveyList(t *testing.T) {
	assert := require.New(t)
	type args struct {
		survey []model.Survey
	}
	tests := []struct {
		name string
		args args
		want []Survey
	}{
		{
			name: "Basic",
			args: args{
				survey: []model.Survey{{
					ID:    1,
					Title: "Hello",
				}},
			},
			want: []Survey{{
				ID:    1,
				Title: "Hello",
			}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSurveyList(tt.args.survey)
			assert.Equal(tt.want, got)
		})
	}
}
