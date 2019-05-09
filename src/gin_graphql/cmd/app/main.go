package main

import (
	graphql "gin_graphql"
	"gin_graphql/internal/middlewares"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	r.Use(middlewares.GinContextToContextMiddleware())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
