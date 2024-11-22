package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Ephim135/RSS-feed-Aggregator/internal/config"
	"github.com/Ephim135/RSS-feed-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: cfg,
	}

	cmds := commands{commandName: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", reset)
	cmds.register("users", users)
	cmds.register("agg", agg)
	cmds.register("addfeed", addfeed)
	cmds.register("feeds", feed)
	cmds.register("follow", follow)
	cmds.register("following", following)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	commandName := args[1]
	var cmdArgs []string
	if len(args) > 2 {
		cmdArgs = args[2:]
	} else {
		cmdArgs = []string{}
	}

	cmd := command{name: commandName, args: cmdArgs}
	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
