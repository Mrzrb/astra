package cli

import "github.com/Mrzrb/astra"

// WithCLI enables the CLI mode for the generator
// This will run the generator to only crawl for the file names, function names and line numbers of the functions that need to be analysed.
func WithCLI() astra.Option {
	return func(s *astra.Service) {
		s.CacheEnabled = true
		s.CLIMode = astra.CLIModeSetup
	}
}

// WithCLIBuilder enables the CLI mode for the generator
// This will run the generator utilising existing cache for the file names, function names and line numbers of the functions that need to be analysed to generate the code.
func WithCLIBuilder() astra.Option {
	return func(s *astra.Service) {
		s.CLIMode = astra.CLIModeBuilder
	}
}
