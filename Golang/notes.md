# Go基础
## Go特性
    - 高性能、高并发
    - 语法简单、学习曲线平缓
    - 丰富的标准库
    - 晚上的工具链
    - 静态链接 静态链接（Static Linking）是指在编译阶段，将程序所需的所有库函数代码直接复制到最终生成的可执行文件中的过程。这使得程序可以在没有安装相应库的系统上运行。
    - 快速编译
    - 跨平台
    - 垃圾回收

## [Golang 需要避免踩的 50 个坑](https://juejin.cn/post/6844903816018542600)
    - 函数定义时 左大括号不能单独放一行 编译器会在)后加;
    - 如果在函数体代码中有未使用的变量，则无法通过编译，不过全局变量声明但不使用是可以的。即使变量声明后为变量赋值，依旧无法通过编译，需在某处使用它。
    - 如果你 import 一个包，但包中的变量、函数、接口和结构体一个都没有用到的话，将编译失败。可以使用 _ 下划线符号作为别名来忽略导入的包，从而避免编译错误，这只会执行 package 的 init() 。
    - 用 := 简短声明的变量只能在函数内部使用。
    - 不能用简短声明方式来单独为一个变量重复声明。
    - 不能使用简短声明来设置字段（结构体的成员变量）的值。
    - nil 是 interface、function、pointer、map、slice 和 channel 类型变量的默认初始值。但声明时不指定类型，编译器也无法推断出变量的具体类型。
    - 允许对值为 nil 的 slice 添加元素，但对值为 nil 的 map 添加元素则会造成运行时 panic
    ```go
    func main(){
        m := make(map[int]int)
        m[1]=2

        var s []int
        s = append(s, 1) // 注意这里要赋值
        
        // 注意区分
        // var s [3]int
        // var s []int
    }
    ```
    - 在创建 map 类型的变量时可以指定容量，但不能像 slice 一样使用 cap() 来检测分配空间的大小
    ```go
    func main(){
        m := make(map[string]int, 99)
        // 非法 println(cap(m))
    }
    ```
    - string 类型的变量值不能为 nil。
    - Go的Array 类型作为函数参数是传值。
    ```Go
    func main(){
        x := [3]int{1,2,3}
        func(arr *[3]int){
            (*arr)[0] = 7
            fmt.Println(arr)
        }(&x)
        fmt.Println(x)

        y := []int{1,2,3}
        func(arr []int){
            arr[0]=7
        }(x)

    }
    ```
    - Go 中的 range 在遍历arr、slice时会生成2个值，第一个是元素索引，第二个是元素的值。
    - slice 和 array 其实是一维数据 需要分步
    ```go
    func main(){
        x := 2
        y := 4

        table := make([][]int, x)
        for i := range  table {
            table[i] = make([]int, y)
        }
    }
    ```
    - 访问 map 中不存在的key
    ```go
    func main(){
        x := map[string]string{"one":"2","two":"","three":"3}
        if _, ok := x["two"]; !ok {
            fmt.Println("key two is no entry")
        }
    }
    ``` 
    - string 类型的值，不可更改
    不允许尝试使用索引遍历字符串，来更新字符串中的个别字符。
    - 一个 UTF8 编码的字符可能会占多个字节，比如汉字就需要 3~4 个字节来存储，此时更新其中的一个字节是错误的。更新字串的正确姿势：将 string 转为 rune slice（此时 1 个 rune 可能占多个 byte），直接更新 rune 中的字符。
    ```Go
    func main(){
        x := "text"
        xRunes := []rune(x)
        xRunes[0]='我'
        x = string(xRunes)
        fmt.Println(x)
    }
    ```
    - 声明语句中 } 折叠到单行后，尾部的 , 不是必需的。
    - Go 的内建函数 len() 返回的是字符串的  byte 数量，而不是像 Python  中那样是计算 Unicode 字符数。如果要得到字符串的字符数，可使用 "unicode/utf8" 包中的 RuneCountInString(str string) (n int) 。RuneCountInString 并不总是返回我们看到的字符数，因为有的字符会占用 2 个 rune。
    - range 迭代 string 得到的值。range 得到的索引是字符值（Unicode point / rune）第一个字节的位置，与其他编程语言不同，这个索引并不直接是字符在字符串中的位置。for range 迭代会尝试将 string 翻译为 UTF8 文本，对任何无效的码点都直接使用 0XFFFD rune（�）UNicode 替代字符来表示。如果 string 中有任何非 UTF8 的数据，应将 string 保存为 byte slice 再进行操作。
    - log.Fatal 和 log.Panic 不只是 log。log 标准库提供了不同的日志记录等级，与其他语言的日志库不同，Go 的 log 包在调用 Fatal*()、Panic*() 时能做更多日志外的事，如中断程序的执行等
    - range 迭代 map不是有序地，想要保证有序性需要其他数据结构保证。
    - switch 语句中的 case 代码块会默认带上 break，但可以使用 fallthrough 来强制执行下一个 case 代码块。
    - Go 特立独行，去掉了前置操作，同时 ++、-- 只作为运算符而非表达式。
    - Go 重用 ^ XOR 操作符来按位取反。同时 ^ 也是按位异或（XOR）操作符。
    - 不导出的 struct 字段无法被 encode。以小写字母开头的字段成员是无法被外部直接访问的，所以 struct 在进行 json、xml、gob 等格式的 encode 操作时，这些私有字段会被忽略，导出时得到零值。
    - 程序默认不等所有 goroutine 都执行完才退出
    - 
    ```go
    // 等待所有 goroutine 执行完毕
    // 使用传址方式为 WaitGroup 变量传参
    // 使用 channel 关闭 goroutine
    func main() {
        var wg sync.WaitGroup
        done := make(chan struct{})
        ch := make(chan interface{})

        workerCount := 2
        for i := 0; i < workerCount; i++ {
            wg.Add(1)
            go doIt(i, ch, done, &wg)	// wg 传指针，doIt() 内部会改变 wg 的值
        }

        for i := 0; i < workerCount; i++ {	// 向 ch 中发送数据，关闭 goroutine
            ch <- i
        }

        close(done) // 当一个通道被关闭后，任何从该通道接收数据的操作（即 <-done）会立即返回零值，而不会阻塞。对于 struct{} 类型的通道，当我们尝试从这个通道接收数据时，如果通道已经被关闭，那么接收操作将立即返回零值，即空的 struct{}。由于 struct{} 没有字段，它的零值是不需要分配内存的，所以接收操作几乎是瞬间完成的，也不会阻塞。
        wg.Wait()
        close(ch)
        fmt.Println("all done!")
    }

    // ch <-chan interface{}: 这个参数是一个接收只读的通道（receive-only channel），其元素类型为interface{}。<-语法表示这个通道只能用于从通道接收数据，不能发送数据。使用接收只读通道作为参数是一种良好的实践，因为它明确地限制了函数只能从通道接收数据，而不能向通道发送数据，有助于提高代码的安全性和可读性。
    func doIt(workerID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
        fmt.Printf("[%v] is running\n", workerID)
        defer wg.Done()
        for {
            select {
            case m := <-ch:
                fmt.Printf("[%v] m => %v\n", workerID, m)
            case <-done:
                fmt.Printf("[%v] is done\n", workerID)
                return
            }
        }
    }

    ```
    - 向已关闭的 channel 发送数据会造成 panic。
    - 在无缓冲的通道上进行通信时，发送和接收操作是同步的，都会阻塞。带缓冲的通道可以在一定程度上解耦发送者和接收者的同步需求，允许一定的异步性。
    - 在一个值为 nil （未初始化）的 channel 上发送和接收数据将永久阻塞。利用这个死锁的特性，可以用在 select 中动态的打开和关闭 case 语句块：
    ```go
    func main() {
        inCh := make(chan int)
        outCh := make(chan int)

        go func() {
            var in <-chan int = inCh
            var out chan<- int
            var val int

            for {
                select {
                case out <- val:
                    println("--------")
                    out = nil
                    in = inCh
                case val = <-in:
                    println("++++++++++")
                    out = outCh
                    in = nil
                }
            }
        }()

        go func() {
            for r := range outCh {
                fmt.Println("Result: ", r)
            }
        }()

        time.Sleep(0)
        inCh <- 1
        inCh <- 2
        time.Sleep(3 * time.Second)
    }

    ```
    - 若函数 receiver 传参是传值方式，则无法修改参数的原有值:
    ```go
    type data struct {
        num   int
        key   *string
        items map[string]bool
    }

    // 指针接收者的方法 (pointerFunc)
    // 这个方法的接收者是指向 data 类型的指针 (*data)。
    // 在方法体内对 this.num 的修改会直接影响到原始的 data 实例，因为 this 指向的是原始实例的地址。
    func (this *data) pointerFunc() {
        this.num = 7
    }

    // 值接收者的方法 (valueFunc)
    // 这个方法的接收者是 data 类型的值。
    // 在方法体内对 this.num 的修改不会影响到原始的 data 实例，因为 this 是原始实例的一个副本。
    // 然而，对于 *this.key 和 this.items 的修改却能影响到原始实例。这是因为：
    // this.key 是一个指针，它指向原始的字符串变量，因此修改指针指向的内容会影响原始的字符串。
    // this.items 是一个映射（map），映射本身是一个引用类型，在Go中传递映射时传递的是映射的引用，而不是整个映射的拷贝。所以，即使 this 是 data 的一个副本，this.items 指向的仍然是原始映射，因此对它的修改会影响到原始映射。
    func (this data) valueFunc() {
        this.num = 8
        *this.key = "valueFunc.key"
        this.items["valueFunc"] = true
    }

    // 在Go语言中，当你定义一个方法时，接收者可以是指针类型或值类型。当调用方法时，Go编译器会自动处理指针和值之间的转换，以确保方法能够被正确调用。
    func main() {
        key := "key1"

        d := data{1, &key, make(map[string]bool)}
        fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)

        d.pointerFunc()	// 修改 num 的值为 7
        fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)

        d.valueFunc()	// 修改 key 和 items 的值
        fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)
    }

    ```
    - Go 是静态类型语言。
    - 在Go语言中，当你将参数作为指针传递给函数时，你不需要显式地解引用指针（即使用*操作符）来访问或修改指针指向的值。这是因为Go语言允许你直接使用点号.操作符来访问结构体字段或调用方法，即使你是通过指针访问这个结构体。然而，如果你需要读取指针所指向的变量的值或者是在非结构体类型的指针上进行操作，那么你还是需要用\*来解引用指针。
    - 在Go语言中，当你定义一个方法时，接收者可以是指针类型或值类型。当调用方法时，Go编译器会自动处理指针和值之间的转换，以确保方法能够被正确调用。
    - WaitGroup变量是传值
- 接口（Interface）与实现
接口是一种抽象类型，它定义了一组方法签名。如果一个类型实现了接口的所有方法，那么这个类型就实现了该接口。Go的接口是隐式实现的，不需要显式声明。
```go
type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    // 这里，s 是一个 Speaker 类型的接口变量，它只知道 Speak() 方法的存在。如果你想访问 Dog 或 Cat 的字段（例如 Name），你需要进行类型断言（type assertion）来将接口变量转换回具体类型。
    var s Speaker
    s = Dog{Name: "Buddy"}
    fmt.Println(s.Speak()) // 输出 "Woof!"

    s = Cat{Name: "Whiskers"}
    fmt.Println(s.Speak()) // 输出 "Meow!"

    // 安全形式的类型断言，以避免潜在的运行时错误。
    if dog, ok := s.(Dog); ok {
        fmt.Println(dog.Name)
    }
    // 直接断言类型，如果失败会导致运行时panic。
    // dog := s.(Dog) // 如果 s 不是 Dog 类型，则会 panic
}
```

## 在 Go 语言中，`panic` 和 `return` 是处理错误和异常情况的两种不同机制，它们各自有特定的应用场景和效果。

### `panic`

**定义与行为**

- **`panic`** 是一个内置函数，用于表示程序遇到了无法恢复的错误状态。当调用 `panic` 时，它会立即停止当前函数的执行，并开始回退当前 goroutine 的栈，运行任何已定义的延迟函数（通过 `defer` 关键字）。如果未被 `recover` 捕获，这种回退会一直持续到该 goroutine 的初始调用处，最终导致程序崩溃。
  
**使用场景**

- **不可恢复的错误**：当你遇到不应该发生的严重逻辑错误时使用 `panic`。例如，数组越界、空指针引用等。
- **初始化失败**：如果某个包级别的初始化过程失败且无法继续正常操作，则可以使用 `panic`。
- **断言检查**：有时开发者会在代码中添加断言来验证某些不应发生的情况，一旦这些条件不满足就触发 `panic`。

**示例**
```go
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

在这个例子中，如果尝试除以零，程序将直接进入恐慌状态并终止执行。

### `return`

**定义与行为**

- **`return`** 用于从函数返回，它可以带回一个或多个结果值给调用者。不同于 `panic`，使用 `return` 不会导致程序崩溃，而是按照正常的控制流结束当前函数，并允许调用者根据返回的结果进行进一步处理。
  
**使用场景**

- **可恢复的错误**：大多数情况下，你应该使用 `return` 来返回错误，让调用者决定如何处理这些错误。这种方式更灵活，允许你采取不同的补救措施，如重试、记录日志或向用户显示友好的错误信息。
- **正常流程控制**：即使没有错误发生，`return` 也用于正常退出函数，并返回计算结果或其他相关信息。

**示例**
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

这里，我们不是直接引发恐慌，而是返回一个错误给调用者，让调用者决定下一步做什么。

### 总结

- **`panic`** 更适合用于那些真正意义上的“致命”错误，即一旦发生，程序无法继续正常运行的情况。使用 `panic` 应该非常谨慎，因为它可能导致程序意外终止。
- **`return`** 则是处理常规错误的标准方式，尤其是那些可以在应用层面上被合理处理的错误。这使得你的应用程序更加健壮和可靠，同时也提高了代码的可读性和维护性。

选择哪种方法主要取决于错误的性质以及你希望程序如何响应这些错误。通常来说，优先考虑使用 `return` 来处理错误，仅在必要时使用 `panic`。此外，在一些特殊情况下，结合 `recover` 可以捕获 `panic` 并提供某种程度的错误恢复能力，但这应该作为最后的手段。

## 多个 defer 语句的执行顺序是后进先出(LIFO, Last In First Out)，即最后一个被 defer 的函数会最先执行。