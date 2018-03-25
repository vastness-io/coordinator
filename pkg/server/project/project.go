package project

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	project_message "github.com/vastness-io/coordinator-svc/project"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/service/project"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectsInformationServer struct {
	service project.Service
	log     *logrus.Entry
}

func NewProjectInformationServer(service project.Service, logger *logrus.Entry) project_message.ProjectsServer {
	return &ProjectsInformationServer{
		service: service,
		log:     logger,
	}
}

func (p *ProjectsInformationServer) GetProject(ctx context.Context, message *project_message.GetProjectMessage) (*project_message.Project, error) {
	prj, err := p.service.GetProject(message.Name, message.Type)

	if err != nil {
		if err == errors.ProjectDoesNotExistErr {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	return prj, nil
}

func (p *ProjectsInformationServer) GetProjects(ctx context.Context, empty *empty.Empty) (*project_message.GetProjectsResponse, error) {
	return &project_message.GetProjectsResponse{
		Projects: p.service.GetProjects(),
	}, nil
}
