package shell

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

func Run(command string, args []string) string {
	out, err := exec.Command(command, args...).Output()

	if err != nil {
		fmt.Printf("%s: %s", command, err.Error())
		os.Exit(1)
	}

	return string(out)
}

func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}
