package project

import (
	"github.com/vastness-io/coordinator-svc/project"
	"github.com/vastness-io/coordinator/pkg/model"
)

func FromProjectModel(in *model.Project) (out *project.Project) {
	out = &project.Project{
		Name: in.Name,
		Type: in.Type,
	}

	repos := make([]*project.Repository, 0, len(in.Repositories))

	for _, repo := range in.Repositories {
		repos = append(repos, FromRepositoryModel(repo))
	}

	out.Repositories = repos
	return
}

func FromRepositoryModel(in *model.Repository) (out *project.Repository) {
	out = &project.Repository{
		Name:  in.Name,
		Owner: in.Owner,
		Type:  in.Type,
	}

	branches := make([]*project.Branch, 0, len(in.Branches))

	for _, branch := range in.Branches {
		branches = append(branches, FromBranchModel(branch))
	}

	out.Branches = branches
	return
}

func FromBranchModel(in *model.Branch) (out *project.Branch) {
	out = &project.Branch{
		Name: in.Name,
		Meta: FromBranchMetaModel(in.Meta),
	}

	commits := make([]*project.Commit, 0, len(in.Commits))

	for _, commit := range in.Commits {
		commits = append(commits, FromCommitModel(commit))
	}

	out.Commits = commits
	return
}

func FromBranchMetaModel(in *model.BranchMeta) (out *project.BranchMeta) {
	out = &project.BranchMeta{}

	langs := make([]*project.Language, 0, len(in.Languages))

	for name, percentage := range in.Languages {
		langs = append(langs, createLanguage(name, percentage))
	}

	out.Languages = langs
	return
}

func createLanguage(name string, percentage float64) (out *project.Language) {
	out = &project.Language{
		Name:       name,
		Percentage: percentage,
	}
	return
}

func FromCommitModel(in *model.Commit) (out *project.Commit) {
	out = &project.Commit{
		Sha:         in.Sha,
		Message:     in.Message,
		Timestamp:   in.Timestamp.UTC().String(),
		AuthorName:  in.AuthorName,
		AuthorEmail: in.AuthorEmail,
		Added:       in.Added,
		Modified:    in.Modified,
		Removed:     in.Removed,
	}
	return
}
