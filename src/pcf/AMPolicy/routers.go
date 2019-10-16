/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AMPolicy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/npcf-am-policy-control/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"PoliciesPolAssoIdDelete",
		strings.ToUpper("Delete"),
		"/policies/:polAssoId",
		PoliciesPolAssoIdDelete,
	},

	{
		"PoliciesPolAssoIdGet",
		strings.ToUpper("Get"),
		"/policies/:polAssoId",
		PoliciesPolAssoIdGet,
	},

	{
		"PoliciesPolAssoIdUpdatePost",
		strings.ToUpper("Post"),
		"/policies/:polAssoId/update",
		PoliciesPolAssoIdUpdatePost,
	},

	{
		"PoliciesPost",
		strings.ToUpper("Post"),
		"/policies",
		PoliciesPost,
	},
}