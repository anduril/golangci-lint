package golinters

import (
	mnd "github.com/tommy-muehle/go-mnd"
	"golang.org/x/tools/go/analysis"

	"github.com/anduril/golangci-lint/pkg/config"
	"github.com/anduril/golangci-lint/pkg/golinters/goanalysis"
)

func NewGoMND(cfg *config.Config) *goanalysis.Linter {
	analyzers := []*analysis.Analyzer{
		mnd.Analyzer,
	}

	var linterCfg map[string]map[string]interface{}
	if cfg != nil {
		linterCfg = cfg.LintersSettings.Gomnd.Settings
	}

	return goanalysis.NewLinter(
		"gomnd",
		"An analyzer to detect magic numbers.",
		analyzers,
		linterCfg,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
