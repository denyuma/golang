// main.go
package main

import (
<<<<<<< HEAD
	"todo/delivery"
	"todo/repository"
	"todo/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// repositoryをインスタンス化
	tr := repository.NewSyncMapTodoRepository()
	// usecaseをインスタンス化
	tu := usecase.NewTodoUsecase(tr)

	// fiberをインスタンス化
	c := fiber.New()

	// CORSの設定
	c.Use(cors.New(cors.Config{
		// https://docs.gofiber.io/api/middleware/cors#config
		AllowCredentials: true,
	}))

	delivery.NewTodoAllGetHandler(c, tu)
	delivery.NewTodoDeleteHandler(c, tu)
	delivery.NewTodoStatusUpdateHandler(c, tu)
	delivery.NewTodoStoreHandler(c, tu)

	c.Listen(":8000")
=======
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var (
		suffix string
	)
	app := cli.NewApp()
	app.Name = "Hello xxxx"
	app.Usage = "Make `Hello xxx` for arbitrary text"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "suffix, s",
			Value:       "!!!",
			Usage:       "text after speaking something",
			Destination: &suffix,
			EnvVar:      "SUFFIX",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "hello",
			Usage: "if use set -t or --text",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "text, t",
					Value: "world",
					Usage: "hello xxx text",
				},
			},

			Action: func(c *cli.Context) error {
				fmt.Printf("Hello %s %s\n", c.String("text"), suffix)
				return nil
			},
		},
	}

	app.Run(os.Args)
>>>>>>> commandline-mingo
}
