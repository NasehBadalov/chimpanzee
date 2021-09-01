package service

import (
	"chimpanzee/internal/mocks"
	"chimpanzee/internal/model"
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_AddSurveyReport(t *testing.T) {
	assert := require.New(t)
	t.Run("Was called once", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("AddSurveyReport", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		svc := ProvideService(repo)
		err := svc.AddSurveyReport(context.TODO(), 1, model.ReportAnswers{})
		assert.NoError(err)
		repo.AssertNumberOfCalls(t, "AddSurveyReport", 1)
	})
}

func TestService_GetReports(t *testing.T) {
	assert := require.New(t)
	t.Run("Was called once", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetReports", mock.Anything).Return([]model.DBReportResult{}, nil)
		svc := ProvideService(repo)
		_, err := svc.GetReports(context.TODO())
		assert.NoError(err)
		repo.AssertNumberOfCalls(t, "GetReports", 1)
	})
}

func TestService_GetSurvey(t *testing.T) {
	assert := require.New(t)
	t.Run("Was called once", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetSurvey", mock.Anything, mock.Anything).Return(model.Survey{}, nil)
		svc := ProvideService(repo)
		_, err := svc.GetSurvey(context.TODO(), 42)
		assert.NoError(err)
		repo.AssertNumberOfCalls(t, "GetSurvey", 1)
	})
}

func TestService_GetSurveys(t *testing.T) {
	assert := require.New(t)
	t.Run("Was called once", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetSurveys", mock.Anything).Return([]model.Survey{}, nil)
		svc := ProvideService(repo)
		_, err := svc.GetSurveys(context.TODO())
		assert.NoError(err)
		repo.AssertNumberOfCalls(t, "GetSurveys", 1)
	})
}
