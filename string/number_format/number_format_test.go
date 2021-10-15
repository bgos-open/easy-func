package number_format

import "testing"

func Test_Do(t *testing.T) {
	type args struct {
		number       float64
		decimals     uint
		decPoint     string
		thousandsSep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case 1:test string number format",
			args: args{
				number: 1234.5678,
			},
			want: "1,235",
		},
		{
			name: "case 2:test string number format",
			args: args{
				number: -1234.5678,
			},
			want: "-1,235",
		},
		{
			name: "case 3:test string number format",
			args: args{
				number: 1234.6578e4,
			},
			want: "12,346,578",
		},
		{
			name: "case 4:test string number format",
			args: args{
				number: -1234.56789e4,
			},
			want: "-12,345,679",
		},
		{
			name: "case 5:test string number format",
			args: args{
				number: 1234.5678,
				decimals: 2,
			},
			want: "1,234.57",
		},
		{
			name: "case 6:test string number format",
			args: args{
				number: -1234.5678,
				decimals: 2,
			},
			want: "-1,234.57",
		},
		{
			name: "case 7:test string number format",
			args: args{
				number: 1234.6578e4,
				decimals: 2,
			},
			want: "12,346,578.00",
		},
		{
			name: "case 8:test string number format",
			args: args{
				number: -1234.56789e4,
				decimals: 2,
			},
			want: "-12,345,678.90",
		},
		{
			name: "case 9:test string number format",
			args: args{
				number: 1234.5678,
				decimals: 2,
				decPoint: ".",
				thousandsSep: " ",
			},
			want: "1 234.57",
		},
		{
			name: "case 10:test string number format",
			args: args{
				number: -1234.5678,
				decimals: 2,
				decPoint: ".",
				thousandsSep: " ",
			},
			want: "-1 234.57",
		},
		{
			name: "case 11:test string number format",
			args: args{
				number: 1234.6578e4,
				decimals: 2,
				decPoint: ".",
				thousandsSep: " ",
			},
			want: "12 346 578.00",
		},
		{
			name: "case 12:test string number format",
			args: args{
				number: -1234.56789e4,
				decimals: 2,
				decPoint: ".",
				thousandsSep: " ",
			},
			want: "-12 345 678.90",
		},
		{
			name: "case 13:test string number format",
			args: args{
				number: 1234.5678,
				decimals: 2,
				decPoint: ",",
				thousandsSep: " ",
			},
			want: "1 234,57",
		},
		{
			name: "case 14:test string number format",
			args: args{
				number: -1234.5678,
				decimals: 2,
				decPoint: ",",
				thousandsSep: " ",
			},
			want: "-1 234,57",
		},
		{
			name: "case 15:test string number format",
			args: args{
				number: 1234.6578e4,
				decimals: 2,
				decPoint: ",",
				thousandsSep: " ",
			},
			want: "12 346 578,00",
		},
		{
			name: "case 16:test string number format",
			args: args{
				number: -1234.56789e4,
				decimals: 2,
				decPoint: ",",
				thousandsSep: " ",
			},
			want: "-12 345 678,90",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.number, tt.args.decimals, tt.args.decPoint, tt.args.thousandsSep); got != tt.want {
				t.Errorf("do() = %v, want %v", got, tt.want)
			}
		})
	}
}
