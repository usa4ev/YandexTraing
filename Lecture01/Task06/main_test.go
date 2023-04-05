package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T){
	type args [4]string
	tests := []struct{
		name string
		args args
		want [2]string
	}{
		{
			name: "test1",
			args: args{"10", "2", "2", "10"},
			want: [2]string{"4", "10"},
		},
		{
			name: "test2",
			args: args{"5", "7", "3", "2"},
			want: [2]string{"9", "5"},
		},
		{
			name: "test21",
			args: args{"3", "2", "5", "7"},
			want: [2]string{"9", "5"},
		},
		{
			name: "test3",
			args: args{"10", "10", "10", "10"},
			want: [2]string{"20", "10"},
		},
		{
			name: "test4",
			args: args{"3", "4", "10", "1"},
			want: [2]string{"4", "10"},
		},
		{
			name: "test5",
			args: args{"1", "10", "3", "4"},
			want: [2]string{"4", "10"},
		},
		{
			name: "test6",
			args: args{"3", "4", "4", "4"},
			want: [2]string{"7", "4"},
		},
		{
			name: "test7",
			args: args{"4", "4", "4", "3"},
			want: [2]string{"7", "4"},
		},
		{
			name: "test8",
			args: args{"3", "4", "4", "5"},
			want: [2]string{"8", "4"},
		},
		{
			name: "test9",
			args: args{"4", "5", "4", "3"},
			want: [2]string{"8", "4"},
		},
		{
			name: "test10",
			args: args{"5", "7", "5", "4"},
			want: [2]string{"11", "5"},
		},
		{
			name: "test11",
			args: args{"5", "4", "5", "7"},
			want: [2]string{"11", "5"},
		},
		{
			name: "test12",
			args: args{"5", "7", "5", "5"},
			want: [2]string{"12", "5"},
		},
		{
			name: "test13",
			args: args{"5", "5", "5", "7"},
			want: [2]string{"12", "5"},
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")
			fmt.Fprintf(f, "%s %s %s %s", tt.args[0], tt.args[1], tt.args[2], tt.args[3],)
		
			main()

			f,_ = os.Open("output.txt")

			var r1,r2 string 
			fmt.Fscanf(f,"%s %s", &r1, &r2)

			if r1 != tt.want[0] || r2 != tt.want[1]{
				t.Errorf("wrong res: got %v,%v; want: %v,%v", r1 ,r2, tt.want[0], tt.want[1])
			}
		})
	}
}