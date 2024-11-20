package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/linus5304/gator/internal/database"
)

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset users: %w", err)
	}
	fmt.Println("Users reset successfully")
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <username>", cmd.name)
	}
	name := cmd.args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}
	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <username>", cmd.name)
	}
	name := cmd.args[0]
	user, err := s.db.GetUserByName(context.Background(), name)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}
	fmt.Println("User switched successfully:")
	printUser(user)
	return nil
}

func handleGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get users: %w", err)
	}
	printUsers(users, s)
	return nil
}

func printUsers(users []database.User, s *state) {
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
}

func printUser(user database.User) {
	fmt.Printf(" * ID: 		%v\n", user.ID)
	fmt.Printf(" * Name: 	%v\n", user.Name)
}
