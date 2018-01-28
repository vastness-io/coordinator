package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	service "github.com/vastness-io/coordinator/pkg/service/webhook"
	"github.com/vastness-io/coordinator/pkg/util"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/vcs-webhook-svc/webhook/github"
)

type githubWebhookServer struct {
	service service.Service
	log     *logrus.Entry
}

func NewGithubWebhookServer(service service.Service, logger *logrus.Entry) github.GithubWebhookServer {
	return &githubWebhookServer{
		service: service,
		log:     logger,
	}
}

func (s *githubWebhookServer) OnPush(ctx context.Context, req *github.PushEvent) (*empty.Empty, error) {

	go func(req *github.PushEvent) {
		langReq := s.constructLangRequest(req)
		languages := s.service.GetLanguagesUsedInRepository(langReq)
		s.log.Info(languages)
	}(req)

	return new(empty.Empty), nil

}

func (s *githubWebhookServer) constructLangRequest(req *github.PushEvent) *linguist.LanguageRequest {
	var (
		headCommit    = req.GetHeadCommit()
		filesInCommit = util.MergeStringSlices(headCommit.GetAdded(), headCommit.Modified)
	)

	return &linguist.LanguageRequest{
		FileNames: filesInCommit,
	}
}
