package main

import "errors"

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, fn func(*state, command) error) {
	c.registeredCommands[name] = fn
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("command not found: " + cmd.name)
	}
	return fn(s, cmd)
}
