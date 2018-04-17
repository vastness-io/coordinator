package vcs_event

import (
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/vcs-webhook-svc/webhook"
	"reflect"
	"testing"
	"time"
)

func TestRemoveRefPrefix(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "refs/head/branch",
			out: "refs/head/branch",
		},
		{
			in:  "refs/heads/branch",
			out: "branch",
		},
		{
			in:  "refs/tag/branch",
			out: "refs/tag/branch",
		},
		{
			in:  "refs/tags/branch",
			out: "branch",
		},
	}

	for _, test := range tests {
		result := RemoveRefPrefix(test.in)

		if result != test.out {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}

func TestConvertEventRepositoryToRepositoryModel(t *testing.T) {
	tests := []struct {
		in  *vcs.Repository
		out *model.Repository
	}{
		{
			in: &vcs.Repository{
				Name: "name",
			},
			out: &model.Repository{
				RepositoryName: "name",
			},
		},
	}

	for _, test := range tests {
		result := ConvertEventRepositoryToRepositoryModel(test.in)

		if !reflect.DeepEqual(result, test.out) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}

func TestConvertEventCommitToCommitModel(t *testing.T) {
	tests := []struct {
		in  *vcs.PushCommit
		out *model.Commit
	}{
		{
			in: &vcs.PushCommit{
				Sha:       "hash",
				Id:        "hash",
				TreeId:    "hash",
				Distinct:  true,
				Message:   "some message",
				Timestamp: "2015-05-05T19:40:15-04:00",
				Url:       "url",
				Author: &vcs.CommitAuthor{
					Name:     "name",
					Email:    "example@example.com",
					Username: "name",
				},
				Committer: &vcs.CommitAuthor{
					Name:     "name",
					Email:    "example@example.com",
					Username: "name",
				},
				Added: []string{
					"file",
				},
				Modified: []string{
					"file2",
				},
				Removed: []string{},
			},
			out: &model.Commit{
				Sha:         "hash",
				Timestamp:   toTime("2015-05-05T19:40:15-04:00"),
				Message:     "some message",
				AuthorName:  "name",
				AuthorEmail: "example@example.com",
				Added: []string{
					"file",
				},
				Modified: []string{
					"file2",
				},
				Removed: []string{},
			},
		},
	}

	for _, test := range tests {
		result := ConvertEventCommitToCommitModel(test.in)

		if !reflect.DeepEqual(result, test.out) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}

func TestConvertToProjectModel(t *testing.T) {
	tests := []struct {
		in  *vcs.VcsPushEvent
		out *model.Project
	}{
		{
			in: &vcs.VcsPushEvent{
				Ref: "refs/heads/branch",
				Commits: []*vcs.PushCommit{
					{
						Sha:       "hash",
						Id:        "hash",
						TreeId:    "hash",
						Distinct:  true,
						Message:   "some message",
						Timestamp: "2015-05-05T19:40:15-04:00",
						Url:       "url",
						Author: &vcs.CommitAuthor{
							Name:     "name",
							Email:    "example@example.com",
							Username: "name",
						},
						Committer: &vcs.CommitAuthor{
							Name:     "name",
							Email:    "example@example.com",
							Username: "name",
						},
						Added: []string{
							"file",
						},
						Modified: []string{
							"file2",
						},
						Removed: []string{},
					},
				},
				Repository: &vcs.Repository{
					Owner: &vcs.User{
						Login: "project_1",
						Url:   "URL",
						Type:  "GITHUB",
						Name:  "project_1",
						Email: "example@example.com",
					},
					Name:     "repo_1",
					FullName: "project_1/repo_1",
					Organization: &vcs.User{
						Login: "project_1",
						Url:   "URL",
						Type:  "GITHUB",
						Name:  "project_1",
						Email: "example@example.com",
					},
				},
				Organization: &vcs.User{
					Login: "project_1",
					Url:   "URL",
					Type:  "GITHUB",
					Name:  "project_1",
					Email: "example@example.com",
				},
				VcsType: vcs.VcsType_GITHUB,
			},
			out: &model.Project{
				Name: "project_1",
				Type: "GITHUB",
				Repositories: []*model.Repository{
					{
						RepositoryName:  "repo_1",
						RepositoryOwner: "project_1",
						RepositoryType:  "GITHUB",
						Branches: []*model.Branch{
							{
								Name:            "branch",
								RepositoryName:  "repo_1",
								RepositoryOwner: "project_1",
								RepositoryType:  "GITHUB",
								Meta:            make(model.BranchMeta),
								Commits: []*model.Commit{
									{
										Sha:         "hash",
										Timestamp:   toTime("2015-05-05T19:40:15-04:00"),
										Message:     "some message",
										AuthorName:  "name",
										AuthorEmail: "example@example.com",
										Added: []string{
											"file",
										},
										Modified: []string{
											"file2",
										},
										Removed: []string{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := ConvertToProjectModel(test.in)

		if !reflect.DeepEqual(result, test.out) {
			t.Fatalf("Expected %v, got %v", test.out, result)
		}
	}
}

func toTime(s string) *time.Time {
	t, _ := time.Parse(time.RFC3339, s)

	t = t.UTC()

	return &t
}
