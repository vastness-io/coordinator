package webhook

import (
	"github.com/vastness-io/linguist-svc"
)

type Service interface {
	GetLanguagesUsedInRepository(*linguist.LanguageRequest) []*linguist.Language
}
