package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/server"
	"gopkg.in/urfave/cli.v1"
)

var config = server.NewConfig()

func run(c *cli.Context) error {
	if err := config.Validate(); err != nil {
		return err
	}

	if config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugf("Running with config: %s", config.ToString())
	}

	config.SetFormatter()

	s, err := server.NewServer(config)
	if err != nil {
		return err
	}

	return s.Listen()
}

func main() {
	app := cli.NewApp()
	app.Name = "restaurant-server"
	app.Authors = []cli.Author{
		{Name: "Umut Özdoğan", Email: "umut.ozdgan@gmail.com"},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "d, debug",
			Destination: &config.Debug,
			EnvVar:      "RESTAURANT_DEBUG",
			Usage:       "Enables logs with level DEBUG to be printed to stdout.",
		},
		cli.IntFlag{
			Name:        "p, port",
			Value:       config.Port,
			Destination: &config.Port,
			EnvVar:      "RESTAURANT_PORT",
			Usage:       "The port to listen to for HTTP interface.",
		},
		cli.StringFlag{
			Name:        "dburl",
			Destination: &config.DbURL,
			EnvVar:      "RESTAURANT_DB_URL",
			Usage:       "The URL of the mongodb server that the server will connect to.",
		},
		cli.StringFlag{
			Name:        "dbname",
			Destination: &config.DbName,
			EnvVar:      "RESTAURANT_DB_NAME",
			Usage:       "The name of the mongodb database that the server will use.",
		},
		cli.StringFlag{
			Name:        "logrus-formatter",
			Value:       config.LogrusFormatter,
			Destination: &config.LogrusFormatter,
			EnvVar:      "RESTAURANT_LOGRUS_FORMATTER",
			Usage:       "The name of the formatter used for logging.",
		},
		cli.DurationFlag{
			Name:        "db-cron-interval",
			Value:       config.DBCronInterval,
			Destination: &config.DBCronInterval,
			EnvVar:      "RESTAURANT_DB_CRON_INTERVAL",
			Usage:       "The length of the intervals that db deletes expired carts (more than 14 days old).",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("Server stopped with error")
	}
}
