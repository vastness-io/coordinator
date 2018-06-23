package vcs_event

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

func (s *vcsEventService) UpdateProject(ctx context.Context, project *model.Project) (*model.Project, error) {

	tx := s.repository.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error() != nil {
		return nil, tx.Error()
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

				if commit.Meta != nil && len(commit.Meta.Languages) == 0 {

					var filesUptoCurrentCommit []string

					for k, v := range fileStatus {
						if v.removed {
							continue
						}
						filesUptoCurrentCommit = append(filesUptoCurrentCommit, k)
					}

					req := linguist.LanguageRequest{
						FileNames: filesUptoCurrentCommit,
					}

					if languages := s.GetLanguages(ctx, &req); len(languages) != 0 {
						commit.Meta.Languages = ConvertToLanguages(languages)
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

			req := linguist.LanguageRequest{
				FileNames: sanitizedFiles,
			}

			if languages := s.GetLanguages(ctx, &req); len(languages) != 0 {
				branch.Meta.Languages = ConvertToLanguages(languages)
			}

		}

	}

	err = s.repository.Update(tx, current)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return current, tx.Commit().Error()

}

func (s *vcsEventService) GetLanguages(ctx context.Context, req *linguist.LanguageRequest) []*linguist.Language {
	res, err := s.linguist.GetLanguages(ctx, req)

	if err != nil {
		s.logger.Error("Unable to detect Language(s)")
		return make([]*linguist.Language, 0)
	}

	return res.GetLanguage()

}

func ConvertToLanguages(langs []*linguist.Language) model.Languages {
	out := make(model.Languages)

	for _, lang := range langs {
		out[lang.GetName()] = lang.GetPercentage()
	}

	return out
}
