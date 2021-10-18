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
		want interface{}
	}{
		{
			name: "case 1:length=0",
			args: args{
				"hello world",
				"w",
				2,
				0,
			},
			want: 0,
		},
		{
			name: "case 2:offset>0 && length<0",
			args: args{
				"hello world",
				"w",
				2,
				-2,
			},
			want: 4,
		},
		{
			name: "case 3:offset>0 && length>0",
			args: args{
				"hello world",
				"w",
				2,
				2,
			},
			want: 2,
		},
		{
			name: "case 4:offset<0 && length<0",
			args: args{
				"hello world",
				"w",
				-8,
				-2,
			},
			want: 3,
		},
		{
			name: "case 5:offset<0 && length<0",
			args: args{
				"hello world",
				"w",
				-8,
				-8,
			},
			want: 0,
		},
		{
			name: "case 6:offset>0 && length+∞",
			args: args{
				"hello world",
				"w",
				2,
				2147483647,
			},
			want: 4,
		},
		{
			name: "case 7:offset=0 && length+∞",
			args: args{
				"hello world",
				"w",
				0,
				2147483647,
			},
			want: 6,
		},
		{
			name: "case 8:offset<0 && length+∞",
			args: args{
				"hello world",
				"w",
				-8,
				2147483647,
			},
			want: 3,
		},
		{
			name: "case 9:two or more char",
			args: args{
				"hello world",
				"ok",
				0,
				2147483647,
			},
			want: 4,
		},
		{
			name: "case 10:offset > len(str)",
			args: args{
				"hello world",
				"ok",
				20,
				2147483647,
			},
			want: false,
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