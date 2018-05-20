
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

	fmt.Println("以指针方式修改值(地址引用)")
	i := 9
	j := 5
	product := 0
	fmt.Printf("地址: i=%p, j=%p, product=%p\n", &i, &j, &product)

	swap_producr(&i, &j, &product)

	fmt.Println(i, j, product) // 5 9 45

	fmt.Println()
	fmt.Println("以传值方式修改值")
	// 非引用交换值,  对于大数组特别🎺性能内存
	i1 := 9
	j1 := 5
	product1 := 0

	fmt.Printf("地址: i1=%p, j1=%p, product1=%p\n", &i1, &j1, &product1)
	i1, j1, product1 = swap_producr_no_ref_address(i1, j1, product1)
	fmt.Println(i1, j1, product1) // 5 9 45

}

func swap_producr(i, j, product *int) {
	fmt.Printf("指针类型接收到的地址: %p, %p,%p\n", i, j, product)
	if *i > *j {
		*i, *j = *j, *i
	}
	*product = *i * *j
}

func swap_producr_no_ref_address(i1, j1, product1 int) (int, int, int) {
	fmt.Printf("值类型接收到的地址: %p, %p,%p\n", &i1, &j1, &product1)
	if i1 > j1 {
		i1, j1 = j1, i1
	}
	product1 = i1 * j1
	return i1, j1, product1

}
```


# struct类型

```go
package main

import "fmt"

func main() {

	zhangsan := person{"zhansan", 1111} // person类型值

	lisi := new(person) // 指向person的指针
	lisi.name, lisi.id = "lisi", 123  // (*lisi).name = "lisi"  (*lisi).id= 123 

	wangwu := &person{"wangwu", 456} // 指向person的指针

	fmt.Printf("%v \n", zhangsan)      // {zhansan 1111}
	fmt.Printf("%v  %v", lisi, wangwu) // &{lisi 123}  &{wangwu 456}

}

type person struct {
	name string
	id   int
}

// new(Type) = &Type{}   所做的操作是相同的

```
