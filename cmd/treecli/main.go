package main

import (
	"fmt"
	"os"

	"treecli/internal/application"
	"treecli/internal/infrastructure"
	"treecli/internal/interfaces"
)

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fsRepo := infrastructure.NewFileSystemRepository()
	treeService := application.NewTreeService(fsRepo)
	cli := interfaces.NewCLI(treeService)

	err := cli.Run(path)
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}
