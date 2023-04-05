package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic a",
			args: args{3, 4, 5},
			want: "YES",
		},
		{
			name: "basic b",
			args: args{4, 3, 5},
			want: "YES",
		},
		{
			name: "basic c",
			args: args{5, 4, 3},
			want: "YES",
		},
		{
			name: "equal",
			args: args{5, 5, 5},
			want: "YES",
		},
		{
			name: "wrong a",
			args: args{5, 1, 2},
			want: "NO",
		},
		{
			name: "wrong b",
			args: args{2, 5, 1},
			want: "NO",
		},
		{
			name: "wrong c",
			args: args{1, 2, 5},
			want: "NO",
		},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")
			fmt.Fprintf(f, "%d\n%d\n%d", tt.args.a, tt.args.b, tt.args.c)

			main()

			f,_ = os.Open("output.txt")
			
			
			var got string
			fmt.Fscan(f, &got)

			if got != tt.want{
				t.Errorf("wrong response: %v, want: %v", got, tt.want)
			}
		})
	} 
}
