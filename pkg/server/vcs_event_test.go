package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/vcs-webhook-svc/webhook"
	"testing"
	"github.com/vastness-io/coordinator/pkg/model"
)

type NOOPService struct{}

func (n *NOOPService) GetLanguagesUsedInRepository(_ *linguist.LanguageRequest) []*linguist.Language {
	var result []*linguist.Language
	return result
}

func (n *NOOPService) UpdateProject(project *model.Project) (*model.Project, error) {
	return &model.Project{}, nil
}

func TestOnPush(t *testing.T) {

	var (
		log = logrus.New().WithField("testing", true)
		svc = &NOOPService{}
		srv = NewVcsEventServer(svc, log)
	)

	tests := []struct {
		req *vcs.VcsPushEvent
		ctx context.Context
	}{
		{
			req: &vcs.VcsPushEvent{},
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
