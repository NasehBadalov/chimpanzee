package dto

import "chimpanzee/internal/model"

// Survey is a set of questions and title
type Survey struct {
	ID        uint        `json:"id"`
	Title     string      `json:"title"`
	Questions *[]Question `json:"questions,omitempty"`
}
// Question is a single question from Survey
type Question struct {
	Type     string         `json:"type"`
	Title    string         `json:"title"`
	Position uint           `json:"position"`
	Answer   QuestionAnswer `json:"answer"`
}

type QuestionAnswer interface{}

// NewSurveyList is serializer for service data
func NewSurveyList(survey []model.Survey) []Survey {
	var sl []Survey
	for _, s := range survey {
		sl = append(sl, Survey{
			ID:    s.ID,
			Title: s.Title,
		})
	}
	return sl
}

// NewSurvey is serializer for service data
func NewSurvey(survey model.Survey) Survey {
	var questions []Question
	for _, q := range survey.Questions {
		var answer QuestionAnswer

		switch q.Type {
		case model.SurveyTypeCheckbox:
			answer = q.Checkbox
		case model.SurveyTypeRadio:
			answer = q.Radio
		case model.SurveyTypeScale:
			answer = q.Scale
		case model.SurveyTypeTextField:
			answer = q.TextField
		}

		questions = append(questions, Question{
			Type:     q.Type,
			Title:    q.Title,
			Position: q.Position,
			Answer:   answer,
		})

	}
	return Survey{
		ID:        survey.ID,
		Title:     survey.Title,
		Questions: &questions,
	}
}
