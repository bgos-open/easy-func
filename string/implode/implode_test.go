package implode

import (
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		sep   string
		array interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{",", []int{1,2,3}},
			wantRes: "1,2,3",
			wantErr: false,
		},
		{
			name: "case-2",
			args: args{"nothing", []int{}},
			wantRes: "",
			wantErr: false,
		},
		{
			name: "case-3",
			args: args{":", []string{"foo", "bar","baz"}},
			wantRes: "foo:bar:baz",
			wantErr: false,
		},
		{
			name: "case-4",
			args: args{":", map[int]interface{}{1:"foo", 2:[]string{"bar", "baz"}, 3:"burp"}},
			wantRes: "foo:slice:burp",
			wantErr: false,
		},
		{
			name: "case-5",
			args: args{":", map[int]interface{}{1:"foo", 2:"bar", 3:1}},
			wantRes: "foo:bar:1",
			wantErr: false,
		},
		{
			name: "case-6",
			args: args{":", []int64{1,2,3,4}},
			wantRes: "1:2:3:4",
			wantErr: false,
		},
		{
			name: "case-7",
			args: args{":", "3434"},
			wantRes: "",
			wantErr: true,
		},
		{
			name: "case-8",
			args: args{":", []float32{1,2,3,4}},
			wantRes:  "1:2:3:4",
			wantErr: false,
		},
		{
			name: "case-9",
			args: args{":", []uint8{1,2,3,4}},
			wantRes:  "1:2:3:4",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Do(tt.args.sep, tt.args.array)
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
