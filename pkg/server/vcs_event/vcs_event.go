package vcs_event

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/coordinator/pkg/service/vcs_event"
	"github.com/vastness-io/vcs-webhook-svc/webhook"
)

type vcsEventServer struct {
	service vcs_event.Service
	log     *logrus.Entry
}

func NewVcsEventServer(service vcs_event.Service, logger *logrus.Entry) vcs.VcsEventServer {
	return &vcsEventServer{
		service: service,
		log:     logger,
	}
}

func (s *vcsEventServer) OnPush(ctx context.Context, event *vcs.VcsPushEvent) (*empty.Empty, error) {

	res := new(empty.Empty)

	project := ConvertToProjectModel(event)

	project, err := s.service.UpdateProject(project)

	if err != nil {
		return res, err
	}

	return res, nil

}
