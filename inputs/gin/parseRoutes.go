package gin

import (
	"sync"

	"github.com/Mrzrb/astra"
)

// ParseRoutes parses routes from a gin routes.
// It will populate the routes with the handler function.
// It will individually call parseRoute for each route.
// createRoutes must be called before this.
func ParseRoutes() astra.ServiceFunction {
	return func(s *astra.Service) error {
		var wg sync.WaitGroup
		s.Log.Debug().Msg("Populating routes from gin routes")
		var mutex sync.Mutex
		for _, route := range s.Routes {
			wg.Add(1)
			go func(s *astra.Service, route astra.Route) error {
				s.Log.Debug().Str("path", route.Path).Str("method", route.Method).Msg("Populating route")
				s.Log.Debug().Str("path", route.Path).Str("method", route.Method).Str("file", route.File).Int("line", route.LineNo).Msg("Parsing route")
				defer wg.Done()

				mutex.Lock()
				err := parseRoute(s, &route)
				mutex.Unlock()
				if err != nil {
					s.Log.Error().Str("path", route.Path).Str("method", route.Method).Str("file", route.File).Int("line", route.LineNo).Err(err).Msg("Failed to parse route")
					return err
				}

				s.ReplaceRoute(route)
				return nil
			}(s, route)
		}
		wg.Wait()
		s.Log.Debug().Msg("Populated service with gin routes")

		return nil
	}
}
