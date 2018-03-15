package webhook

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/coordinator/pkg/repository"
	"github.com/vastness-io/linguist-svc"
	"time"
)

type vcsEventService struct {
	linguist   linguist.LinguistClient
	logger     *logrus.Entry
	repository repository.ProjectRepository
}

func NewVcsEventService(logger *logrus.Entry, linguistClient linguist.LinguistClient, projectRepository repository.ProjectRepository) Service {
	return &vcsEventService{
		linguist:   linguistClient,
		logger:     logger,
		repository: projectRepository,
	}
}

func (s *vcsEventService) UpdateProject(project *model.Project) (*model.Project, error) {

	tx := s.repository.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return nil, tx.Error
	}

	current, err := s.repository.GetProject(tx, project.Name, project.Type)

	if err != nil {
		if err != errors.ProjectDoesNotExistErr {
			return nil, err
		}
		if err := s.repository.Create(tx, project); err != nil {
			tx.Rollback()
			return nil, err
		}
		current = project
	}


	for _, repo := range current.Repositories {
		for _, branch := range repo.Branches {
			req := linguist.LanguageRequest{}

			var files []string
			for _, commit := range branch.Commits {
				files = append(files, commit.Added...)
				files = append(files, commit.Modified...)
				//TODO handle removed files
			}

			req.FileNames = files

			branch.Languages = convertLangResponse(s.GetLanguagesUsedInRepository(&req))
		}

	}

	err = s.repository.Update(tx, current)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return current, tx.Commit().Error

}

func (s *vcsEventService) GetLanguagesUsedInRepository(req *linguist.LanguageRequest) []*linguist.Language {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Don't think this is the right way to do it?
	defer cancel()

	res, err := s.linguist.GetLanguages(ctx, req)

	if err != nil {
		s.logger.Error("Unable to detect Language(s) for repository")
		return make([]*linguist.Language, 0)
	}

	return res.GetLanguage()

}

func convertLangResponse(res []*linguist.Language) []*model.Language {
	var out []*model.Language

	for _, lang := range res {
		out = append(out, &model.Language{
			Name:       lang.GetName(),
			Percentage: lang.GetPercentage(),
		})
	}
	return out
}
