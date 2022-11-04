package main

import (
	"flag"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"swagger-gin-simple/api"
	"swagger-gin-simple/dbs"
	"swagger-gin-simple/seeds"
)

var (
	laddr           = flag.String("addr", ":3000", "Local address for the HTTP API")
	loglevel        = flag.String("loglevel", "INFO", "Log-level (ERROR|WARN|INFO|DEBUG|TRACE)")
	initialSeedFile = flag.String("initialSeedFile", "", "Run one-time seeds passing path to a valid JSON seed file")
)

func configureLogging() error {
	l, err := log.ParseLevel(*loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(l)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableHTMLEscape: true,
	})
	return nil
}

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	flag.Parse()
	if err := configureLogging(); err != nil {
		log.Fatal(err)
	}

	// DB env
	os.Setenv("DB","sqlite")
	os.Setenv("DATABASE_PATH","./db.sqlite3")

	err := dbs.InitializeDatabaseLayer()
	if err != nil {
		log.Fatal(err)
	}

	if initialSeedFile != nil && *initialSeedFile != "" {
		if err = seeds.RunSeeds(*initialSeedFile); err != nil {
			log.Fatal(err)
		}
	}

    // Gin instance
	api := api.NewRESTApiV1()
	if err = api.Serve(*laddr); err != nil {
		log.Fatal(err)
	}
}

