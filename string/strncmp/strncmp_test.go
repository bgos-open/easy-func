package strncmp

import (
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		str1   string
		str2   string
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
		wantErr bool
	}{
		//FROM Example:https://github.com/php/php-src/blob/master/ext/standard/tests/strings/strncmp_basic.phpt
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				str1:   "str1",
				str2:   "str2",
				length: -1,
			},
			wantRes: 0,
			wantErr: true,
		},
		{
			name: "case-2",
			args: args{
				str1:   "str1",
				str2:   "str2",
				length: 0,
			},
			wantRes: 0,
			wantErr: false,
		},
		{
			name: "case-3",
			args: args{
				str1:   "str1",
				str2:   "str2",
				length: 5,
			},
			wantRes: -1,
			wantErr: false,
		},
		{
			name: "case-4",
			args: args{
				str1:   "str1",
				str2:   "str2",
				length: 3,
			},
			wantRes: 0,
			wantErr: false,
		},
		{
			name: "case-5",
			args: args{
				str1:   "str1",
				str2:   "str2",
				length: 4,
			},
			wantRes: -1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Do(tt.args.str1, tt.args.str2, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("Do() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
