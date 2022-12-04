package th1

/*
 给定一个字符串数组
 [“I”,“am”,“stupid”,“and”,“weak”]
 用 for 循环遍历该数组并修改为
 [“I”,“am”,“smart”,“and”,“strong”]

TestJump 实现
 
*/

import (
	"fmt"
	"testing"
	"time"
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


