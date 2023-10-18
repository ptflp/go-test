package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, prefix string, isLast bool, depth int) {

}

func main() {
	// получение флага
	var depth int
	//flag.IntVar
	//flag.Parse()
	path := flag.Arg(0)
	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}
	printTree(path, "", true, depth)
}
