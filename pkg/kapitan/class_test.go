package kapitan

import (
	"reflect"
	"testing"
)

func Test_className(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty path should return empty name",
			args: args{
				path: "",
			},
			want: "",
		},
		{
			name: "file path should return correctly formatted class name",
			args: args{
				path: "infra/cloud_provider/aws.yaml",
			},
			want: "infra.cloud_provider.aws",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := className(tt.args.path); got != tt.want {
				t.Errorf("className() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillUsedBy(t *testing.T) {
	type args struct {
		classes []Class
	}
	tests := []struct {
		name string
		args args
		want []Class
	}{
		{
			name: "Class UsedBy should be correctly filled",
			args: args{
				classes: []Class{
					{
						Path:        "foo/bar.yaml",
						Name:        "foo.bar",
						Description: "a foo/bar class",
						Uses:        []string{"bar.foo"},
					},
					{
						Path:        "bar/foo.yaml",
						Name:        "bar.foo",
						Description: "a bar/foo class",
					},
				},
			},
			want: []Class{
				{
					Path:        "foo/bar.yaml",
					Name:        "foo.bar",
					Description: "a foo/bar class",
					Uses:        []string{"bar.foo"},
				},
				{
					Path:        "bar/foo.yaml",
					Name:        "bar.foo",
					Description: "a bar/foo class",
					UsedBy:      []string{"foo.bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillUsedBy(tt.args.classes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillUsedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
