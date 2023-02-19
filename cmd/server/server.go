package main

import (
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"
)

func main() {
	fx.New().Run()
}
