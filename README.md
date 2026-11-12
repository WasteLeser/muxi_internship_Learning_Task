
# 个人木犀实习期学习任务仓库 

## ⭐ Star me on GitHub: your support motivates me a lot! 🙏


**READ仅用于记录进度，学习内容与作业将放入文件夹用于分类与展示**




个人博客在重建中

<details>
<summary>第一周主线任务</summary>

- [x] 注册[GitHub](https://github.com/WasteLeser)、ChatGPT账号并成功访问


---

- [x] 学习Markdown的写法，重点关注代码块、多级标题、数字/圆点列表、加粗、删除线......
  
    ---

- [x] 利用Markdown编写简单的个人博客，并利用github pages部署

    ---

- [x]  学习go的数组、切片，并完成以下任务：
     - 给定s := make([]byte, 5)，那么len(s) 和 cap(s) 分别是多少？令s = s[2:4]，len(s) 和 cap(s) 又分别是多少？
     - 比较字符串"hello，世界"的长度和for range该字符串的循环次数

---

- [x]  将学习的笔记/代码上传到GitHub，建议新建一个学习仓库，把每周的任务分文件夹上传


</details>



<details>
<summary>第二周主线任务</summary>

- [x]  理解map和slice内存扩容的区别

---

- [x] 利用闭包实现一个计数器

---

- [x] 基于下面这个结构体实现四个方法：


- Average：返回平均分；Scores 为空时返回 0。


- AddScore：追加一个成绩，调用方能看见这个改动。


- Top：返回最高的 n 个成绩，从高到低。n 超过长度时返回全部，n <= 0 时返回 nil。不能修改 s.Scores 本来的顺序，也不能让返回的切片存在影响原数据的可能。


- Merge：合并两个人的成绩，返回一个新的 Student（Name 用接收者的）。合并结果和两个原对象必须互不影响。


        type Student struct {
            Name   string
            Scores []int // 记录若干项成绩
        }


  
---


- [x] 思考以下程序在并发出现panic的原因，并解决
        package main
      
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
              //进行异步消费
              go func() {
                 m := consumeMSG[:ConsumeNum]
                 fn(m)
              }()
              // 更新上次消费时间
              lastConsumeTime = time.Now()
              // 清除插入的数据
              consumeMSG = consumeMSG[ConsumeNum:]
           } else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
              // 如果距离上次消费已经超过5分钟且有未处理的消息
              if len(consumeMSG) > 0 {
                 //进行异步消费 
                 go func() {
                    m := consumeMSG[:ConsumeNum]
                    fn(m)
                 }()
                 // 更新上次消费时间
                 lastConsumeTime = time.Now()
                 // 清空插入的数据
                 consumeMSG = consumeMSG[ConsumeNum:]
                }
             }
            }
          }

        func fn(m []MSG) {
            fmt.Printf("本次消费了%d条消息\n", len(m))
        }


- [x] 使用for循环生成20个goroutine,每个goroutine随机休眠0~1000ms，并向一个channel传入随机数和goroutine编号(从1-20)，等待这些goroutine都生成完后，想办法给这些随机数按照编号进行排序(输出排序前和排序后的结果,要求不使用额外的空间存储这20个数据)


---

  
- [x] 拓展题：只使用channel进行并发控制，交叉打印下面两个字符串（要求一个打印完，另一个会继续打印）
"ABCDEFGHIJKLMNOPQRSTUVWXYZ" "0123..."
得到："AB01CD23EF34..."（不会的话可以降低难度，使用其他的并发控制工具）


</details>


<details>
<summary>第三周主线任务</summary>

> 请输入文本


</details>


<details>
<summary>第四周主线任务</summary>

> 请输入文本


</details>


<details>
<summary>神秘</summary>

<details>
<summary>神秘</summary>

 <details>
<summary>神秘</summary>

<details>
<summary>神秘</summary>

<details>
<summary>神秘</summary>

<details>
<summary>神秘</summary>

<details>
<summary>神秘</summary>

<details>
<summary>神秘</summary>

啥也没有


</details>


</details>


</details>


</details>


</details>


</details>



</details>


</details>
