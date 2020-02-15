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

	var size int
	if size, err = tmpfile.Write(input); err != nil {
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
		log.Fatal("abort renames : ", err)
	}

	readFile, err := os.Open(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	buf := make([]byte, 1024)
	for {
		n, err := readFile.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			log.Fatal(size, err)
		}
		fmt.Print(string(buf))
	}
}
