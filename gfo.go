package main

import "fast_gin/core"

//
//type isALock struct {
//	t int32
//}
//
//func (l *isALock) Lock() {
//	for {
//		println("正在尝试获得锁.......")
//		if ok := atomic.CompareAndSwapInt32(&l.t, 0, 1); ok {
//			return
//		}
//	}
//}
//func (l *isALock) Unlock() {
//	atomic.StoreInt32(&l.t, 0)
//}
//
//type mapCache struct {
//	data map[string]string
//	l    *isALock
//}
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	//go func(ctx context.Context) {
//	//	select {
//	//	case <-ctx.Done():
//	//		return
//	//	}
//	//}(ctx)
//	wg := sync.WaitGroup{}
//	c := &mapCache{
//		data: make(map[string]string),
//		l:    &isALock{},
//	}
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		c.l.Lock()
//		defer c.l.Unlock()
//		c.data["key"] = "value"
//		println(c.data["key"])
//	}()
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		c.l.Lock()
//		//defer c.l.Unlock()
//		c.data["key"] = "value"
//		println(c.data["key"])
//	}()
//	wg.Wait()
//	<-ctx.Done()
//}

//
//import (
//	"fmt"
//	"math/rand"
//)
// 协程A 往channel传值 协程B 接受channel的值并计算 协程C 打印结果值
// 考察点：waitgroup、channel
//func main() {
//
//	var wg sync.WaitGroup
//
//	wg.Add(3)
//	numChan := make(chan int)
//	resChan := make(chan int)
//
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 100; i++ {
//			numChan <- rand.Intn(100)
//		}
//		close(numChan)
//	}()
//
//	go func() {
//		defer wg.Done()
//		res := 0
//		for i := range numChan {
//			res += i*i
//		}
//		resChan <- res
//	}()
//
//	go func() {
//		defer wg.Done()
//		res := <- resChan
//		fmt.Println(res)
//	}()
//
//	wg.Wait()
//}

// 单个chan是阻塞的
//package main
//
//import (
//	"fmt"
//	"math/rand"
//)
//
//// 使用两个chan来进行计算，
//func main() {
//	random := make(chan int)
//	done := make(chan bool)
//
//	go func() {
//		res := 0
//		for {
//			num, ok := <-random
//			if ok {
//				fmt.Print(num, "   ")
//				res += num * num
//			} else {
//				break
//			}
//		}
//		fmt.Println(res)
//		done <- true
//	}()
//
//	go func() {
//		defer close(random)
//
//		for i := 0; i < 5; i++ {
//			random <- rand.Intn(5)
//		}
//	}()
//
//	//<-done
//	//close(done)
//	select {
//	case <-done:
//		return
//	}
//}

//要求每秒钟调用一次proc并保证程序不退出?
//package main
//
//import (
//"fmt"
//"time"
//)
//
//func main() {
//	go func() {
//		// 1 在这里需要你写算法
//		// 2 要求每秒钟调用一次proc函数
//		// 3 要求程序不能退出
//		for {
//			go func() {
//				defer func() {
//					if err := recover(); err != nil {
//						fmt.Println(err)
//						return
//					}
//				}()
//				proc()
//			}()
//			time.Sleep(1 * time.Second)
//		}
//
//	}()
//
//	select {}
//}
//
//func proc() {
//	panic("ok")
//}

// 使用Ticker
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	go func() {
//		// 1 在这里需要你写算法
//		// 2 要求每秒钟调用一次proc函数
//		// 3 要求程序不能退出
//		ticker := time.NewTicker(1 * time.Second)
//
//		for {
//			select {
//			case <-ticker.C:
//				go func() {
//					defer func() {
//						if err := recover(); err != nil {
//							fmt.Println(err)
//						}
//					}()
//					proc()
//				}()
//			}
//		}
//	}()
//
//	select {}
//}
//
//func proc() {
//	panic("ok")
//}

// 实现两个协程交互打印数字和字母
// package main
//
// import "fmt"
//
//	func main() {
//		letterChan := make(chan bool)
//		numChan := make(chan bool)
//		doneChan := make(chan bool)
//
//		go func() {
//			num := 1
//			for {
//				select {
//				case <-numChan:
//					fmt.Println(num)
//					num += 1
//					letterChan <- true
//				}
//			}
//		}()
//
//		go func() {
//			letter := 'A'
//			for {
//				select {
//				case <-letterChan:
//					if letter > 'Z' {
//						doneChan <- true
//					} else {
//						fmt.Println(string(letter))
//						letter++
//						numChan <- true
//					}
//				}
//			}
//		}()
//
//		numChan <- true
//
//		select {
//		case <-doneChan:
//			return
//		}
//	}
func main() {
	core.EsConnect()
}
