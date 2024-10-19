package application

import (
	"path/filepath"
	"sort"

	"treecli/internal/domain"
)

type TreeService struct {
	Repo         domain.TreeRepository
	ExcludeGlobs []string
	MaxDepth     int
}

func NewTreeService(repo domain.TreeRepository, excludeGlobs []string, maxDepth int) *TreeService {
	return &TreeService{
		Repo:         repo,
		ExcludeGlobs: excludeGlobs,
		MaxDepth:     maxDepth,
	}
}

func (ts *TreeService) BuildTree(path string, currentDepth int) (*domain.TreeNode, error) {
	if ts.MaxDepth > 0 && currentDepth > ts.MaxDepth {
		return &domain.TreeNode{
			Name:            filepath.Base(path),
			IsDepthExceeded: true,
		}, nil
	}

	isDir, err := ts.Repo.IsDir(path)
	if err != nil {
		return nil, err
	}

	node := &domain.TreeNode{
		Name: path,
	}

	shouldExclude, err := ts.shouldExclude(path)
	if err != nil {
		return nil, err
	}
	if shouldExclude {
		node.IsExcluded = true
		return node, nil
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
		fullPath := filepath.Join(path, entry)
		childNode, err := ts.BuildTree(fullPath, currentDepth+1)
		if err != nil {
			return nil, err
		}
		childNode.Name = entry
		childNode.IsLast = i == len(entries)-1
		if !childNode.IsExcluded {
			node.Children = append(node.Children, childNode)
		}
	}

	return node, nil
}

func (ts *TreeService) shouldExclude(path string) (bool, error) {
	base := filepath.Base(path)
	for _, pattern := range ts.ExcludeGlobs {
		match, err := filepath.Match(pattern, base)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}
