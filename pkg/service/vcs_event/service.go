package vcs_event

import (
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/linguist-svc"
)

type Service interface {
	UpdateProject(project *model.Project) (*model.Project, error)
	GetLanguagesUsedInRepository(*linguist.LanguageRequest) []*linguist.Language
}
