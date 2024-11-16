package main

import (
	"github.com/Ephim135/RSS-feed-Aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	s := state{cfg: cfg}
}
