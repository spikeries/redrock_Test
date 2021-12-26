package main

//func main() {
//	var mu sync.Mutex
//
//	go func() {
//		fmt.Println("有点强人锁男")
//		mu.Lock()
//	}()
//
//	mu.Unlock()
//}
//go func 与mu.Unlock并行执行，导致mu.lock还未执行就开始执行mu.Unlock
//出现错误