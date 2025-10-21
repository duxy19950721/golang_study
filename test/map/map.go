package main

import (
	"fmt"
	"time"
)

var a map[string]string = map[string]string{
	"name": "duxy",
	"age":  "30",
}

func main() {
	// loop()
	// add()
	// remove()
	bingfa()
}

func loop() {
	b := map[string]string{
		"name": "duxy",
		"age":  "30",
	}
	for k, v := range b {
		fmt.Println(k, v)
	}

	fmt.Println("len", len(b))
}

func add() {
	b := map[string]string{
		"name": "duxy",
		"age":  "30",
	}
	b["aaaa"] = "123"
	fmt.Println("len", len(b))
}

func remove() {
	b := map[string]string{
		"name": "duxy",
		"age":  "30",
	}
	delete(b, "name")
	fmt.Println("len", len(b))
}

func bingfa() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"]++
		}
	}()

	go func() {
		for {
			m["a"]++
			fmt.Println(m["a"])
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}
