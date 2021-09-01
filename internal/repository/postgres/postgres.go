package postgres

import (
	"chimpanzee/internal/config"
	"chimpanzee/internal/model"
	"chimpanzee/internal/repository"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"time"
)

// Repository struct implements repository.Repository using postgresql connection
type Repository struct {
	c *sqlx.DB
}

// ProvideRepository provides *Repository as repository.Repository for dependency injection
func ProvideRepository(cfg *config.Config) repository.Repository {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Database,
	)

	var db *sqlx.DB
	var err error
	const retries = 10
	toSleep := 1
	for i := 1; i <= retries; i++ {
		db, err = sqlx.Connect("postgres", connString)
		if err != nil && i != retries {
			zap.S().Infof("Coundn't connect to postgres, waiting %d seconds and retrying", toSleep)
			time.Sleep(time.Duration(toSleep) * time.Second)
			toSleep *= 2
			continue
		} else if err != nil {
			zap.S().Fatal(err)
		}
	}

	repo := &Repository{c: db}

	err = repo.migrate(context.TODO())
	if err != nil {
		zap.S().Fatal(err)
	}

	return repo
}

// migrate runs initial scripts in db
// must contain idempotent operations (sql statements)
func (r *Repository) migrate(ctx context.Context) error {
	tx, err := r.c.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`create table if not exists survey
(
    id        serial,
    title     varchar(120),
    questions jsonb,
    primary key (id)
)`)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`create table if not exists reports
(
    id        serial,
    answers   jsonb,
    survey_id int
        constraint reports_survey_id_fk
            references survey (id)
            on update cascade on delete cascade,
    primary key (id)
);`)
	if err != nil {
		return err
	}

	// add survey if not exists
	_, err = tx.Exec(`insert into survey(title, questions)
select 'Customer Satisfaction Survey Template', '[
  {
    "type": "scale",
    "title": "How likely is it that you would recommend this company to a friend or colleague?",
    "position": 1,
    "scale": [
      0,
      10
    ]
  },
  {
    "type": "radio",
    "title": "Overall, how satisfied or dissatisfied are you with our company?",
    "position": 2,
    "radio": [
      "Very satisfied",
      "Somewhat dissatisfied",
      "Somewhat satisfied",
      "Very dissatisfied",
      "Neither satisfied nor dissatisfied"
    ]
  },
  {
    "type": "checkbox",
    "title": "Which of the following words would you use to describe our products? Select all that apply.",
    "position": 3,
    "checkbox": [
      "Reliable",
      "Overpriced",
      "High quality",
      "Impractical",
      "Useful",
      "Unique",
      "Poor quality",
      "Good value for money",
      "Unreliable"
    ]
  },
  {
    "type": "radio",
    "title": "How well do our products meet your needs?",
    "position": 4,
    "radio": [
      "Extremely well",
      "Not so well",
      "Very well",
      "Not at all well",
      "Somewhat well"
    ]
  },
  {
    "type": "radio",
    "title": "How would you rate the quality of the product?",
    "position": 5,
    "radio": [
      "Very high quality",
      "Low quality",
      "High quality",
      "Very low quality",
      "Neither high nor low quality"
    ]
  },
  {
    "type": "radio",
    "title": "How would you rate the value for money of the product?",
    "position": 6,
    "radio": [
      "Excellent",
      "Below average",
      "Above average",
      "Poor",
      "Average"
    ]
  },
  {
    "type": "radio",
    "title": "How responsive have we been to your questions or concerns about our products?",
    "position": 7,
    "radio": [
      "Extremely responsive",
      "Not so responsive",
      "Very responsive",
      "Not at all responsive",
      "Somewhat responsive",
      "Not applicable"
    ]
  },
  {
    "type": "radio",
    "title": "How long have you been a customer of our company?",
    "position": 8,
    "radio": [
      "This is my first purchase",
      "1 - 2 years",
      "Less than six months",
      "3 or more years",
      "Six months to a year",
      "I havent made a purchase yet"
    ]
  },
  {
    "type": "radio",
    "title": "How long have you been a customer of our company?",
    "position": 9,
    "radio": [
      "Extremely likely",
      "Not so likely",
      "Very likely",
      "Not at all likely",
      "Somewhat likely"
    ]
  },
  {
    "type": "textfield",
    "title": "Do you have any other comments, questions, or concerns?",
    "position": 10
  }
]'
where not exists(
        select id from survey
    );`)
	if err != nil {
		return err
	}

	// add report if not exists
	_, err = tx.Exec(`insert into reports(survey_id, answers) 
select 1, '[
  {
    "type": "scale",
    "position": 1,
    "scale": 5
  },
  {
    "type": "radio",
    "position": 2,
    "radio": "Very satisfied"
  },
  {
    "type": "checkbox",
    "position": 3,
    "checkbox": [
      "Reliable",
      "High quality",
      "Impractical",
      "Good value for money"
    ]
  },
  {
    "type": "radio",
    "position": 4,
    "radio": "Somewhat well"
  },
  {
    "type": "radio",
    "position": 5,
    "radio": "High quality"
  },
  {
    "type": "radio",
    "position": 6,
    "radio": "Poor"
  },
  {
    "type": "radio",
    "position": 7,
    "radio": "Very responsive"
  },
  {
    "type": "radio",
    "position": 8,
    "radio": "1 - 2 years"
  },
  {
    "type": "radio",
    "position": 9,
    "radio": "Very likely"
  },
  {
    "type": "textfield",
    "position": 10,
    "radio": "Good job!"
  }
]'
where not exists(
        select id from reports
    );`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// GetSurveys gets surveys list without actual questions
func (r *Repository) GetSurveys(ctx context.Context) ([]model.Survey, error) {
	var surveys []model.Survey
	err := r.c.SelectContext(ctx, &surveys, "select id, title from survey")
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurvey get survey by id
func (r *Repository) GetSurvey(ctx context.Context, id uint) (model.Survey, error) {
	var survey model.Survey
	stmt, err := r.c.PreparexContext(ctx, "select id, title, questions from survey where id=$1")
	if err != nil {
		return survey, err
	}
	err = stmt.GetContext(ctx, &survey, id)
	if err != nil {
		return survey, err
	}
	return survey, err
}

// AddSurveyReport add survey report by survey ID
func (r *Repository) AddSurveyReport(ctx context.Context, surveyID int, report model.ReportAnswers) error {
	stmt, err := r.c.PreparexContext(ctx, "insert into reports (survey_id, answers) values ($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, surveyID, report)
	if err != nil {
		return err
	}
	return nil
}

// GetReports gets all the reports from datastore
func (r *Repository) GetReports(ctx context.Context) ([]model.DBReportResult, error) {
	var rr []model.DBReportResult
	err := r.c.SelectContext(ctx, &rr, `select r.id,
       s.id as survey_id,
       type,
       qs.title as question,
       qs.position,
       scale,
       textfield,
       checkbox,
       radio
from reports r
         join survey s on s.id = r.survey_id
         cross join lateral jsonb_to_recordset(questions) as qs(position integer, title text)
         cross join lateral jsonb_to_recordset(answers) as ans(position integer, type text,
                                                               scale integer, textfield text, checkbox text[],
                                                               radio text)
where qs.position = ans.position`)
	if err != nil {
		return nil, err
	}

	return rr, nil
}
