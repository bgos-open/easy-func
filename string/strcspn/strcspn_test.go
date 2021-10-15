package strcspn

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		haystack string
		char     string
		offset   int
		length   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1:offset > len(str)",
			args: args{
				"hello world",
				"wor",
				100,
				0,
			},
			want: 0,
		},
		{
			name: "case 2:offset>0",
			args: args{
				"hello world",
				"w",
				2,
				0,
			},
			want: 4,
		},
		{
			name: "case 3:offset < 0",
			args: args{
				"hello world",
				"w",
				-2,
				0,
			},
			want: 2,
		},
		{
			name: "case 4:offset == 0",
			args: args{
				"hello world",
				"w",
				0,
				0,
			},
			want: 6,
		},
		{
			name: "case 5:length>0",
			args: args{
				"hello world",
				"w",
				0,
				2,
			},
			want: 2,
		},
		{
			name: "case 6:length>len(str)",
			args: args{
				"hello world",
				"w",
				0,
				2147483647,
			},
			want: 6,
		},
		{
			name: "case 7:two or more char",
			args: args{
				"hello world",
				"world",
				0,
				2147483647,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.haystack, tt.args.char, tt.args.offset, tt.args.length); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}