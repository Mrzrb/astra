package petstore

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Mrzrb/astra/tests/integration/10-struct-name-collision-avoidance/nested/types"
	topLevelTypes "github.com/Mrzrb/astra/tests/integration/10-struct-name-collision-avoidance/types"
)

func topLevelHandler(c *gin.Context) {
	c.JSON(http.StatusOK, topLevelTypes.TestType{
		TopLevelField: "topLevel",
	})
}

func nestedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, types.TestType{
		NestedField: "nested",
	})
}
