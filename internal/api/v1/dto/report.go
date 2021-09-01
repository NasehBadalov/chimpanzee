package dto

import (
	"chimpanzee/internal/model"
	"errors"
	"sort"
)

// CannotConvertRequest indicates that malformed data is passed that is can't be processed
var CannotConvertRequest = errors.New("cannot convert request")

// Report is struct that represents report data with answers and corresponding survey id
type Report struct {
	SurveyID int            `json:"survey_id"`
	Answers  []ReportAnswer `json:"answers"`
}

// ReportAnswer is a single answer for survey
type ReportAnswer struct {
	Position int            `json:"position"`
	Type     string         `json:"type"`
	Answer   QuestionAnswer `json:"answer"`
}

// NewReportAnswer converts dto to model
func NewReportAnswer(r []ReportAnswer) ([]model.ReportAnswer, error) {
	var ra []model.ReportAnswer

	for _, a := range r {
		t := model.ReportAnswer{
			Type:     a.Type,
			Position: a.Position,
		}

		// conditionally get answer using Type field to match which field to get it from
		var ok bool
		switch a.Type {
		case model.SurveyTypeCheckbox:
			cb, ok := a.Answer.([]interface{})
			if !ok {
				return nil, CannotConvertRequest
			}
			var cba []string
			for _, c := range cb {
				cc, ok := c.(string)
				if !ok {
					return nil, CannotConvertRequest
				}
				cba = append(cba, cc)
			}

			t.Checkbox = cba
		case model.SurveyTypeRadio:
			t.Radio, ok = a.Answer.(string)
			if !ok {
				return nil, CannotConvertRequest
			}
		case model.SurveyTypeScale:
			scaleF, ok := a.Answer.(float64)
			if !ok {
				return nil, CannotConvertRequest
			}
			t.Scale = int(scaleF)
		case model.SurveyTypeTextField:
			t.TextField, ok = a.Answer.(string)
			if !ok {
				return nil, CannotConvertRequest
			}
		}

		ra = append(ra, t)

	}
	return ra, nil
}

// ReportResultSorted is representation of report result for api layer
// map[SurveyID]map[ID] []...
type ReportResultSorted map[int]map[int][]ReportEntry

type ReportResult struct {
	SurveyID int            `json:"survey_id"`
	Reports  []ReportEntity `json:"reports"`
}

type ReportEntitySorter []ReportEntity

func (r ReportEntitySorter) Len() int {
	return len(r)
}

func (r ReportEntitySorter) Less(i, j int) bool {
	return r[i].ReportID < r[j].ReportID
}

func (r ReportEntitySorter) Swap(i, j int) {
	r[i].ReportID, r[j].ReportID = r[j].ReportID, r[i].ReportID
}

type ReportEntity struct {
	ReportID int           `json:"report_id"`
	Answers  []ReportEntry `json:"answers"`
}

type ReportEntry struct {
	Type     string         `json:"type"`
	Question string         `json:"question"`
	Position int            `json:"position"`
	Answer   QuestionAnswer `json:"answer"`
}

// NewReportResults converts model to dto
func NewReportResults(reportResult []model.ReportResult) ReportResult {
	rrs := make(ReportResultSorted)

	for _, rr := range reportResult {
		var answer QuestionAnswer

		// conditionally set answer using Type field to match which field to fill
		switch rr.Type {
		case model.SurveyTypeCheckbox:
			answer = rr.Checkbox
		case model.SurveyTypeRadio:
			answer = rr.Radio
		case model.SurveyTypeScale:
			answer = rr.Scale
		case model.SurveyTypeTextField:
			answer = rr.TextField
		}

		if rrs[rr.SurveyID] == nil {
			rrs[rr.SurveyID] = make(map[int][]ReportEntry)
		} else if rrs[rr.SurveyID][rr.ID] == nil {
			rrs[rr.SurveyID][rr.ID] = []ReportEntry{}
		}
		rrs[rr.SurveyID][rr.ID] = append(rrs[rr.SurveyID][rr.ID], ReportEntry{
			Type:     rr.Type,
			Question: rr.Question,
			Position: rr.Position,
			Answer:   answer,
		})
	}

	var rr ReportResult
	for surveyID, idEntry := range rrs {
		rr.SurveyID = surveyID
		for id, entries := range idEntry {
			rr.Reports = append(rr.Reports, ReportEntity{ReportID: id, Answers: entries})
		}
	}
		sort.Sort(ReportEntitySorter(rr.Reports))
	return rr
}
