/**
 * @Author: sonic
 * @File:  chromedp.go
 * @Date: 2020/10/15 14:30
 * @Description:
 */

package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
)

func main() {
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.Flag("default-browser-check", true))
	defer cancel()

	ctx, cancle := chromedp.NewContext(allocCtx)
	defer cancle()

	//ctx, cancle := chromedp.NewContext(context.Background())
	//defer cancle()

	if err := chromedp.Run(ctx); err != nil{
		panic(err)
	}

	var htmlstring string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.baidu.com"),
		chromedp.OuterHTML("html", & htmlstring),
		); err != nil{
		panic(err)
	}



	time.Sleep(time.Second * 10)

	fmt.Println(htmlstring)
}


