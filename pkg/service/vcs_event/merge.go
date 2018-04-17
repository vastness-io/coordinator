package vcs_event

import (
	"github.com/vastness-io/coordinator/pkg/model"
)

func MergeProjects(old, new *model.Project) *model.Project {
	if old.Name == new.Name && old.Type == new.Type {
		old.Repositories = MergeRepositories(old.Repositories, new.Repositories)
		return old
	}
	return new
}

func MergeRepositories(old, new []*model.Repository) []*model.Repository {
	for i, _ := range new {
		ok, index := containsRepository(new[i], old)
		if ok {
			new[i].Name = new[i].RepositoryName
			new[i].Owner = new[i].RepositoryOwner
			new[i].Type = new[i].RepositoryType
			old[index].Branches = MergeBranches(old[index].Branches, new[i].Branches)
		} else {
			old = append(old, new[i])
		}
	}

	return old
}

func containsRepository(repository *model.Repository, repositories []*model.Repository) (bool, int) {
	for i, _ := range repositories {
		if (repositories[i].Name == repository.RepositoryName) && (repositories[i].Owner == repository.RepositoryOwner) && (repositories[i].Type == repository.RepositoryType) {
			return true, i
		}
	}
	return false, -1
}

func MergeBranches(old, new []*model.Branch) []*model.Branch {
	for i, _ := range new {
		ok, index := containsBranch(new[i], old)
		if ok {
			old[index].Commits = MergeCommits(old[index].Commits, new[i].Commits)
		} else {
			old = append(old, new[i])
		}
	}

	return old
}

func containsBranch(branch *model.Branch, branches []*model.Branch) (bool, int) {
	for i, _ := range branches {
		if (branches[i].Name == branch.Name) && (branches[i].RepositoryName == branch.RepositoryName) && (branches[i].RepositoryOwner == branch.RepositoryOwner) && (branches[i].RepositoryType == branch.RepositoryType) {
			return true, i
		}
	}
	return false, -1
}

func MergeCommits(old, new []*model.Commit) []*model.Commit {

	for i, _ := range new {
		ok, index := containsCommit(new[i], old)
		if ok {
			old[index] = new[i]
		} else {
			old = append(old, new[i])
		}
	}

	return old
}

func containsCommit(commit *model.Commit, commits []*model.Commit) (bool, int) {
	for i, _ := range commits {
		if commits[i].Sha == commit.Sha {
			return true, i
		}
	}
	return false, -1
}
