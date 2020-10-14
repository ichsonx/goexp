/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2020/7/28 9:05
 * @Description:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("it will close after 10 secs...")
	time.Sleep(10 * time.Second)
}
