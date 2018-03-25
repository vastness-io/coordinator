package project

import "github.com/vastness-io/coordinator-svc/project"

type Service interface {
	GetProject(string, string) (*project.Project, error)
	GetProjects() []*project.Project
}
