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
	maxDepth := flag.Int("max-depth", 0, "Limite de profundidade da árvore (0 para ilimitado)")
	includeExts := flag.String("ext", "", "Lista de extensões de arquivo a serem incluídas (por exemplo, .go,.md)")
	excludeExts := flag.String("exclude-ext", "", "Lista de extensões de arquivo a serem excluídas (por exemplo, .txt,.log)")
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

	var includeExtList []string
	if *includeExts != "" {
		includeExtList = strings.Split(*includeExts, ",")
		for i, ext := range includeExtList {
			includeExtList[i] = strings.TrimSpace(strings.ToLower(ext))
			if !strings.HasPrefix(includeExtList[i], ".") {
				includeExtList[i] = "." + includeExtList[i]
			}
		}
	}

	var excludeExtList []string
	if *excludeExts != "" {
		excludeExtList = strings.Split(*excludeExts, ",")
		for i, ext := range excludeExtList {
			excludeExtList[i] = strings.TrimSpace(strings.ToLower(ext))
			if !strings.HasPrefix(excludeExtList[i], ".") {
				excludeExtList[i] = "." + excludeExtList[i]
			}
		}
	}

	fsRepo := infrastructure.NewFileSystemRepository()
	treeService := application.NewTreeService(fsRepo, excludeGlobs, *maxDepth, includeExtList, excludeExtList)
	cli := interfaces.NewCLI(treeService)

	err := cli.Run(path)
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}
