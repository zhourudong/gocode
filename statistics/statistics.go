package main

import (
	"net/http"
	"sort"

	"fmt"
	"strconv"
	"strings"

	"log"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type statistics struct {
	numbers []float64
	mean    float64 // 平均数
	median  float64 // 中位数
}

func getStats(numbers []float64) (stats statistics) {
	// 处理数据 求中间值 与平均数
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers)) // 求和
	stats.median = median(numbers)
	return stats
}
func sum(numbers []float64) (total float64) {
	// 求和
	for _, x := range numbers {
		total += x
	}
	return total
}
func median(numbers []float64) float64 {
	// 求中间数
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}
func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server:", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	// 处理表单
	err := request.ParseForm() // Must be called before writing response

	fmt.Fprint(writer, pageTop, form) // 写入 resp
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers) // // 处理数据 求中间值 与平均数
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}
func processRequest(r *http.Request) ([]float64, string, bool) {
	// 处理form提交的数据
	var numbers []float64
	if slice, found := r.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			// 转换成float64数值型
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is incalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // 第一次没有数据被显示
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
