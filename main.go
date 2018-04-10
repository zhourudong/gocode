package main

import (
	"flag"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var uaList = []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36 Core/1.47.277.400 QQBrowser/9.4.7658.400",

	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36 SE 2.X MetaSr 1.0",

	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.154 Safari/537.36 LBBROWSER",

	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36 TheWorld 7"}

type resource struct {
	//	构建资源
	url    string
	target string
	start  int
	end    int
}

func rurlResource() []resource {
	var res []resource
	r1 := resource{
		// 首页
		url:    "http://127.0.0.1:5000",
		target: "",
		start:  0,
		end:    0,
	}

	r2 := resource{
		//	列表页
		url:    "http://127.0.0.1:5000/list/${id}.html",
		target: "${id}",
		start:  1,
		end:    21,
	}

	r3 := resource{
		//	详情页
		url:    "http://127.0.0.1:5000/movie/${id}.html",
		target: "${id}",
		start:  1,
		end:    12924,
	}

	res = append(append(append(res, r1), r2), r3)

	return res

}

func buildUrl(res []resource) []string {
	var list []string

	for _, resItem := range res {
		if len(resItem.target) == 0 {
			list = append(list, resItem.url)
		} else {
			for i := resItem.start; i <= resItem.end; i++ {
				urlStr := strings.Replace(resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}
func makeLog(current, refer, ua string) string {
	//	返回日志
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramStr := u.Encode()

	logTemplate := "127.0.0.1 - - [05/Apr/2018:13:35:29 +0800] \"GET /dig{$paramStr} HTTP/1.1\" 200 43 \"http://127.0.0.1:5000/admin/moviecol/list/1/\" \"{$ua}\" \"-\""
	log := strings.Replace(logTemplate, "{$paramStr}", paramStr, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)

	return log

}

func randInt(min, max int) int {
	//	返回随机值函数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if min > max {
		return max
	}

	return r.Intn(max-min) + min

}
func main() {

	total := flag.Int("total", 100, "模拟多少行日志")
	filePath := flag.String("filepath", "/Users/zrd/tmp/nginx-1.12.2/logs/dig.log", "文件路径")

	flag.Parse()

	// 需要构造出真实的网站url集合
	res := rurlResource()

	list := buildUrl(res)

	// 按照要求，生成$total行日志内容，源自于这个网站url的集合
	logStr := ""
	for i := 0; i <= *total; i++ {
		currentUrl := list[randInt(0, len(list))]
		referUrl := list[randInt(0, len(list))]
		ua := uaList[randInt(0, len(uaList))]

		logStr = logStr + makeLog(currentUrl, referUrl, ua) + "\n"
		//	写到文件
		//ioutil.WriteFile(*filePath, []byte(logStr), 0644) 覆盖写
	}
	fd, _ := os.OpenFile(*filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 600)
	fd.Write([]byte(logStr))
	fd.Close()

	println("done\n")

}
