package kapitan

import "path/filepath"

const (
	// Classes directory in a typical kapitan project.
	dirClasses = "inventory/classes"
)

// Project describes a kapitan project.
type Project struct {
	Classes []Class
}

// NewProject returns a newly instanciated project initialized from project
// directory path.
func NewProject(dir string) (*Project, error) {
	parsedClasses, err := parse(filepath.Join(dir, dirClasses))
	if err != nil {
		return nil, err
	}
	classes, err := NewClasses(parsedClasses)
	if err != nil {
		return nil, err
	}
	return &Project{
		Classes: classes,
	}, nil
}
