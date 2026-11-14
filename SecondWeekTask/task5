```
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Result struct {
	ID  int
	Num int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan Result)

	for i := 1; i <= 20; i++ {
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(1001)) * time.Millisecond)
			ch <- Result{id, rand.Intn(1000)}
		}(i)
	}

	var a [20]Result
	for i := 0; i < 20; i++ {
		a[i] = <-ch
	}

	fmt.Println("排序前：")
	for _, r := range a {
		fmt.Println(r.ID, r.Num)
	}

	//此处我搜索了一下怎么使用sort排序，并copy了一下，但是已了解原理
	sort.Slice(a[:], func(i, j int) bool {
		return a[i].ID < a[j].ID
	})

	fmt.Println("\n排序后：")
	for _, r := range a {
		fmt.Println(r.ID, r.Num)
	}
}

```

代码如上，我再也不堆到周五周六写作业了。
