源代码运行后报错panic

> panic: runtime error: slice bounds out of range [:5] with capacity 3

一眼切片超过容量了

又发现常量ConsumeNum = 5

通过ctrl+点击的方式查看用到该常量的切片有consumeMSG[]

我看不懂了，开始阅读整体代码。。。

大致理解了代码

发现距离上次消费已经超过5分钟且有未处理的消息时消费的物品或不足五件此时操作仍以五件进行操作，

遂改变了消费模块的第二个分支中的

      `  m := consumeMSG[:ConsumeNum]`
	  
为

     `   m := consumeMSG[:len(consumeMSG)]`

依旧

> panic: runtime error: slice bounds out of range [:5] with capacity 3

通过前面思路得到大概率仍是第二个分支内的错误，发现

```
        		// 清空插入的数据
				consumeMSG = consumeMSG[ConsumeNum:]
```

此处仍也需要修改 

我将ConsumeNum更改为了len(consumeMSG)

我要崩溃了。

我重拾了信心。

我阅读了代码。

我没看懂。

我决定调试。

我在每个模块内加入了输出特定的数字，

发现依旧是if else分支内有问题。

且第二个分支内

`
         if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
`	

的数字从未被输出过

我检查了条件，好吧好像我没找出来而且panic是切片那边的

顺便我看到了第一个分支

```
    	if len(consumeMSG) >= ConsumeNum {
			//进行异步消费
			go func() {
				m := consumeMSG[:ConsumeNum]
				fn(m)
			}()
			// 更新上次消费时间
			lastConsumeTime = time.Now()
			// 清除插入的数据
			consumeMSG = consumeMSG[ConsumeNum:]
		}
```

内有一个我学习协程时被经常提到的警告，go内外都有使用consumeMSG

哦我看懂了，我尝试把m移出来变成consumeMSG的一份copy

但我不太会啊，我百度了一下。

这是更改后
```
		if len(consumeMSG) >= ConsumeNum {
			m := append([]MSG(nil), consumeMSG[:ConsumeNum]...)
			//进行异步消费
			go func() {
				fn(m)
			}()
			// 更新上次消费时间
			lastConsumeTime = time.Now()
			// 清除插入的数据
			consumeMSG = consumeMSG[ConsumeNum:]
		}
```
我又发现第二个分支是一样的问题，

更改后代码发现能够运行了

```

package main

import (
	"fmt"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time // 记录上次消费的时间
	msgs := make(chan MSG)

	//这里源源不断的生产信息
	go func() {
		for i := 0; ; i++ {
			msgs <- MSG{
				ms: message{
					Topic:     "消费主题",
					Partition: 0,
					Offset:    0,
				},
				feedEvent: FeedEventDM{
					Type:    "grade",
					UserID:  i,
					Title:   "成绩提醒",
					Content: "您的成绩是xxx",
				},
			}
			//每次发送信息会停止0.01秒以模拟真实的场景
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//不断接受消息进行消费
	for msg := range msgs {

		// 添加新的值到events中
		consumeMSG = append(consumeMSG, msg)
		// 如果数量达到额定值就批量消费
		if len(consumeMSG) >= ConsumeNum {
			m := append([]MSG(nil), consumeMSG[:ConsumeNum]...)
			//进行异步消费
			go func() {
				fn(m)
			}()
			// 更新上次消费时间
			lastConsumeTime = time.Now()
			// 清除插入的数据
			consumeMSG = consumeMSG[ConsumeNum:]
		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			// 如果距离上次消费已经超过5分钟且有未处理的消息
			if len(consumeMSG) > 0 {
				m := append([]MSG(nil), consumeMSG[:ConsumeNum]...)

				//进行异步消费
				go func() {
					fn(m)
				}()
				// 更新上次消费时间
				lastConsumeTime = time.Now()
				// 清空插入的数据
				consumeMSG = consumeMSG[len(consumeMSG):]
			}
		}
	}
}

func fn(m []MSG) {
	fmt.Printf("本次消费了%d条消息\n", len(m))
}

```
