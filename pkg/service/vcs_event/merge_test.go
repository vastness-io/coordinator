package vcs_event

import (
	"github.com/vastness-io/coordinator/pkg/model"
	"reflect"
	"testing"
)

func TestMergeProjects(t *testing.T) {
	tests := []struct {
		old *model.Project
		new *model.Project
		out *model.Project
	}{
		{
			old: &model.Project{
				Name: "1",
			},
			new: &model.Project{
				Name: "1",
				Repositories: []*model.Repository{
					{
						Name: "1",
					},
				},
			},

			out: &model.Project{
				Name: "1",
				Repositories: []*model.Repository{
					{
						Name: "1",
					},
				},
			},
		},
		{
			old: &model.Project{
				Name: "random",
			},
			new: &model.Project{
				Name: "1",
			},
			out: &model.Project{
				Name: "1",
			},
		},
	}

	for _, test := range tests {
		result := MergeProjects(test.old, test.new)

		if !reflect.DeepEqual(result, test.out) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}

func TestMergeRepositories(t *testing.T) {
	tests := []struct {
		old []*model.Repository
		new []*model.Repository
		out []*model.Repository
	}{
		{
			old: []*model.Repository{
				{
					Name:  "1",
					Owner: "project_1",
					Type:  "GITHUB",
				},
			},
			new: []*model.Repository{
				{
					RepositoryName:  "1",
					RepositoryOwner: "project_1",
					RepositoryType:  "GITHUB",
					Branches: []*model.Branch{
						{
							Name:            "1",
							RepositoryName:  "1",
							RepositoryOwner: "project_1",
							RepositoryType:  "GITHUB",
						},
					},
				},
			},

			out: []*model.Repository{
				{
					Name:  "1",
					Owner: "project_1",
					Type:  "GITHUB",
					Branches: []*model.Branch{
						{
							Name:            "1",
							RepositoryName:  "1",
							RepositoryOwner: "project_1",
							RepositoryType:  "GITHUB",
						},
					},
				},
			},
		},
		{
			old: []*model.Repository{
				{
					Name:  "1",
					Owner: "project_1",
					Type:  "GITHUB",
				},
			},
			new: []*model.Repository{
				{
					RepositoryName:  "2",
					RepositoryOwner: "project_1",
					RepositoryType:  "GITHUB",
				},
			},

			out: []*model.Repository{
				{
					Name:  "1",
					Owner: "project_1",
					Type:  "GITHUB",
				},
				{
					RepositoryName:  "2",
					RepositoryOwner: "project_1",
					RepositoryType:  "GITHUB",
				},
			},
		},
	}

	for _, test := range tests {
		result := MergeRepositories(test.old, test.new)

		if !reflect.DeepEqual(result, test.out) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}

func TestMergeBranches(t *testing.T) {
	tests := []struct {
		old []*model.Branch
		new []*model.Branch
		out []*model.Branch
	}{
		{
			old: []*model.Branch{
				{
					Name: "1",
				},
			},
			new: []*model.Branch{
				{
					Name: "1",
					Commits: []*model.Commit{
						{
							Sha: "1",
						},
					},
				},
			},

			out: []*model.Branch{
				{
					Name: "1",
					Commits: []*model.Commit{
						{
							Sha: "1",
						},
					},
				},
			},
		},
		{
			old: []*model.Branch{
				{
					Name: "random",
				},
			},
			new: []*model.Branch{
				{
					Name: "1",
				},
			},

			out: []*model.Branch{
				{
					Name: "random",
				},
				{
					Name: "1",
				},
			},
		},
	}

	for _, test := range tests {
		result := MergeBranches(test.old, test.new)

		if !reflect.DeepEqual(result, test.out) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}

func TestMergeCommits(t *testing.T) {
	tests := []struct {
		old []*model.Commit
		new []*model.Commit
		out []*model.Commit
	}{
		{
			old: []*model.Commit{
				{
					Sha: "1",
				},
			},
			new: []*model.Commit{
				{
					Sha:     "1",
					Message: "override",
				},
			},

			out: []*model.Commit{
				{
					Sha:     "1",
					Message: "override",
				},
			},
		},
		{
			old: []*model.Commit{
				{
					Sha: "random",
				},
			},
			new: []*model.Commit{
				{
					Sha: "1",
				},
			},

			out: []*model.Commit{
				{
					Sha: "random",
				},
				{
					Sha: "1",
				},
			},
		},
	}

	for _, test := range tests {
		result := MergeCommits(test.old, test.new)

		if !reflect.DeepEqual(result, test.out) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}
