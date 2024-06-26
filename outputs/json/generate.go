package json

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/Mrzrb/astra"
)

// JSONOutput is the output of the JSON output.
// It will in essence be a copy of the service's output (routes and components).
type JSONOutput struct {
	Routes     []astra.Route `json:"routes"`
	Components []astra.Field `json:"components"`
}

// Generate will create the JSON output.
// It will marshal the JSONOutput struct and write it to a file.
func Generate(filePath string) astra.ServiceFunction {
	return func(s *astra.Service) error {
		s.Log.Info().Msg("Generating JSON output")
		output := JSONOutput{
			Routes:     s.Routes,
			Components: s.Components,
		}

		s.Log.Debug().Msg("Generated JSON output")
		file, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			s.Log.Error().Err(err).Msg("Failed to marshal JSON output")
			return err
		}

		if !strings.HasSuffix(filePath, ".json") {
			s.Log.Debug().Str("filePath", filePath).Msg("Adding .json suffix to file path")
			filePath += ".json"
		}

		s.Log.Debug().Str("filePath", filePath).Msg("Writing JSON output to file")
		filePath = path.Join(s.WorkDir, filePath)
		err = os.WriteFile(filePath, file, 0644)
		if err != nil {
			s.Log.Error().Err(err).Msg("Failed to write JSON output to file")
			return err
		}

		s.Log.Info().Msg("Generated JSON output")
		return nil
	}
}
