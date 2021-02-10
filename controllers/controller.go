package controllers

import (
	"awacs_smart_api_service/graph"
	"awacs_smart_api_service/graph/generated"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brkelkar/common_utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//LogResponceEvent responce logger
func logUA(ua string) {
	logger.Info("UA_details", zap.String("UA", ua))

}

//GraphqlHandler Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		c.Request.UserAgent()
		logUA(c.Request.UserAgent())
		h.ServeHTTP(c.Writer, c.Request)
	}
}

//PlaygroundHandler Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		URI := c.Request.RequestURI
		userAgent := c.Request.UserAgent()
		fmt.Println(URI, userAgent)
		h.ServeHTTP(c.Writer, c.Request)
	}
}
