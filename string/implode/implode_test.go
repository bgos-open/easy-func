package implode

import (
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		sep   string
		array []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string

	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{":", []string{"foo", "bar","baz"}},
			wantRes: "foo:bar:baz",

		},
		{
			name: "case-2",
			args: args{":", []string{"a"}},
			wantRes: "a",

		},
		{
			name: "case-3",
			args: args{"", []string{"foo", "bar","baz"}},
			wantRes: "foobarbaz",

		},
		{
			name: "case-4",
			args: args{":", []string{""}},
			wantRes: "",

		},
		{
			name: "case-5",
			args: args{":", []string{}},
			wantRes: "",

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := Do(tt.args.sep, tt.args.array)
			if gotRes != tt.wantRes {
				t.Errorf("Do() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
