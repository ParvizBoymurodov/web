package main

import (
	"context"
	"flag"
	"github.com/ParvizBoymurodov/web/cmd/crud/app"
	"github.com/ParvizBoymurodov/web/pkg/crud/services/burgers"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var (
	host = flag.String("host", "", "Server host")
	port = flag.String("port", "", "Server port")
	dsn  = flag.String("dsn", "", "Postgres DSN")
)

const (
	envHost = "HOST"
	envPort = "PORT"
	envDSN  = "DATABASE_URL"
)


func fromFLagOrEnv(flag *string, envName string) (server string, ok bool){
	if *flag != ""{
		return *flag, true
	}
	return os.LookupEnv(envName)
}

func main() {
	flag.Parse()
	hostf, _ := fromFLagOrEnv(host, envHost)
	portf, _ := fromFLagOrEnv(port, envPort)
	dsnf, _ := fromFLagOrEnv(dsn, envDSN)

	addr := net.JoinHostPort(hostf, portf)
	start(addr, dsnf)
}

func start(addr string, dsn string) {
	router := app.NewExactMux()
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	burgersSvc := burgers.NewBurgersSvc(pool)
	server := app.NewServer(
		router,
		pool,
		burgersSvc,
		filepath.Join("web", "templates"),
		filepath.Join("web", "assets"),
	)
	server.InitRoutes()

	panic(http.ListenAndServe(addr, server))
}




//;lkmascxkscsbjdscdskncvdsklc;sdjk