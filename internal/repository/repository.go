package repository

import (
	"chimpanzee/internal/model"
	"context"
)

// Repository is a contract for communicating with underlying datastore
type Repository interface {
	// GetSurveys gets surveys list without actual questions
	GetSurveys(ctx context.Context) ([]model.Survey, error)
	// GetSurvey get survey by id
	GetSurvey(ctx context.Context, id uint) (model.Survey, error)
	// AddSurveyReport add survey report by survey ID
	AddSurveyReport(ctx context.Context, surveyID int, report model.ReportAnswers) error
	// GetReports gets all the reports from datastore
	GetReports(ctx context.Context) ([]model.DBReportResult, error)
}
