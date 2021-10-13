package stristr

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		str string
		sep string
		before_search bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case-1",
			args: args{"tEsT sTrInG", "tEsT", false},
			want: "tEsT sTrInG",
			wantErr: false,
		},
		{
			name: "case-2",
			args: args{"tEsT sTrInG", "stRiNg", false},
			want: "sTrInG",
			wantErr: false,
		},
		{
			name: "case-3",
			args: args{"tEsT sTrInG", "stRiN", false},
			want: "sTrInG",
			wantErr: false,
		},
		{
			name: "case-4",
			args: args{"tEsT sTrInG", "t S", false},
			want: "T sTrInG",
			wantErr: false,
		},
		{
			name: "case-5",
			args: args{"tEsT sTrInG", "g", false},
			want: "G",
			wantErr: false,
		},
		{
			name: "case-6",
			args: args{"", "", false},
			want: "",
			wantErr: true,
		},
		{
			name: "case-7",
			args: args{"a", "", false},
			want: "",
			wantErr: true,
		},
		{
			name: "case-8",
			args: args{"", "a", false},
			want: "",
			wantErr: true,
		},
		{
			name: "case-9",
			args: args{"\\\\a\\", "\\a", false},
			want: "\\a\\",
			wantErr: false,
		},
		{
			name: "case-10",
			args: args{"tEsT sTrInG", " ", false},
			want: " sTrInG",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Do(tt.args.str, tt.args.sep, tt.args.before_search)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}
