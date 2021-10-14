package join

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		glue   string
		pieces []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"case-1:blank",
			args:args {
				glue: ",",
				pieces: []string{"1", "2", "3", "4"},
			},
			want:"1,2,3,4",
		},
		{
			name:"case-2:blank",
			args:args {
				glue: ",",
				pieces: []string{"Red", "Green", "Blue", "Black", "White"},
			},
			want:"Red,Green,Blue,Black,White",
		},
		{
			name:"case-3:blank",
			args:args {
				glue: "0",
				pieces: []string{"element1", "element2"},
			},
			want:"element10element2",
		},
		{
			name:"case-4:blank",
			args:args {
				glue: "1",
				pieces: []string{"element1", "element2"},
			},
			want:"element11element2",
		},
		{
			name:"case-5:blank",
			args:args {
				glue: "12345",
				pieces: []string{"element1", "element2"},
			},
			want:"element112345element2",
		},
		{
			name:"case-6:blank",
			args:args {
				glue: "-2345",
				pieces: []string{"element1", "element2"},
			},
			want:"element1-2345element2",
		},
		{
			name:"case-7:blank",
			args:args {
				glue: "10.5",
				pieces: []string{"element1", "element2"},
			},
			want:"element110.5element2",
		},
		{
			name:"case-8:blank",
			args:args {
				glue: "-10.5",
				pieces: []string{"element1", "element2"},
			},
			want:"element1-10.5element2",
		},
		{
			name:"case-9:blank",
			args:args {
				glue: "10.1234567e10",
				pieces: []string{"element1", "element2"},
			},
			want:"element110.1234567e10element2",
		},
		{
			name:"case-10:blank",
			args:args {
				glue: "10.7654321E-10",
				pieces: []string{"element1", "element2"},
			},
			want:"element110.7654321E-10element2",
		},{
			name:"case-11:blank",
			args:args {
				glue: ".5",
				pieces: []string{"element1", "element2"},
			},
			want:"element1.5element2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.glue, tt.args.pieces); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
