package domain

type TreeNode struct {
	Name       string
	Children   []*TreeNode
	IsLast     bool
	IsExcluded bool
}

type TreeRepository interface {
	ReadDir(path string) ([]string, error)
	IsDir(path string) (bool, error)
}
