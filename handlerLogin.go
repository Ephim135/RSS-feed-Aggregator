package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("expecting one argument username")
	}
	if len(cmd.args) != 1 {
		return errors.New("to many arguments")
	}

	s.cfg.CurrentUserName = cmd.args[0]
	fmt.Println("User has been set!")
}
