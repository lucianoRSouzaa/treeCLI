package interfaces

import (
	"fmt"

	"treecli/internal/application"
	"treecli/internal/domain"
)

type CLI struct {
	Service *application.TreeService
}

func NewCLI(service *application.TreeService) *CLI {
	return &CLI{
		Service: service,
	}
}

func (cli *CLI) Run(path string) error {
	tree, err := cli.Service.BuildTree(path)
	if err != nil {
		return err
	}

	cli.printTree(tree, "")
	return nil
}

func (cli *CLI) printTree(node *domain.TreeNode, prefix string) {
	if node.IsExcluded && prefix != "" {
		fmt.Println(prefix + node.Name + " (conteúdo oculto)")
		return
	}

	fmt.Println(prefix + node.Name)
	cli.printChildren(node.Children, prefix)
}

func (cli *CLI) printChildren(children []*domain.TreeNode, prefix string) {
	for _, child := range children {
		connector := "├── "
		newPrefix := prefix + "│   "
		if child.IsLast {
			connector = "└── "
			newPrefix = prefix + "    "
		}

		if child.IsExcluded {
			fmt.Println(prefix + connector + child.Name + " (conteúdo oculto)")
		} else {
			fmt.Println(prefix + connector + child.Name)
			if len(child.Children) > 0 {
				cli.printChildren(child.Children, newPrefix)
			}
		}
	}
}
