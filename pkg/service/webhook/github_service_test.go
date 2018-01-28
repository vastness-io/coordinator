package webhook

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist-svc/mock/linguist"
	"testing"
)

func TestGetLanguagesUsedInRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		mockClient            *mock_linguist_svc.MockLinguistClient
		expectedDetectedLangs []*linguist.Language
		err                   error
	}{
		{
			mockClient: mock_linguist_svc.NewMockLinguistClient(ctrl),
			expectedDetectedLangs: []*linguist.Language{
				{
					Name:       "Go",
					Percentage: float64(50),
				},
				{
					Name:       "Java",
					Percentage: float64(50),
				},
			},
		},
		{
			mockClient:            mock_linguist_svc.NewMockLinguistClient(ctrl),
			expectedDetectedLangs: []*linguist.Language{},
			err: errors.New("example"),
		},
	}

	for _, test := range tests {

		var (
			log     = logrus.New().WithField("testing", true)
			lingSvc = NewGithubWebhookService(log, test.mockClient)
			langReq = &linguist.LanguageRequest{
				FileNames: []string{
					"main.go",
					"app.java",
				},
			}

			langRes = &linguist.LanguageResponse{
				test.expectedDetectedLangs,
			}
		)
		test.mockClient.EXPECT().GetLanguages(gomock.Any(), langReq).Return(langRes, test.err)

		languages := lingSvc.GetLanguagesUsedInRepository(langReq)

		if len(languages) != len(test.expectedDetectedLangs) {
			t.Fatal("Should equal")
		}

		for _, expected := range test.expectedDetectedLangs {
			var found bool
			for _, actual := range languages {
				if expected == actual {
					found = true
				}
			}
			if !found {
				t.Fatalf("Should have detected %s language", expected)
			}

			found = false
		}
	}
}
