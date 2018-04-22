package project

import (
	"github.com/vastness-io/coordinator-svc/project"
	"github.com/vastness-io/coordinator/pkg/model"
	"reflect"
	"testing"
	"time"
)

func TestFromProjectModel(t *testing.T) {

	var (
		testTime    = time.Time{}
		testTimePtr = &testTime
	)
	tests := []struct {
		in       *model.Project
		expected *project.Project
	}{
		{
			in: &model.Project{
				Name: "project_1",
				Type: "GITHUB",
				Repositories: []*model.Repository{
					{
						Name:  "repo_1",
						Owner: "project_1",
						Type:  "GITHUB",
						Branches: []*model.Branch{
							{
								Name: "branch_1",
								Meta: map[string]interface{}{
									"languages": map[string]interface{}{
										"Go": 15.0,
									},
								},
								Commits: []*model.Commit{
									{
										Sha:         "something",
										Message:     "some message",
										Timestamp:   testTimePtr,
										AuthorName:  "author name",
										AuthorEmail: "author email",
										Added: []string{
											"added",
										},
										Modified: []string{
											"modified",
										},
										Removed: []string{
											"removed",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: &project.Project{
				Name: "project_1",
				Type: "GITHUB",
				Repositories: []*project.Repository{
					{
						Name:  "repo_1",
						Owner: "project_1",
						Type:  "GITHUB",
						Branches: []*project.Branch{
							{
								Name: "branch_1",
								Meta: &project.BranchMeta{
									Languages: []*project.Language{
										{
											Name:       "Go",
											Percentage: 15.0,
										},
									},
								},
								Commits: []*project.Commit{
									{
										Sha:         "something",
										Message:     "some message",
										Timestamp:   testTime.String(),
										AuthorName:  "author name",
										AuthorEmail: "author email",
										Added: []string{
											"added",
										},
										Modified: []string{
											"modified",
										},
										Removed: []string{
											"removed",
										},
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
		out := FromProjectModel(test.in)

		if !reflect.DeepEqual(test.expected, out) {
			t.Fatalf("Expected %v, got %v", test.expected, out)
		}
	}

}
