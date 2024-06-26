package azureFunctions

import (
	"strings"

	"github.com/Mrzrb/astra"
)

// Adapted from https://learn.microsoft.com/en-us/aspnet/web-api/overview/web-api-routing-and-actions/attribute-routing-in-web-api-2#route-constraints
var acceptedTypeMap = map[string]string{
	"string":  "alpha",
	"int":     "long",
	"int32":   "int",
	"int64":   "long",
	"uint":    "long",
	"uint32":  "int",
	"uint64":  "long",
	"float":   "double",
	"float32": "float",
	"float64": "double",
	"bool":    "bool",

	// Passthrough types
	"time.Time":                   "datetime",
	"github.com/google/uuid.UUID": "guid",
}

func convertRoute(route astra.Route) string {
	routeString := route.Path

	for _, pathParams := range route.PathParams {
		if azureType, ok := acceptedTypeMap[pathParams.Field.Type]; ok {
			if azureType == "" {
				return ""
			}

			if pathParams.IsRequired {
				routeString = strings.Replace(routeString, ":"+pathParams.Name, "{"+pathParams.Name+":"+azureType+"}", 1)
			} else {
				routeString = strings.Replace(routeString, "*"+pathParams.Name, "{"+pathParams.Name+":"+azureType+"?}", 1)
			}
		}
	}

	return routeString
}
