package document

import "testing"

func Test_multiline(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "multiline string should be displayed as preformatted text",
			args: args{
				s: "ingress:\n  enabled: true\n  annotations:\n    kubernetes.io/ingress.class: traefik-external\n    traefik.frontend.rule.type: PathPrefixStrip\n  hosts:\n  - host: stream.quortex.io\n    paths:\n    - /",
			},
			want: "<pre>ingress:<br />  enabled: true<br />  annotations:<br />    kubernetes.io/ingress.class: traefik-external<br />    traefik.frontend.rule.type: PathPrefixStrip<br />  hosts:<br />  - host: stream.quortex.io<br />    paths:<br />    - /</pre>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiline(tt.args.s); got != tt.want {
				t.Errorf("multiline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_escaped(t *testing.T) {
	type args struct {
		runes string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "escaping should work corrrectly",
			args: args{
				runes: "|*",
				s:     "foo | bar*",
			},
			want: "foo \\| bar\\*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escaped(tt.args.runes, tt.args.s); got != tt.want {
				t.Errorf("escaped() = %v, want %v", got, tt.want)
			}
		})
	}
}
