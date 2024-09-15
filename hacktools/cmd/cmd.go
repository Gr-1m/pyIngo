package cmd

import (
	"bytes"
	"os/exec"
	"strings"
)

func Run(cmd string) (Stdout, Stderr string) {

	if strings.Contains(cmd, "|") {
		c1c2str := strings.Split(cmd, "|")
		c1 := strings.Split(c1c2str[0], " ")
		c2 := strings.Split(c1c2str[1], " ")

		var stdout, stderr bytes.Buffer
		r1 := exec.Command(c1[0], c1...)
		r2 := exec.Command(c2[0], c2...)
		r2.Stdin, _ = r1.StdoutPipe()
		r2.Stdout = &stdout
		r2.Stderr = &stderr
		// err := c.Run()
		_ = r2.Start()
		_ = r1.Run()
		_ = r2.Wait()

		return stdout.String(), stderr.String()
	}

	c := strings.Split(cmd, " ")
	runner := exec.Command(c[0], c[1:]...)

	var stdout, stderr bytes.Buffer
	runner.Stdout = &stdout
	runner.Stderr = &stderr
	err := runner.Run()
	if err != nil {
		return "", err.Error()
	}

	return stdout.String(), stderr.String()
}
