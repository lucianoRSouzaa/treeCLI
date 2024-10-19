package application

import (
	"sort"

	"treecli/internal/domain"
)

type TreeService struct {
	Repo domain.TreeRepository
}

func NewTreeService(repo domain.TreeRepository) *TreeService {
	return &TreeService{
		Repo: repo,
	}
}

func (ts *TreeService) BuildTree(path string) (*domain.TreeNode, error) {
	isDir, err := ts.Repo.IsDir(path)
	if err != nil {
		return nil, err
	}

	node := &domain.TreeNode{
		Name: path,
	}

	if !isDir {
		return node, nil
	}

	entries, err := ts.Repo.ReadDir(path)
	if err != nil {
		return nil, err
	}

	sort.Strings(entries)

	for i, entry := range entries {
		fullPath := path + "/" + entry
		childNode, err := ts.BuildTree(fullPath)
		if err != nil {
			return nil, err
		}
		childNode.Name = entry
		childNode.IsLast = i == len(entries)-1
		node.Children = append(node.Children, childNode)
	}

	return node, nil
}
