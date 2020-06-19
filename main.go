package main

import (
	"buncms/graph"
	"buncms/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bunred/buncms"
	"github.com/bunred/buncms/types"
	"github.com/gofiber/fiber"
)

func main() {
	// buncms
	server := buncms.NewServer(types.NewServer{
		EnableCORS: true,
	})

	server.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!111")
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	serverHandler := srv.Handler()
	playgroundHandler := playground.Handler("GraphQL playground", "/query")

	server.Use("/query", func(c *fiber.Ctx) {
		serverHandler(c.Fasthttp)
	})

	server.Use("/playground", func(c *fiber.Ctx) {
		playgroundHandler(c.Fasthttp)
	})

	buncms.StartServer(types.StartServer{
		Server: server,
	})
}
