package main

import (
	"context"
	"fmt"

	"github.com/linus5304/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not logged in: %w", err)
		}
		return handler(s, cmd, user)
	}
}
