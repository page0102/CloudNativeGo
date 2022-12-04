package th1

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T){

	Myarray1 := [5]string{"I","am","stupid","and","weak"}
	for i,s := range Myarray1{
		fmt.Println(i,s)
		if s == "stupid"{
			Myarray1[i] = "smart"
		}else if s == "weak"{
			Myarray1[i] = "strong"
		}
	}

	for i,s := range Myarray1{
		fmt.Println(i,s)
	}
}
