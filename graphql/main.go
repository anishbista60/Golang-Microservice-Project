package graphql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL string `envconfig:"ACCCOUNT_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func maini() {
	var cfg AppConfig
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}

	server, err := NewGraphqlServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.GraphQL(server.ToExcecutableSchema()))
	http.Handle("/playground", handler.GraphQL())

}
