
package main

import (
"fmt"
)
//采用了素数筛的思路
//将原始素数的倍数从输入的数字中排除出去
//我自己的理解写到了注释中
func main() {
	get_prime(1000000)
}
func get_prime(n int) {
	origin, wait := make(chan int), make(chan struct{})
	go FilterPrime(origin, wait)
	for i := 2; i <= n; i++ {//不包括1，将（2,1000000)的数字轮流输入到origin
		origin <- i
	}
	close(origin)
	<-wait //等待同步 防止主线程退出
}

func FilterPrime(seq chan int, wait chan struct{}) {
	prime, ok := <-seq
	if !ok {
		close(wait)
		return
	}
	fmt.Println(prime) //打印素数
	out := make(chan int)
	go FilterPrime(out, wait) //启动另一个goroutine
	for num := range seq {//遍历seq
		if num%prime != 0 {//素数筛的思想，排除了原始素数的倍数后，将后面的数字传给out
			out <- num
		}
	}
	close(out)
}
//————————————————
//版权声明：本文为CSDN博主「chrispink_yang」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
//原文链接：https://blog.csdn.net/m0_37422289/article/details/92759872