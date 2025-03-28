package main

import (
	"auth/auth/xauth"
	"auth/auth/xnotificaciones"
	"auth/graph"
	"auth/graph_auth"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func conexion() *sql.DB {
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	// loc := "America%2FLa_Paz"
	loc := "UTC"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=%s", dbuser, dbpass, dbhost, dbname, loc)
	db, err := sql.Open("mysql", dsn)
	fmt.Println(dsn)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	er := db.Ping()
	if er != nil {
		panic(er.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

const defaultPort = "8020"

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := conexion()
	router := chi.NewRouter()
	router.Use(xauth.AuthMiddleware(db))
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	// subs := make(map[string]chan *model.XNotificacion)
	xnotificaciones.InitializeGlobal()
	resolver_auth := &graph_auth.Resolver{DB: db}
	srv_auth := handler.NewDefaultServer(graph_auth.NewExecutableSchema(graph_auth.Config{Resolvers: resolver_auth}))

	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// websocket
	srvws := handler.New(graph_auth.NewExecutableSchema(graph_auth.Config{Resolvers: resolver_auth}))
	srvws.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 40 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: xauth.UaserIDMiddleware(db),
	})
	// fin websocket

	show_playground := os.Getenv("PLAYGROUND")
	if show_playground == "1" {
		router.Handle("/auth", playground.Handler("GraphQL playground", "/query_auth"))
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	router.Handle("/query_auth", srv_auth)
	router.Handle("/query", srv)
	router.Handle("/ws", srvws)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
