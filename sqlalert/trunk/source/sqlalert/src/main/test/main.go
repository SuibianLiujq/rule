package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./func.tmp")
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		if line == "" {
			continue
		}

		key := strings.Fields(line)[0]
		fmt.Println(key)
	}
}
