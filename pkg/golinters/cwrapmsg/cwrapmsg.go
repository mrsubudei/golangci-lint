package cwrapmsg

import (
	"github.com/mrsubudei/cwrapmsg"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.CwrapmsgSettings) *goanalysis.Linter {
	cfg := &cwrapmsg.Config{}

	if settings != nil {
		cfg.IgnoreCyrillic = settings.IgnoreCyrillic
		cfg.CaseInsensitive = settings.CaseInsensitive
	}

	return goanalysis.
		NewLinterFromAnalyzer(cwrapmsg.NewAnalyzer(cfg)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
