package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T){
	type args [5]int
	tests := []struct{
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{1,1,1,1,1},
			want: "YES",
		},
		{
			name: "!ok",
			args: args{2,2,2,1,1},
			want: "NO",
		},
		{
			name: "ok2",
			args: args{2,1,1,1,1},
			want: "YES",
		},
		{
			name: "ok3",
			args: args{1,2,1,2,1},
			want: "YES",
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")
			fmt.Fprintf(f, "%d\n%d\n%d\n%d\n%d", tt.args[0], tt.args[1], tt.args[2], tt.args[3], tt.args[4])

			main()

			var got string
			f,_ = os.Open("output.txt")
			fmt.Fscan(f, &got)

			if got != tt.want{
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}