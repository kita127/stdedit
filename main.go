package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	body, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print(string(body))
}
