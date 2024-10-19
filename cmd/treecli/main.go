package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"treecli/internal/application"
	"treecli/internal/infrastructure"
	"treecli/internal/interfaces"
)

func main() {
	exclude := flag.String("exclude", "", "Padrões de exclusão (wildcards), separados por vírgula")
	flag.Parse()

	path := "."
	args := flag.Args()
	if len(args) > 0 {
		path = args[0]
	}

	var excludeGlobs []string
	if *exclude != "" {
		excludeGlobs = strings.Split(*exclude, ",")
		for i, pattern := range excludeGlobs {
			excludeGlobs[i] = strings.TrimSpace(pattern)
		}
	}

	fsRepo := infrastructure.NewFileSystemRepository()
	treeService := application.NewTreeService(fsRepo, excludeGlobs)
	cli := interfaces.NewCLI(treeService)

	err := cli.Run(path)
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}
