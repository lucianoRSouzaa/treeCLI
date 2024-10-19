package domain

type TreeNode struct {
	Name     string
	Children []*TreeNode
	IsLast   bool
}

type TreeRepository interface {
	ReadDir(path string) ([]string, error)
	IsDir(path string) (bool, error)
}
