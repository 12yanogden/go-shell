package shell

import (
	"os"
	"os/exec"

	"github.com/12yanogden/errors"
	"golang.org/x/term"
)

func Run(command string, args []string) string {
	out, err := exec.Command(command, args...).Output()

	if err != nil {
		errors.Scream(err.Error())
	}

	return string(out)
}

func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}
