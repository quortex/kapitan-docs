package kapitan

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type WalkFunc func(path []string, node yaml.Node) error

// parsed wraps data about a yaml file
type parsed struct {
	path string
	info *os.FileInfo
	node *yaml.Node
}

// parse deep parse the given directory to return data about yaml files it
// contains.
func parse(dir string) ([]*parsed, error) {
	res := []*parsed{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// we only return yaml files
		if isYaml(info) {
			f, err := ioutil.ReadFile(path)
			if err != nil {
				log.Error(fmt.Errorf("Read file error: %w", err))
				return err
			}

			// Yaml file content unmarshal
			var node yaml.Node
			err = yaml.Unmarshal(f, &node)
			if err != nil {
				log.Errorf("Cannot unmarshal yaml: %v", err)
				return err
			}

			// Get relative path from directory
			p, err := filepath.Rel(dir, path)
			if err != nil {
				log.Errorf("Cannot get relative path: %v", err)
				return err
			}
			res = append(res, &parsed{path: p, info: &info, node: &node})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Walk walks a yaml node, calling fn for each node in the tree, including root.
func Walk(root yaml.Node, fn WalkFunc) error {
	for _, node := range root.Content {
		if err := fn([]string{}, *node); err != nil {
			return err
		}
		if err := walk(*node, []string{}, fn); err != nil {
			return err
		}
	}
	return nil
}

// walk walks the yaml node, calling fn for each node in the tree, including
// root. It also computes the nodes path in the file.
func walk(root yaml.Node, path []string, fn WalkFunc) error {
	for i := range root.Content {
		node := root.Content[i]
		if node.Kind == yaml.MappingNode && i > 0 {
			path = append(path, root.Content[i-1].Value)
		}
		if err := fn(path, *node); err != nil {
			return err
		}
		if err := walk(*node, path, fn); err != nil {
			return err
		}
	}
	return nil
}

// isYaml returns if given file is a yaml file.
func isYaml(info os.FileInfo) bool {
	return !info.IsDir() && slices.Contains([]string{".yaml", ".yml"}, filepath.Ext(info.Name()))
}
