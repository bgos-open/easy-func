package debugZvalDump

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// reference: https://github.com/php/php-src/blob/5b01c4863fe9e4bc2702b2bbf66d292d23001a18/ext/standard/tests/general_functions/debug_zval_dump_v.phpt
		{
			name: "case-1",
			args: args{
				values: []interface{}{},
			},
		},
		{
			name: "case-2",
			args: args{
				values: []interface{}{
					"1", "abc",
				},
			},
		},
		{
			name: "case-3",
			args: args{
				values: []interface{}{
					"1",
					"abc",
					struct {
					}{},
					struct {
						a string
					}{
						a: "123",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Do(tt.args.values...)
		})
	}
}
