package main

import "fmt"

type Command rune

func (c Command) String() string {
	return fmt.Sprintf("%c", c)
}

type CommandLoop struct {
	commands []Command
	counter int
}

func (c *CommandLoop) Next() Command {
	commandIdx := c.counter
	c.counter = (c.counter + 1) % len(c.commands)
	return c.commands[commandIdx]
}

func (c *CommandLoop) Reset() {
	c.counter = 0
}

const (
	Left = Command('L')
	Right = Command('R')
)