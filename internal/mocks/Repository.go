// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	model "chimpanzee/internal/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddSurveyReport provides a mock function with given fields: ctx, surveyID, report
func (_m *Repository) AddSurveyReport(ctx context.Context, surveyID int, report model.ReportAnswers) error {
	ret := _m.Called(ctx, surveyID, report)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, model.ReportAnswers) error); ok {
		r0 = rf(ctx, surveyID, report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetReports provides a mock function with given fields: ctx
func (_m *Repository) GetReports(ctx context.Context) ([]model.DBReportResult, error) {
	ret := _m.Called(ctx)

	var r0 []model.DBReportResult
	if rf, ok := ret.Get(0).(func(context.Context) []model.DBReportResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DBReportResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSurvey provides a mock function with given fields: ctx, id
func (_m *Repository) GetSurvey(ctx context.Context, id uint) (model.Survey, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Survey
	if rf, ok := ret.Get(0).(func(context.Context, uint) model.Survey); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Survey)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSurveys provides a mock function with given fields: ctx
func (_m *Repository) GetSurveys(ctx context.Context) ([]model.Survey, error) {
	ret := _m.Called(ctx)

	var r0 []model.Survey
	if rf, ok := ret.Get(0).(func(context.Context) []model.Survey); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Survey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
