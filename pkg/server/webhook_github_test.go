package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/vcs-webhook-svc/webhook/github"
	"testing"
)

type NOOPService struct{}

func (n *NOOPService) GetLanguagesUsedInRepository(_ *linguist.LanguageRequest) []*linguist.Language {
	var result []*linguist.Language
	return result
}

func TestOnPush(t *testing.T) {

	var (
		log = logrus.New().WithField("testing", true)
		svc = &NOOPService{}
		srv = NewGithubWebhookServer(svc, log)
	)

	tests := []struct {
		req *github.PushEvent
		ctx context.Context
	}{
		{
			req: &github.PushEvent{
				HeadCommit: &github.PushCommit{
					Added:    []string{},
					Modified: []string{},
					Removed:  []string{},
				},
			},
			ctx: context.Background(),
		},
	}

	for _, test := range tests {
		result, err := srv.OnPush(test.ctx, test.req)

		expected := new(empty.Empty)

		if result != expected || err != nil {
			t.Fatalf("expected pointer of empty, got: %s", expected)
		}
	}

}
