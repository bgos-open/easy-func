package strcasecmp

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		string1 string
		string2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				string1: "Hello, world",
				string2: "world",
			},
			want: 1,
		},
		{
			name: "case-2",
			args: args{
				string1: "Hello, world'S",
				string2: "world",
			},
			want: 1,
		},
		{
			name: "case-3",
			args: args{
				string1: "Hello, worldS",
				string2: "world",
			},
			want: 1,
		},
		{
			name: "case-4",
			args: args{
				string1: "Hello\\0world",
				string2: "Hello",
			},
			want: 1,
		},
		{
			name: "case-5",
			args: args{
				string1: "Hello\\0world",
				string2: "Helloworld",
			},
			want: 1,
		},
		{
			name: "case-6",
			args: args{
				string1: "\\x0",
				string2: "\\0",
			},
			want: 1,
		},
		{
			name: "case-7",
			args: args{
				string1: "\\000",
				string2: "\\0",
			},
			want: 1,
		},
		{
			name: "case-8",
			args: args{
				string1: "\\x00",
				string2: "",
			},
			want: 1,
		},
		{
			name: "case-9",
			args: args{
				string1: "10.5",
				string2: "10.5",
			},
			want: 0,
		},
		{
			name: "case-10",
			args: args{
				string1: "10.5",
				string2: "1.5",
			},
			want: 1,
		},
		{
			name: "case-11",
			args: args{
				string1: ";acc",
				string2: "acc",
			},
			want: 1,
		},
		{
			name: "case-12",
			args: args{
				string1: ";acc",
				string2: ";acc",
			},
			want: 0,
		},
		{
			name: "case-13",
			args: args{
				string1: "aC",
				string2: "aC",
			},
			want: 0,
		},
		{
			name: "case-14",
			args: args{
				string1: "aC",
				string2: "ac",
			},
			want: 0,
		},
		{
			name: "case-15",
			args: args{
				string1: "acCc",
				string2: "acc",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.string1, tt.args.string2); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
