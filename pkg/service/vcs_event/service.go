package vcs_event

import (
	"context"
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/linguist-svc"
)

type Service interface {
	UpdateProject(ctx context.Context, project *model.Project) (*model.Project, error)
	GetLanguagesUsedInBranch(context.Context, *linguist.LanguageRequest) []*linguist.Language
}
