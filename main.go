package main

import (
	"fmt"
	"gpool/src"
)

type logger2 struct {
}

func (log logger2) Print(v ...interface{})  {
	fmt.Println(v)
}

func main()  {
	task := src.NewTask(func() error {
		fmt.Println("test")
		return nil
	})

	src.NewExecutor(task, src.NewPool(3), logger2{}).Processor()
}
