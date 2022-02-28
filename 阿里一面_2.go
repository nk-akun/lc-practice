package main

// 两个线程交替打印
// 线程1打印1,3,5,7,...
// 线程2打印2,4,6,8,...

import "fmt"

func print1(sig1 chan bool, sig2 chan bool, msg string) {
	for i := 1; i <= 100; i += 2 {
		if exist := <-sig1; exist {
			fmt.Println(i, msg)
			sig2 <- true
		}
	}
}

func print2(sig1 chan bool, sig2 chan bool, msg string, stop chan bool) {
	for i := 2; i <= 100; i += 2 {
		if exist := <-sig2; exist {
			fmt.Println(i, msg)
			sig1 <- true
		}
	}
	stop <- true
}

func main() {
	sig1 := make(chan bool, 1)
	sig2 := make(chan bool, 1)
	stop := make(chan bool, 1)
	go print1(sig1, sig2, "first")
	go print2(sig1, sig2, "second", stop)
	sig1 <- true
	<-stop
}
