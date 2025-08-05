package main

import (
	"errors"
	"fmt"
	"context"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
func resetCommand(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("All users deleted successfully!")
	return nil
}

func usersCommand(s *state, cmd command) error {
	users, err := s.queries.GetUsers(CmdContext(cmd))
	if err != nil {
		return err
	}
	if len(users) == 0 {
		fmt.Println("No users found.")
		return nil
	}

	//Pretty-print each user as JSON
	for _, u := range users{
		data, err := json.MarshalIndent(u, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to arsh user: %w", err)
		}
		fmt.Println(string(data))
	}
	return nil
}
