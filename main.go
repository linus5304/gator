package main

import (
	"fmt"
	"os"

	"github.com/linus5304/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handleLogin(s *state, c *command) error {
	if len(c.args) != 1 {
		return fmt.Errorf("usage: login <username>")
	}
	err := s.cfg.SetUser(c.args[0])
	if err != nil {
		return err
	}
	fmt.Println("Logged in as", c.args[0])
	return nil
}

type commands struct {
	value map[string]func(*state, *command) error
}

func (c *commands) register(name string, fn func(*state, *command) error) {
	c.value[name] = fn
}

func (c *commands) run(s *state, cmd *command) error {
	fn, ok := c.value[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return fn(s, cmd)
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}
	s := state{cfg: &cfg}
	commands := commands{value: make(map[string]func(*state, *command) error)}
	commands.register("login", handleLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}
	cmd := command{name: args[1], args: args[2:]}
	err = commands.run(&s, &cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
