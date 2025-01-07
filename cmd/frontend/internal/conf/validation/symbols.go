package validation

import (
	"fmt"

	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/ctags_config"
	"github.com/khulnasoft/khulnasoft/lib/codeintel/languages"
)

func init() {
	conf.ContributeValidator(func(c conftypes.SiteConfigQuerier) (problems conf.Problems) {
		configuration := c.SiteConfig().SyntaxHighlighting
		if configuration == nil || configuration.Symbols == nil {
			return nil
		}

		for language, engine := range configuration.Symbols.Engine {
			parser_engine, err := ctags_config.ParserNameToParserType(engine)
			if err != nil {
				return conf.NewSiteProblems(fmt.Sprintf("Not a valid Symbols.Engine: `%s`.", engine))
			}

			language = languages.NormalizeLanguage(language)
			if !ctags_config.LanguageSupportsParserType(language, parser_engine) {
				return conf.NewSiteProblems(fmt.Sprintf("Not a valid Symbols.Engine for language: %s `%s`.", language, engine))
			}

		}

		return nil
	})
}
