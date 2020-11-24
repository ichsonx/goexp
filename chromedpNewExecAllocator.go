// 2020-10-14
// 当前结合gooogle doc 查看的chromedp最新的启动调用实践
// 之前会有使用runner来启动chrome，现在应该弃用了
// 1、目前都是先使用NewExecAllocator/NewRemoteAllocator来启动chrome实例
//    （直接使用chromedp.Run()会默认使用DefaultExecAllocatorOptions()的设置来启动chrome实例（第一次Run()时）
//    可通过ExecAllocatorOption（或chromedp.Flag()函数）来设定chrome实例启动参数（NewRemoteAllocator无效）
// 2、NewExecAllocator为本地启动chrome实例，NewRemoteAllocator为连接并调用远程的已启动的chrome实例
// 3、本代码是使用NewExecAllocator的，实例化chrome后，使用返回的context来获得chromedp的context。
//    然后就可通过chromedp的context来运行任务了。
// 4、复用同一个chrome实例，打开多个tab标签页的示例：https://godoc.org/github.com/chromedp/chromedp#example-NewContext--ManyTabs
// 5、复用同一个chrome实例的同一个tab标签页示例：https://godoc.org/github.com/chromedp/chromedp#example-NewContext--ReuseBrowser
// 6、其实（4、5）主要是针对context的使用，用同一个chromedp的context来Run()会复用同一个tab标签页。
//    使用子context的话则会创建新的tab标签页（需要设置timeout context来关闭tab，否则好可能爆内存）
package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), defaultOpts()...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// ensure that the browser process is started
	// 创建chrome实例，判断是否启动成功
	if err := chromedp.Run(taskCtx); err != nil {
		panic(err)
	}
}

func defaultOpts() []chromedp.ExecAllocatorOption {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Headless,
		chromedp.NoSandbox,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("remote-debugging-port", 9222),
	)
	return opts
}
