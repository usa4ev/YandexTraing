package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T){
	type args [4]string
	type want [3]string
	
	tests := []struct{
		name string
		args args
		want want
	}{
		{
			name: "",
			args: args{"8(495)430-23-97",
					"+7-4-9-5-43-023-97",
					"4-3-0-2-3-9-7",
					"8-495-430"},
			want: want{
				"YES",
				"YES",
				"NO",				
			},

		},

		{
			name: "",
			args: args{"86406361642",
					"83341994118",
					"86406361642",
					"83341994118"},
			want: want{
				"NO",
				"YES",
				"NO",				
			},

		},
		{
			name: "",
			args: args{"+78047952807",
					"+78047952807",
					"+76147514928",
					"88047952807"},
			want: want{
				"YES",
				"NO",
				"YES",				
			},

		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")

			fmt.Fprintf(f, "%s\n%s\n%s\n%s", tt.args[0],  tt.args[1], tt.args[2], tt.args[3])

			main()

			f,_ = os.Open("output.txt")

			got := want{}
			fmt.Fscanf(f, "%s\n%s\n%s", &got[0], &got[1], &got[2])

			for i,v := range got{
				if v != tt.want[i]{
					t.Errorf("wrong value: %v, want: %v", v, tt.want[i])
				}
			}
		})
	}
}