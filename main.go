package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	f, err := ioutil.TempFile("", "stdedit-temp")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(f.Name()) // clean up
	}()

	if _, err = f.Write(input); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	editor := os.Getenv("STDEDIT")
	if editor == "" {
		editor = "vim"
	}
	cmd := exec.Command(editor, f.Name())
	cmd.Stdin = tty.Input()
	cmd.Stdout = tty.Output()
	cmd.Stderr = tty.Output()
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	output, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(os.Stdout, string(output))
}
