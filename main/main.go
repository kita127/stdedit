package main

import (
	"fmt"
	"github.com/kita127/stdedit"
	"io/ioutil"
	"os"
)

func main() {
	body, _ := ioutil.ReadAll(os.Stdin)
	edited := stdedit.Edit(body)
	fmt.Print(string(edited))
}
