package kapitan

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// commentMark defines characters indicating comments to include in
// documentation.
const commentMark = "# --"

// extractComment extract comment from a given string. This function extracts
// the comment mark at the beginning of the comment, as well as the character #
// at the beginning of each line.
func extractComment(s string) string {
	res := []string{}
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		if len(res) > 0 {
			res = append(res, strings.TrimLeft(l, "# "))
		} else if strings.HasPrefix(l, commentMark) {
			res = append(res, strings.TrimLeft(l, commentMark+" "))
		}
	}
	return strings.Join(res, "\n")
}

// removeComments remove comments from a yaml multiline string.
func removeComments(s string) string {
	res := []string{}
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		if !strings.HasPrefix(strings.TrimLeft(l, " "), "#") {
			res = append(res, l)
		}
	}
	return strings.Join(res, "\n")
}

// PairedNodes wraps a pair of nodes (used to get key / value in the context of
// a MappingNode).
type PairedNodes struct {
	a *yaml.Node
	b *yaml.Node
}

// groupNodes groups given slice of nodes as paired nodes.
func groupNodes(nodes []*yaml.Node) []PairedNodes {
	res := make([]PairedNodes, 0, len(nodes)/2)
	for i := 1; i < len(nodes); i += 2 {
		res = append(res, PairedNodes{nodes[i-1], nodes[i]})
	}
	return res
}

// nodesAsStringSlice returns given nodes values as a slice of strings.
func nodesAsStringSlice(nodes []*yaml.Node) []string {
	res := make([]string, len(nodes))
	for i, n := range nodes {
		res[i] = n.Value
	}
	return res
}
