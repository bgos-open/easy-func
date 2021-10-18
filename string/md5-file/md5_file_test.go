package md5

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/md5_file.phpt
		{
			"case-1:No filename",
			args{``},
			"文件路径不能为空",
		},
		{
			"case-2:Invalid filename",
			args{`aZrq16u`},
			"打开文件失败，filename=aZrq16u, err=open ./aZrq16u: The system cannot find the file specified.",
		},
		{
			"case-3:Scalar value as filename",
			args{`12`},
			"打开文件失败，filename=12, err=open ./12: The system cannot find the file specified.",
		},
		{
			"case-4:Normal filename",
			args{`./file-md5.txt`},
			"0b42f1bec4ecbf96877136cf408468b8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.filename); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
