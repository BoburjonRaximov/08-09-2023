package main

import (
	"context"
	"fmt"
	"playground/cpp-bootcamp/api"
	"playground/cpp-bootcamp/api/handler"
	"playground/cpp-bootcamp/config"
	"playground/cpp-bootcamp/pkg/logger"
	"playground/cpp-bootcamp/storage/db"
	"runtime"
	"sync"
	"time"
)

func main() {

	//context
	// withCancel()

	//
	//token
	// m := make(map[string]interface{})
	// m["user_id"] = "22cb933c-6c4c-47d3-9a49-80c9c4f2ad14"
	// m["branch_id"] = "b6ded900-b4f9-4df7-9e32-ad257d30c3fe"
	// m["name"] = "Alex"
	// token, _ := helper.GenerateJWT(m, time.Hour, "mySecretKey")
	// fmt.Println(token, err)
	// claims, _ := helper.ExtractClaims(token, "mySecretKey")
	// info, _ := helper.ParseClaims(token, "mySecretKey")
	// fmt.Println(info)
	// request := &models.CreatePerson{
	// 	Name: faker.FirstName(),
	// 	Job:  faker.LastName(),
	// 	Age:  rand.Intn(100),
	// }
	// fmt.Println(request)
	// return
	cfg := config.Load()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := db.NewStorage(context.Background(), cfg)
	if err != nil {
		return
	}

	h := handler.NewHandler(strg, log)

	r := api.NewServer(h)
	r.Run()

}

//
func withCancel() {
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// // cancel()
	// go func() {
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	timeConsumingFunc(ctx, 5*time.Second, "hello")
}

//

//

//

//

//

//

//

//
//select
func Select() {
	var ch1, ch2 = make(chan interface{}), make(chan interface{})
	// go sendMsg(ch1, false)
	go sendMsg(ch2, true)

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				break
			}
			fmt.Println("from ch1:", msg)
		case msg, ok := <-ch2:
			if !ok {
				break
			}
			fmt.Println("from ch2:", msg)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("waiting")
		}
	}
}
func sendMsg(ch chan interface{}, char bool) {
	if char {
		for i := 'a'; i < 'a'+30; i++ {
			ch <- string(i)
			time.Sleep(499 * time.Millisecond)
		}
		close(ch)
	} else {
		for i := 0; i < 30; i++ {
			ch <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(ch)
	}
}

//goroutine
func G1() {
	go fmt.Println("hello")

	fmt.Println("hi")
	fmt.Println("how are you?")
	time.Sleep(time.Second)
}

//WaitGroup
func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("hello")
	}()

	fmt.Println("hi")
	wg.Wait()
}

//Channel
func Channel() {
	ch := make(chan bool)
	go func() {
		fmt.Println("hello", time.Now())
		ch <- true

		ch <- false
		fmt.Println("send 2")
	}()
	fmt.Println("hi")
	time.Sleep(time.Second)
	fmt.Println(<-ch, time.Now())

	fmt.Println(<-ch, time.Now())
	time.Sleep(time.Second)

}
func ChannelClose() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello")
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	for {
		v, open := <-ch
		if !open {
			break
		}
		fmt.Println(v)
	}
	fmt.Println("hi")
}
func ChannelNil() {
	var ch chan bool
	go func() {
		close(ch)  // panic
		ch <- true //block
		fmt.Println("hello")
	}()
	// fmt.Println(<-ch) //block
	fmt.Println("hi")
	for {
	}
}

func chanRW() {
	ch := make(chan int)
	go readChan(ch)
	println("writing")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("num:", runtime.NumCPU())
	// go writeChan(ch)

	// println("reading")
	// for v := range ch {
	// 	fmt.Println(v)
	// }

}
func readChan(ch chan int) {
	println("reading")

	for v := range ch {
		fmt.Println(v)
	}
}
func writeChan(ch chan int) {
	println("writing")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func BufferedChan() {
	ch := make(chan int, 100)
	go func() {
		fmt.Println("writing time:", time.Now())
		ch <- 1
		fmt.Println("writing2 time:", time.Now())
		ch <- 2
		fmt.Println("writing end time:", time.Now())
		close(ch)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("reading time:", time.Now())
	fmt.Println(<-ch)

	time.Sleep(5 * time.Second)

	fmt.Println("hi")
}

func timeConsumingFunc(ctx context.Context, d time.Duration, message string) {
	for {
		select {
		case <-time.After(d):
			// name := ctx.Value("name")
			// fmt.Println("name:", name)
			fmt.Println(message)
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			return
		}
	}

}
