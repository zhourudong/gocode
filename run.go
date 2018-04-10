package run

import "flag"

func run() {
	// 三个参数 第一个是 --total 第二个是默认值 第三是 参数说明
	total := flag.Int("total", 100, "模拟多少行日志")
	filePath := flag.String("filepath", "/Users/zrd/tmp/nginx-1.12.2/logs/access.log", "文件路径")

	flag.Parse()

	print(*total, *filePath)
}

/*
	// 三个参数 第一个是 --total 第二个是默认值 第三是 参数说明
	total := flag.Int("total", 100, "模拟多少行日志")
	filePath := flag.String("filepath", "/Users/zrd/tmp/nginx-1.12.2/logs/access.log", "文件路径")

	flag.Parse()

	print(*total, *filePath)



# ----
(sx3.5.3) ➜  go_web   go build
(sx3.5.3) ➜  go_web  ls
go_web run.go
(sx3.5.3) ➜  go_web  ./go_web  --filepath=ssss
100ssss%


*/
