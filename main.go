package main

import (
	"fmt"
	"os"
)

func main() {
	envs := os.Environ()
	for _, e := range envs {
		fmt.Println(e)
	}
}
