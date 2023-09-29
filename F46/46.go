package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

const interval = 500
const shortDuration = 1 * time.Millisecond

type keyType string

func main() {

	// Context with Value
	key1 := keyType("Name")
	ctx1 := context.WithValue(context.Background(), key1, "aish")
	Context1(ctx1, key1)

	// Determining If a Context is Done
	ctx2, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * interval * time.Millisecond)
		cancel()
	}()
	Context2(ctx2)

	//Context with Cancel
	ctx3 := context.Background()
	ctx3, cancel = context.WithCancel(ctx3)

	time.AfterFunc(2*time.Second, cancel)

	Context3(ctx3, 6*time.Second, "progress")

	// Context with Deadline

	d := time.Now().Add(shortDuration)
	ctx4, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx4.Done():
		fmt.Println(ctx4.Err())
	}
	// Context with Timeout
	ctx5, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go Context5(ctx5)
	select {
	case <-ctx5.Done():
		fmt.Println("exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}

// Context with Value
func Context1(ctx1 context.Context, k1 keyType) {
	value := ctx1.Value(k1)
	if value != nil {
		fmt.Print("The context value is :", value)
		return
	}
	fmt.Print("Unable to find the context value")
}

// Context with Deadline
func Context5(ctx5 context.Context) {
	for {
		select {
		case <-ctx5.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("doing something amazing")
		}
		time.Sleep(500 * time.Millisecond)
	}

}

// Context with Cancel
func Context3(ctx3 context.Context, d time.Duration, name string) {
	select {
	case <-time.After(d):
		fmt.Print("Your name is ", name)
	case <-ctx3.Done():
		err := ctx3.Err()
		fmt.Print(err)
	}
}

// Determining If a Context is Done
func Context2(ctx2 context.Context) {
	ticker := time.NewTicker(interval * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			log.Println("run")
		case <-ctx2.Done():
			return
		}
	}

}
