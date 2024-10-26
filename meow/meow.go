package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	app := &cli.App{
		Name:  "Say meow",
		Usage: "Just say meow",
		Action: func(ctx *cli.Context) error {
			meows := [...]string{"meow", "murr", "maow", "meoww", "nyaaa", "nya"}
			if ctx.NArg() == 0 {
				return fmt.Errorf("you need to set amount of meow")
			}
			repeats, err := strconv.Atoi(ctx.Args().First())
			if err != nil {
				return fmt.Errorf("args need to be an int")
			}
			for i := 0; i < repeats; i++ {
				fmt.Printf("%s ", meows[rand.Intn(len(meows))])
			}
			if rand.Intn(4) > 2 {
				fmt.Print(":3")
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
