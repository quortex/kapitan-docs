package kapitan

import (
	"fmt"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

// Class defines a class in a kapitan project.
type Class struct {
	Path        string      // the class path in the project structure
	Name        string      // the class name
	Description string      // a description for the class (extracted from file header comments)
	Uses        []string    // a slice of classes from which this class inherits
	UsedBy      []string    // a slice of classes that inherit from this class
	Parameters  []Parameter // the parameters of this class
}

// IsEmpty returns if this class is suitable to be displayed in the
// documentation.
func (c Class) IsEmpty() bool {
	return c.Description == "" && len(c.Parameters) == 0
}

// Parameter defines a kapitan parameter used in a class.
type Parameter struct {
	Key          string        // the key in the yaml
	Kind         ParameterType // the parameter kind
	Description  string        // a description for the parameter (extracted from parameter header comment in the class)
	DefaultValue string        // the parameter default value (the value defined in the class)
}

// ParameterType defines the type of parameter.
type ParameterType string

const (
	ParameterTypeUnknown = "unknown"
	ParameterTypeString  = "string"
	ParameterTypeNumber  = "number"
	ParameterTypeObject  = "object"
	ParameterTypeList    = "list"
)

// parameterTypeFromTag instantiates a parameter type from yaml node tag
func parameterTypeFromTag(tag string) ParameterType {
	switch tag {
	case "!!str":
		return ParameterTypeString
	case "!!int", "!!float":
		return ParameterTypeNumber
	case "!!map":
		return ParameterTypeObject
	case "!!seq":
		return ParameterTypeList
	}

	return ParameterTypeUnknown
}

// NewClasses returns a slice of non empty classes from given parsed yaml files.
func NewClasses(parsed []*parsed) ([]Class, error) {
	res := []Class{}
	for _, p := range parsed {
		c, err := classFromParsed(*p)
		if err != nil {
			return nil, err
		}
		if !c.IsEmpty() {
			res = append(res, *c)
		}
	}
	res = fillUsedBy(res)

	return res, nil
}

// classFromParsed instantiate a Class from the given parsed yaml files.
func classFromParsed(parsed parsed) (*Class, error) {
	res := Class{Path: parsed.path, Name: className(parsed.path)}
	classesParsed := false

	// get class description from document node header comment
	if parsed.node.Kind == yaml.DocumentNode {
		res.Description = extractComment(parsed.node.HeadComment)
	}

	// walk over yaml nodes to instantiate Class correctly.
	err := Walk(*parsed.node, func(path []string, node yaml.Node) error {
		log.Info(fmt.Sprintf("Walking on node: %s", node.Value))

		// we consider only the Mapping nodes for the parsing of the attributes of
		// the class
		if node.Kind == yaml.MappingNode {
			for _, gn := range groupNodes(node.Content) {

				// parses classes from which this class inherits
				if gn.a.Value == "classes" && !classesParsed {
					classesParsed = true
					res.Uses = nodesAsStringSlice(gn.b.Content)
				}

				// fill class parameters from commented ones
				comment := extractComment(gn.a.HeadComment)
				if comment != "" {
					defaultValue, err := yaml.Marshal(&gn.b)
					if err != nil {
						log.Fatal(err)
					}
					res.Parameters = append(res.Parameters, Parameter{
						Key:          strings.Join(append(path, gn.a.Value), "."),
						Kind:         parameterTypeFromTag(gn.b.Tag),
						Description:  comment,
						DefaultValue: removeComments(string(defaultValue)),
					})
				}
			}
			return nil
		} else {
			return nil
		}
	})
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// fillUsedBy iterate over a slice of class to determine for each class, the
// classes that inherit from this one.
func fillUsedBy(classes []Class) []Class {
	res := make([]Class, len(classes))
	for i, c := range classes {
		for _, c2 := range classes {
			if slices.Contains(c2.Uses, c.Name) {
				c.UsedBy = append(c.UsedBy, c2.Name)
			}
		}
		res[i] = c
	}
	return res
}

// className computes the class name from it's file path
func className(path string) string {
	return strings.Join(strings.Split(strings.TrimSuffix(path, filepath.Ext(path)), "/"), ".")
}
