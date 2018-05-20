# 连接字符串

```go
func ex1() {

	var buffer bytes.Buffer
	for {
		if piece, ok := getNextValidString(); ok {
			buffer.WriteString(piece)
		} else {
			break
		}
	}

	fmt.Print(buffer.String(), "\n")
}

```

# 查找一行字,首尾一个单词

```go 
func find_first_last_string() {
	line := "啊dsfsfbjhdsf fsdaf hvhjvg hvh ABCSD"
	f_space := strings.Index(line, " ") // 查找第一个空格位置
	first_word := line[:f_space]
	l_space := strings.LastIndex(line, " ")
	last_word := line[l_space+1:]
	fmt.Println(first_word, last_word)
}
```

# 字符串分割

```go
func strings_pkg() {
	names := "中国.北京.天津"
	//	分割字符串
	for _, name := range strings.Split(names, ".") {
		fmt.Println(name)
	}
}

```
