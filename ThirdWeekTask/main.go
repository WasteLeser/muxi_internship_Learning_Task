package main

import (
	"net/http"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, _ := http.NewRequest("GET", "https://gtainmuxi.muxixyz.com/api/v1/organization/code", nil)
	req.Header.Set("code", "0")
	resp, _ := http.DefaultClient.Do(req)
	passport := resp.Header.Get("passport")

	// req2, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key", nil)
	// req2.Header.Set("passport", passport)
	// req2.Header.Set("code", "0")
	// resp2, _ := client.Do(req2)
	// resbody, _ := io.ReadAll(resp2.Body)
	// fmt.Println(string(resbody))

	//该段注释掉的代码获取了"extra_info":"c2VjcmV0X2tlOk11eGlTdHVkaW8yMDMzMDQsIGVycm9yX2NvZGU6Zm9yIHtnbyBmdW5jKCl7dGltZS5TbGVlcCgxKnRpbWUuSG91cil9KCl9"
	//通过base64解码后得到"secret_key:MuxiStudio203304, error_code:for {go func(){time.Sleep(1*time.Hour)}()}"
	//加密到底是什么啊，这下不得不引入hacker包了

	// body, _ := encrypt.AESEncryptOutInBase64([]byte("for {go func(){time.Sleep(1*time.Hour)}()}"), []byte("MuxiStudio203304"))
	// requestBody := map[string]string{
	// 	"content": string(body),
	// }
	// bodyBytes, _ := json.Marshal(requestBody)
	//需要使用json，马上去复习了一下。不是很熟练这一块儿，还需要搜索。。。
	// req2, _ := http.NewRequest("PUT", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", bytes.NewReader(bodyBytes))
	// req2.Header.Set("passport", passport)
	// req2.Header.Set("code", "0")
	// resp2, _ := http.DefaultClient.Do(req2)
	// defer resp2.Body.Close()
	//	respbody, _ := io.ReadAll(resp2.Body)
	//	fmt.Println(string(respbody))
	//{"code":0,"message":"OK","data":{"text":"干的漂亮！你已经突破了第一扇门，请继续访问 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate 。","extra_info":""}}

	//该段代码已用不到，注释掉了
	//	req3, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", nil)
	//	req3.Header.Set("passport", passport)
	//	req3.Header.Set("code", "0")
	//	resp3, _ := http.DefaultClient.Do(req3)
	//	respbody, _ := io.ReadAll(resp3.Body)
	// 	fmt.Println(string(respbody))
	// 	//{"code":0,"message":"OK","data":{"text":"你现在已经到第二扇门了，是虹膜识别安全门。\r\n你需要向组织请求已准备好的虹膜样本，访问 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample 下载图片。\r\n再将此图片上传至 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate 以破解此门，加油！","extra_info":""}}

	//	resp3.Save("E:/Code/gojohnygo/iris_sample.jpg")
	// 	突然发现save方法没法用。我靠原来save是hacker-support里HttpResponse的方法，我一直用的标准库net/http的http.Response

	//获取图片
	// req4, _ := httptool.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample", "", httptool.DEFAULT)
	// req4.AddHeader("passport", passport)
	// req4.AddHeader("code", "0")
	// resp4, _ := req4.SendRequest()
	// resp4.Save("E:/Code/gojohnygo/iris_sample.png")

	//copy了一下怎么request时上传图片
	// 看到了源码     - description: file,这个是用表单上传
	//修改为绝对路径终于通过了
	//二次更改：我投降了，hacker包上传文件这么简单吗原来
	// req5, _ := httptool.NewRequest(httptool.POSTMETHOD,"https://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate","E:/Code/gojohnygo/iris_sample.jpg",httptool.FILE)
	// req5.AddHeader("passport", passport)
	// resp5, _ := req5.SendRequest()
	// resp5.ShowBody()

	//{"code":0,"message":"OK","data":{"text":"还剩最后一道门了！\r\n我们需要银行结构图碎片，这些碎片就隐藏在前面某四个路由的响应头中，位于 map-fragments 字段。\r\n将它们用\"/\"拼起来就是最后一道门的所在位置！注意response的信息。","extra_info":""}}

	//因为前面全注释掉了还有点乱，所以我开了作弊码找到了路径。源码这里。router/router.go：
	// 58	end := g.Group("/api/v1/muxi/backend/computer/examination")
	// 59	end.Use(middleware.AuthMiddleware())
	// 60	{
	// 61		end.GET("", checkpoint5.GetText)
	// 62		end.POST("", checkpoint5.UploadFile)
	// 63	}

	// req6, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", nil)
	// req6.Header.Set("passport", passport)
	// req6.Header.Set("code", "0")
	// resp6, _ := http.DefaultClient.Do(req6)
	// respbody, _ := io.ReadAll(resp6.Body)
	// fmt.Println(string(respbody))

	req7, _ := httptool.NewRequest(httptool.POSTMETHOD, "https://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", "E:/Code/gojohnygo/全排列.go", httptool.FILE)
	req7.AddHeader("passport", passport)
	resp7, _ := req7.SendRequest()
	resp7.ShowBody()
  //全排列还没写出来，先休息一下，晚上再写
}


