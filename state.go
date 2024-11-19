package main

import (
	"github.com/Ephim135/RSS-feed-Aggregator/internal/config"
	"github.com/Ephim135/RSS-feed-Aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
