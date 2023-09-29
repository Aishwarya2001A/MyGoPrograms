package main

import (
	"context"
	"fmt"
	"time"
)

func myfunc(ctx context.Context) {
	fmt.Println("context learning")
}
func myfunc1(c context.Context, arg string) string {

	time.Sleep(2 * time.Second)
	fmt.Println(arg)
	return arg
}

func main() {
	ctx := context.TODO()
	myfunc(ctx)

	c := context.Background()
	myfunc1(c, "aish")

	fmt.Println("Bye")

}
