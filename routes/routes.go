package routes

import (
	"github.com/stef-k/xrplnsapi/controllers"

	"github.com/labstack/echo"
)

// InitRoutes initializes all routes and handlers
func InitRoutes(e *echo.Echo) {
	// setup a group of routes with common prefix v1
	g := e.Group("/v1")
	// setup routes
	g.GET("/social-networks", controllers.SocialNetworks)
	g.GET("/resolve/social/:network/:username", controllers.Resolve)
	g.GET("/resolve/xrplaccount/:xrplaccount/:tag", controllers.ResolveXRPLAccount)
	g.GET("/resolve/user/:slug", controllers.ResolveUser)
}
