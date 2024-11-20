package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/linus5304/gator/internal/config"
	"github.com/linus5304/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handleLogin)
	commands.register("register", handleRegister)
	commands.register("reset", handleReset)
	commands.register("users", handleGetUsers)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}
	cmd := command{name: args[1], args: args[2:]}
	err = commands.run(&s, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
