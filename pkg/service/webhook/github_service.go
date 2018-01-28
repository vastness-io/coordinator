package webhook

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist-svc"
	"time"
)

type githubWebhookService struct {
	linguist linguist.LinguistClient
	logger   *logrus.Entry
}

func NewGithubWebhookService(logger *logrus.Entry, linguistClient linguist.LinguistClient) Service {
	return &githubWebhookService{
		linguist: linguistClient,
		logger:   logger,
	}
}

func (s *githubWebhookService) GetLanguagesUsedInRepository(req *linguist.LanguageRequest) []*linguist.Language {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Don't think this is the right way to do it?
	defer cancel()

	res, err := s.linguist.GetLanguages(ctx, req)

	if err != nil {
		s.logger.Error("Unable to detect Language(s) for repository")
		return make([]*linguist.Language, 0)
	}

	return res.GetLanguage()

}
