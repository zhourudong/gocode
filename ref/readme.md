
# 指针

```go

package main

import "fmt"

func main() {
	z := 37    // int
	pi := &z   // pi 类型为 *int (指向int型的指针)
	ppi := &pi // ppi 类型为 **int (指向int型的指针)

	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi)              // z=37, *pi=37, **ppi=37
	fmt.Printf("type:z=%#T, type:*pi=%T, type:**ppi=%T\n", z, *pi, **ppi) //

	**ppi++                                                  // 等同于  (*(*ppi))++ 和 *(*ppi)++; 操作的结果z++和 *pi++ 相同
	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi) // z=38, *pi=38, **ppi=38
	z++
	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi) // z=39, *pi=39, **ppi=39

	(*(*ppi))++
	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi) // z=40, *pi=40, **ppi=40
	*(*ppi)++
	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi) // z=41, *pi=41, **ppi=41

	(*pi)++
	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi) // z=42, *pi=42, **ppi=42
	*(pi)++
	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi) // z=43, *pi=43, **ppi=43
	*pi++
	fmt.Printf("z=%#v, *pi=%#v, **ppi=%#v\n", z, *pi, **ppi) // z=44, *pi=44, **ppi=44

}

```

![指针](https://raw.githubusercontent.com/zhourudong/gocode/master/ref.jpg)

# 交换值

```go
package main

import "fmt"

func main() {

	i := 9
	j := 5
	product := 0

	swap_producr(&i, &j, &product)
	fmt.Println(i, j, product) // 5 9 45

	// 非引用交换值,  对于大数组特别🎺性能内存
	i1 := 9
	j1 := 5
	product1 := 0

	i1, j1, product1 = swap_producr_no_ref_address(i1, j1, product1)
	fmt.Println(i1, j1, product1) // 5 9 45

}

func swap_producr(x, y, product *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
	*product = *x * *y
}

func swap_producr_no_ref_address(x, y, product int) (int, int, int) {
	if x > y {
		x, y = y, x
	}
	product = x * y
	return x, y, product

}
```
