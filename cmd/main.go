package main

import (
	"github.com/gl00my4eyes/mygoftp/internal/config"
)

func main() {
	cfg := config.NewFromFile("./config/config.yml")

	cfg.Host()

}
