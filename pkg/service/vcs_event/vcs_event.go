package vcs_event

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/coordinator/pkg/repository"
	"github.com/vastness-io/coordinator/pkg/shared"
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

	if tx.GetError() != nil {
		return nil, tx.GetError()
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
	} else {
		current = MergeProjects(current, project)
	}

	for _, repo := range current.Repositories {
		for _, branch := range repo.Branches {
			req := linguist.LanguageRequest{}

			type timeDiff struct {
				removed bool
				t       time.Time
			}

			fileStatus := make(map[string]*timeDiff)

			for _, commit := range branch.Commits {

				timestamp := *commit.Timestamp

				for _, add := range commit.Added {

					v, ok := fileStatus[add]

					if ok {
						if timestamp.After(v.t) {
							v.removed = false
							v.t = timestamp
						}
					} else {
						fileStatus[add] = &timeDiff{
							t: timestamp,
						}
					}

				}

				for _, mod := range commit.Modified {

					v, ok := fileStatus[mod]

					if ok {
						if timestamp.After(v.t) {
							v.removed = false
							v.t = timestamp
						}
					} else {
						fileStatus[mod] = &timeDiff{
							t: timestamp,
						}
					}

				}

				for _, rem := range commit.Removed {
					fileStatus[rem] = &timeDiff{
						removed: true,
						t:       timestamp,
					}

				}
			}

			var sanitizedFiles []string

			for k, v := range fileStatus {
				if v.removed {
					continue
				}
				sanitizedFiles = append(sanitizedFiles, k)
			}

			req.FileNames = RemoveDirectoryPrefix(sanitizedFiles)

			if languages := s.GetLanguagesUsedInBranch(&req); len(languages) != 0 {
				branch.Meta.SetLanguages(ConvertToBranchLanguages(languages))
			}

		}

	}

	err = s.repository.Update(tx, current)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return current, tx.Commit().GetError()

}

func (s *vcsEventService) GetLanguagesUsedInBranch(req *linguist.LanguageRequest) []*linguist.Language {

	ctx, cancel := context.WithTimeout(context.Background(), shared.LanguageReqTimeout) // Don't think this is the right way to do it?
	defer cancel()

	res, err := s.linguist.GetLanguages(ctx, req)

	if err != nil {
		s.logger.Error("Unable to detect Language(s) for branch")
		return make([]*linguist.Language, 0)
	}

	return res.GetLanguage()

}

func ConvertToBranchLanguages(langs []*linguist.Language) model.BranchLanguages {
	out := make(model.BranchLanguages)

	for _, lang := range langs {
		out[lang.GetName()] = lang.GetPercentage()
	}

	return out
}
