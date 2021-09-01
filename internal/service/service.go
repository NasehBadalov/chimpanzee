package service

import (
	"chimpanzee/internal/model"
	"chimpanzee/internal/repository"
	"context"
)

// IService is a contract for interacting with business logic layer
type IService interface {
	// GetSurveys gets surveys list
	GetSurveys(ctx context.Context) ([]model.Survey, error)
	// GetSurvey get survey by id
	GetSurvey(ctx context.Context, id uint) (model.Survey, error)
	// AddSurveyReport add survey report by survey ID
	AddSurveyReport(ctx context.Context, surveyID int, report model.ReportAnswers) error
	// GetReports gets all the reports
	GetReports(ctx context.Context) ([]model.ReportResult, error)
}

// Service struct is an actual implementation of IService
type Service struct {
	repo repository.Repository
}

// ProvideService provides *Service as IService for dependency injection
func ProvideService(repo repository.Repository) IService {
	return &Service{repo: repo}
}

// GetSurveys gets surveys list
func (s *Service) GetSurveys(ctx context.Context) ([]model.Survey, error) {
	return s.repo.GetSurveys(ctx)
}

// GetSurvey get survey by id
func (s *Service) GetSurvey(ctx context.Context, id uint) (model.Survey, error) {
	return s.repo.GetSurvey(ctx, id)
}

// AddSurveyReport add survey report by survey ID
func (s *Service) AddSurveyReport(ctx context.Context, surveyID int, report model.ReportAnswers) error {
	return s.repo.AddSurveyReport(ctx, surveyID, report)
}

// GetReports gets all the reports
func (s *Service) GetReports(ctx context.Context) ([]model.ReportResult, error) {
	drp, err := s.repo.GetReports(ctx)
	if err != nil {
		return nil, err
	}
	var rr []model.ReportResult
	for _, r := range drp {
		rr = append(rr, model.DBReportResultToReportResult(r))
	}
	return rr, nil
}
