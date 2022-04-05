package kapitan

import "testing"

func Test_extractComment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string should return empty comment",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "string with no comment mark should return empty comment",
			args: args{
				s: "This is a dummy string\nWith no comment mark!",
			},
			want: "",
		},
		{
			name: "string with misplaced comment mark should return empty comment",
			args: args{
				s: "This is a dummy string\nWith a # -- misplaced comment mark!",
			},
			want: "",
		},
		{
			name: "string with comment mark should return comment starting at comment mark position",
			args: args{
				s: "This is a dummy string\n# -- With a comment mark!\n# And another line.\n# -- and a comment mark to ignore",
			},
			want: "With a comment mark!\nAnd another line.\n-- and a comment mark to ignore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractComment(tt.args.s); got != tt.want {
				t.Errorf("extractComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
