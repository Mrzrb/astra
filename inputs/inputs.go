package inputs

import (
	"github.com/gin-gonic/gin"
	"github.com/Mrzrb/astra"
	astraGin "github.com/Mrzrb/astra/inputs/gin"
)

const (
	InputModeGin astra.InputMode = "gin" // github.com/gin-gonic/gin web framework.
)

func addInput(mode astra.InputMode, createRoutes astra.ServiceFunction, parseRoutes astra.ServiceFunction) astra.Option {
	return func(s *astra.Service) {
		s.Inputs = append(s.Inputs, astra.Input{
			Mode:         mode,
			CreateRoutes: createRoutes,
			ParseRoutes:  parseRoutes,
		})
	}
}

// WithGinInput adds gin as an input to the service.
// CreateRoutes is called before ParseRoutes.
// CreateRoutes is the only function that will have access to the routes - it will create the routes and refer to the handler function by name, file and line number.
// ParseRoutes will populate the routes with the handler function, should not need access to the routes because there will be cases where the routes is nil (CLI).
func WithGinInput(router *gin.Engine) astra.Option {
	return addInput(
		InputModeGin,
		astraGin.CreateRoutes(router),
		astraGin.ParseRoutes(),
	)
}
