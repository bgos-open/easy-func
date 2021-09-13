package rtrim

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		str       string
		character []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case-1:blank",
			args: args{
				str:       " 12 3 ",
				character: []string{},
			},
			want: " 12 3",
		},
		{
			name: "case-2.1:newline",
			args: args{
				str:       "\n\n\t\v\u00000\v\v\v 12 3\n\n\t\v\u00000\v\v\v ",
				character: []string{},
			},
			want: "\n\n\t\v\u00000\v\v\v 12 3",
		},
		{
			name: "case-2.2:multi",
			args: args{
				str:       "\n\v 12 3 \n\v",
				character: []string{},
			},
			want: "\n\v 12 3",
		},
		{
			name: "case-3.1:different char",
			args: args{
				str:       `abAc 12 3 abAc`,
				character: []string{"bacAcab"},
			},
			want: "abAc 12 3 ",
		},
		{
			name: "case-4.1:none",
			args: args{
				str:       ``,
				character: []string{"bacAcab"},
			},
			want: "",
		},
		{
			name: "case-4.2:none",
			args: args{
				str:       ` `,
				character: []string{""},
			},
			want: " ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.str, tt.args.character...); got != tt.want {
				t.Errorf("Trim() = |%s|, want |%s|", got, tt.want)
			}
		})
	}

}

func TestDo1(t *testing.T) {
	type args struct {
		s       string
		cutSets []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.s, tt.args.cutSets...); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
