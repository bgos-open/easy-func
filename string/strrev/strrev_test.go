package strrev

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		//FROM EXAMPLE:https://github.com/php/php-src/blob/master/ext/standard/tests/strings/strrev_basic.phpt
		{
			"A",
			args{
				"hello\\\\world",
			},
			"dlrow\\\\olleh",
		},
		{
			"B",
			args{
				"hello\\$world",
			},
			"dlrow$\\olleh",
		},
		{
			"C",
			args{
				"\\ttesting\\ttesting\\tstrrev",
			},
			"verrtst\\gnitsett\\gnitsett\\",
		},
		{
			"D",
			args{
				"ababababababa",
			},
			"ababababababa",
		},
		{
			"E",
			args{
				"a",
			},
			"a",
		},
		{
			"F",
			args{
				" ",
			},
			" ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.str); got != tt.want {
				t.Errorf("Do(%v) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
