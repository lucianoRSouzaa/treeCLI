package application

import (
	"path/filepath"
	"sort"
	"strings"

	"treecli/internal/domain"
)

type TreeService struct {
	Repo         domain.TreeRepository
	ExcludeGlobs []string
	MaxDepth     int
	IncludeExts  []string
	ExcludeExts  []string
}

func NewTreeService(repo domain.TreeRepository, excludeGlobs []string, maxDepth int, includeExts, excludeExts []string) *TreeService {
	return &TreeService{
		Repo:         repo,
		ExcludeGlobs: excludeGlobs,
		MaxDepth:     maxDepth,
		IncludeExts:  includeExts,
		ExcludeExts:  excludeExts,
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
		Name: filepath.Base(path),
	}

	shouldExclude, err := ts.shouldExclude(path, isDir)
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
		childIsDir, err := ts.Repo.IsDir(fullPath)
		if err != nil {
			return nil, err
		}

		shouldExclude, err := ts.shouldExclude(fullPath, childIsDir)
		if err != nil {
			return nil, err
		}
		if shouldExclude {
			continue
		}

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

func (ts *TreeService) shouldExclude(path string, isDir bool) (bool, error) {
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

	if !isDir {
		ext := strings.ToLower(filepath.Ext(base))

		for _, excludeExt := range ts.ExcludeExts {
			if strings.EqualFold(excludeExt, ext) {
				return true, nil
			}
		}

		if len(ts.IncludeExts) > 0 {
			included := false
			for _, includeExt := range ts.IncludeExts {
				if strings.EqualFold(includeExt, ext) {
					included = true
					break
				}
			}
			if !included {
				return true, nil
			}
		}
	}

	return false, nil
}
