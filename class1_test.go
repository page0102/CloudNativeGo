package th1

/*
1.  TestJump 实现：
    给定一个字符串数组
    [“I”,“am”,“stupid”,“and”,“weak”]
    用 for 循环遍历该数组并修改为
    [“I”,“am”,“smart”,“and”,“strong”]
2.  TestChanel +（producer、consummer）实现：
    队列：
    队列长度 10，队列元素类型为 int
    生产者：
    每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
    消费者：
    每1秒从队列中获取一个元素并打印，队列为空时消费者阻塞
 
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

func producer(ch chan <-int ){

	for i:=0;i<10;i++{
		ch <- i
		time.Sleep(time.Second)
		fmt.Println("producer insert ", i)
	}
	close(ch)
}

func consummer(ch chan int){
	for i := range ch{
		time.Sleep(time.Second)
		fmt.Println("consummer got ", i)
	}
}


func TestChanel(t *testing.T) {

	ch := make(chan int,10)
	go producer(ch)
	consummer(ch)

}
