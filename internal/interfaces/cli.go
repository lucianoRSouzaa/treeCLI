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
	tree, err := cli.Service.BuildTree(path, 0)
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

		line := prefix + connector + child.Name

		if child.IsExcluded {
			line += " (conteúdo oculto)"
			fmt.Println(line)
		} else if child.IsDepthExceeded {
			line += " (profundidade máxima atingida)"
			fmt.Println(line)
		} else {
			fmt.Println(line)
			if len(child.Children) > 0 {
				cli.printChildren(child.Children, newPrefix)
			}
		}
	}
}
