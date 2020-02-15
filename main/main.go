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

	tmpfile, err := ioutil.TempFile("", "stdedit-temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err = tmpfile.Write(input); err != nil {
		log.Fatal(err)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	tty, err := tty.Open()
	if err != nil {
		log.Fatal("tty error : ", err)
	}
	defer tty.Close()
	cmd := exec.Command("vi", tmpfile.Name())
	cmd.Stdin = tty.Input()
	cmd.Stdout = tty.Output()
	cmd.Stderr = tty.Output()
	if err := cmd.Run(); err != nil {
		log.Fatal("abort edit : ", err)
	}

	output, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(os.Stdout, string(output))
}
