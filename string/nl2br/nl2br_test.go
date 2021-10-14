package nl2br

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		str     string
		isXhtml bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case-1:str does not contain new line chars, with XHTML",
			args: args{
				str : "test",
				isXhtml : true,
			},
			want: "test",
		},
		{
			name: "case-2:str does not contain new line chars,without XHTML",
			args: args{
				str : "test",
				isXhtml : false,
			},
			want: "test",
		},
		{
			name: "case-3:empty str,with XHTML",
			args: args{
				str : "",
				isXhtml : true,
			},
			want: "",
		},
		{
			name: "case-4:empty str, without XHTML",
			args: args{
				str : "",
				isXhtml : false,
			},
			want: "",
		},
		{
			name: "case-5:str contain new line chars(windows), with XHTML",
			args: args{
				str : "\r\n",
				isXhtml : true,
			},
			want: "<br />\r\n",
		},
		{
			name: "case-6:str contain new line chars(windows), without XHTML",
			args: args{
				str : "\r\n",
				isXhtml : false,
			},
			want: "<br>\r\n",
		},
		{
			name: "case-7:str contain new line chars(unix)，with XHTML",
			args: args{
				str : "\n",
				isXhtml : true,
			},
			want: "<br />\n",
		},
		{
			name: "case-8:str contain new line chars(unix),without XHTML",
			args: args{
				str : "\n",
				isXhtml : false,
			},
			want: "<br>\n",
		},
		{
			name: "case-9:str contain new line chars(unix & windows), with XHTML",
			args: args{
				str : "\n\r",
				isXhtml : true,
			},
			want: "<br />\n\r",
		},
		{
			name: "case-10:str contain new line chars(unix & windows),without XHTML",
			args: args{
				str : "\n\r",
				isXhtml : false,
			},
			want: "<br>\n\r",
		},
		{
			name: "case-11:str contain new line chars(unix & windows), with XHTML",
			args: args{
				str : "\n\r\r\n\r\r\r\r",
				isXhtml : true,
			},
			want: "<br />\n\r<br />\r\n<br />\r<br />\r<br />\r<br />\r",
		},
		{
			name: "case-12:str contain new line chars(unix & windows), without XHTML",
			args: args{
				str : "\n\r\r\n\r\r\r\r",
				isXhtml : false,
			},
			want: "<br>\n\r<br>\r\n<br>\r<br>\r<br>\r<br>\r",
		},
		{
			name: "case-13:str contain new line chars(unix & windows),with XHTML",
			args: args{
				str : "\n\r\n\n\r\n\r\r\n\r\n",
				isXhtml : true,
			},
			want: "<br />\n\r<br />\n<br />\n\r<br />\n\r<br />\r\n<br />\r\n",
		},
		{
			name: "case-14:str contain new line chars(unix & windows),without XHTML",
			args: args{
				str : "\n\r\n\n\r\n\r\r\n\r\n",
				isXhtml : false,
			},
			want: "<br>\n\r<br>\n<br>\n\r<br>\n\r<br>\r\n<br>\r\n",
		},
		{
			name: "case-15:str contain new line chars(unix & windows), with XHTML",
			args: args{
				str : "\n\r\n\n\n\n\r\r\r\r\n\r",
				isXhtml : true,
			},
			want: "<br />\n\r<br />\n<br />\n<br />\n<br />\n\r<br />\r<br />\r<br />\r\n<br />\r",
		},
		{
			name: "case-16:str contain new line chars(unix & windows), without XHTML",
			args: args{
				str : "\n\r\n\n\n\n\r\r\r\r\n\r",
				isXhtml : false,
			},
			want: "<br>\n\r<br>\n<br>\n<br>\n<br>\n\r<br>\r<br>\r<br>\r\n<br>\r",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.str, tt.args.isXhtml); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
