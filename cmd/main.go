package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jhinmainksta/habr-clone/graph"
	"github.com/jhinmainksta/habr-clone/graph/my_model"
	"github.com/jhinmainksta/habr-clone/repository"
	"github.com/jhinmainksta/habr-clone/repository/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	subs := make(map[string]map[string]chan *my_model.Comment)

	repo := repository.NewRepository(db)
	resolver := graph.NewResolver(repo, subs, viper.GetInt("limit"), viper.GetInt("offset"))

	queryHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	queryHandler.AddTransport(&transport.Websocket{})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataloaderMiddleware(db, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}

func InitConfig() error {
	viper.AddConfigPath("cfgs")
	viper.SetConfigName("cfg")
	return viper.ReadInConfig()
}
