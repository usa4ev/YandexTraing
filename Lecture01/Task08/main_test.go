package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T){
	type args [4]int
	type want []int
	tests := []struct{
		name string
		args args
		want []int
	}{
		{
			name: "ok",
			args: args{1,3,3,2},
			want: want{5,7},
		},
		{
			name: "bad",
			args: args{1,5,1,2},
			want: want{-1},
		},
		{
			name: "okinv",
			args: args{3,1,2,3},
			want: want{5,7},
		},
		{
			name: "ones",
			args: args{1,1,1,1},
			want: want{1,3},
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")
			fmt.Fprintf(f, "%d\n%d\n%d\n%d", tt.args[0], tt.args[1], tt.args[2], tt.args[3])

			main()

			f,_ = os.Open("output.txt")
			var gotMin, gotMax int
			n,_ := fmt.Fscanf(f, "%d %d", &gotMin, &gotMax)
			if len(tt.want) != n || gotMin != tt.want[0] || (len(tt.want) == 2 && gotMax != tt.want[1]){
				t.Errorf("got max = %v; min = %v, want = %v", gotMax, gotMin, tt.want)
			}
		})
	}

}