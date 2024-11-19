package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	commandName map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandName[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	val, ok := c.commandName[cmd.name]
	if !ok {
		return errors.New("command does not exist")
	}
	err := val(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
