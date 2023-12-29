package shell

import (
	"os"
	"os/exec"

	"golang.org/x/term"
)

func Run(command string, args []string) string {
	out, err := exec.Command(command, args...).Output()

	if err != nil {
		panic(err)
	}

	return string(out)
}

func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

