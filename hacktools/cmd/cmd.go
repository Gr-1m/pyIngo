package cmd

import (
	"bytes"
	"os/exec"
	"strings"
)

type Cmder interface {
	Run() error
}

type Cmd struct {
	Command []string

	Stdout, Stderr bytes.Buffer
}

func (c *Cmd) Run() error {
	cmd := exec.Command(c.Command[0], c.Command[1:]...)
	cmd.Stdout = &c.Stdout
	cmd.Stderr = &c.Stderr

	return cmd.Run()
}

func Run(cmd string) (Stdout, Stderr string) {

	c := &Cmd{
		Command: strings.Split(cmd, " "),
	}

	err := c.Run()
	if err != nil {
		return "", err.Error()
	}

	return c.Stdout.String(), c.Stderr.String()
}
