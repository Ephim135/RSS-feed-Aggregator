package main

import (
	"github.com/Ephim135/RSS-feed-Aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
