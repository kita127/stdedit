package main

import (
	"fmt"
	_ "github.com/kita127/stdedit"
	"io"
	"io/ioutil"
	"log"
	"os"
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

	readFile, err := os.Open(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}

	output := make([]byte, size)
	if size, err = readFile.Read(output); err == io.EOF {
		log.Fatal(size, err)
	}

	if err := readFile.Close(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("file path : %s\n", tmpfile.Name())
	// fmt.Printf("file input : %s\n", string(output))
	//edited := stdedit.Edit(body)
	fmt.Println(string(output))
}
