package dto

import (
	"chimpanzee/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewReportAnswer(t *testing.T) {
	assert := require.New(t)

	type args struct {
		r []ReportAnswer
	}
	tests := []struct {
		name    string
		args    args
		want    []model.ReportAnswer
		wantErr bool
	}{
		{
			name: "Radio",
			args: args{
				r: []ReportAnswer{{
					Position: 1,
					Type:     model.SurveyTypeRadio,
					Answer:   "meow",
				}},
			},
			want: []model.ReportAnswer{{
				Type:     model.SurveyTypeRadio,
				Position: 1,
				Radio:    "meow",
			}},
			wantErr: false,
		},
		{
			name: "Scale",
			args: args{
				r: []ReportAnswer{{
					Position: 1,
					Type:     model.SurveyTypeScale,
					Answer:   float64(4),
				}},
			},
			want: []model.ReportAnswer{{
				Type:     model.SurveyTypeScale,
				Position: 1,
				Scale:    4,
			}},
			wantErr: false,
		},
		{
			name: "non float64 Scale",
			args: args{
				r: []ReportAnswer{{
					Position: 1,
					Type:     model.SurveyTypeScale,
					Answer:   4,
				}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Checkbox",
			args: args{
				r: []ReportAnswer{{
					Position: 1,
					Type:     model.SurveyTypeCheckbox,
					Answer:   []interface{}{"hello", "dunya"},
				}},
			},
			want: []model.ReportAnswer{{
				Type:     model.SurveyTypeCheckbox,
				Position: 1,
				Checkbox: []string{"hello", "dunya"},
			}},
			wantErr: false,
		},
		{
			name: "non []interface Checkbox",
			args: args{
				r: []ReportAnswer{{
					Position: 1,
					Type:     model.SurveyTypeCheckbox,
					Answer:   []string{"hello", "dunya"},
				}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "TextField",
			args: args{
				r: []ReportAnswer{{
					Position: 1,
					Type:     model.SurveyTypeTextField,
					Answer:   "Hello, dunya",
				}},
			},
			want: []model.ReportAnswer{{
				Type:      model.SurveyTypeTextField,
				Position:  1,
				TextField: "Hello, dunya",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReportAnswer(tt.args.r)
			if tt.wantErr == false {
				assert.NoError(err)
			}
			assert.Equal(tt.want, got)
		})
	}
}

func TestNewReportResults(t *testing.T) {
	assert := require.New(t)
	randomString := "Hello world"
	type args struct {
		r []model.ReportResult
	}
	tests := []struct {
		name string
		args args
		want ReportResult
	}{
		{
			name: "Radio",
			args: args{
				r: []model.ReportResult{{
					Type:  model.SurveyTypeRadio,
					Radio: &randomString,
				}},
			},
			want: ReportResult{
				Reports: []ReportEntity{{
					Answers: []ReportEntry{{
						Type:   model.SurveyTypeRadio,
						Answer: &randomString,
					}},
				}},
			},
		},
		{
			name: "Checkbox",
			args: args{
				r: []model.ReportResult{{
					Type:     model.SurveyTypeCheckbox,
					Checkbox: []string{"hello", "world"},
				}},
			},
			want: ReportResult{Reports: []ReportEntity{{
				Answers: []ReportEntry{{
					Type:   model.SurveyTypeCheckbox,
					Answer: []string{"hello", "world"},
				}},
			}},
			},
		},
		{
			name: "Scale",
			args: args{
				r: []model.ReportResult{{
					Type:  model.SurveyTypeScale,
					Scale: &[]int{4}[0],
				}},
			},
			want: ReportResult{
				Reports: []ReportEntity{{
					Answers: []ReportEntry{{
						Type:   model.SurveyTypeScale,
						Answer: &[]int{4}[0],
					}},
				}},
			},
		},
		{
			name: "TextField",
			args: args{
				r: []model.ReportResult{{
					Type:      model.SurveyTypeTextField,
					TextField: &randomString,
				}},
			},
			want: ReportResult{
				Reports: []ReportEntity{{
					Answers: []ReportEntry{{
						Type:   model.SurveyTypeTextField,
						Answer: &randomString,
					}},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewReportResults(tt.args.r)
			assert.Equal(tt.want, got)
		})
	}
}
