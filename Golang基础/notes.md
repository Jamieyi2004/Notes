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
## serInfo{}：这是一个复合字面量（Composite Literal），用于创建一个 UserInfo 类型的实例。复合字面量允许你以一种紧凑的方式初始化数据结构（如结构体、数组、切片和映射）。当你使用 {} 并且不传递任何初始值时，所有字段都将被赋予其类型的零值；&：这是取地址操作符，在这里用来获取新创建的 UserInfo 实例的地址。这意味着你得到的是一个指向该实例的指针而不是实例本身。这对于需要修改原对象或避免复制大型结构体的情况特别有用。

## map[string]interface{} 是 Go 语言中的一种数据结构，它表示一个键为字符串、值为任意类型的映射（即字典）。这里的 interface{} 是 Go 的空接口类型，它可以存储任何类型的值。因此，map[string]interface{} 可以用来创建一个动态的键值对集合，其中每个键都是字符串，而对应的值可以是任意类型的数据。
## 在 Go 中，`模块`是一组相关的包（packages），这些包一起工作以提供一组功能。从 Go 1.11 开始引入了 `go mod` 工具来支持模块，允许开发者更灵活地管理项目依赖。每个模块都有一个唯一的模块路径（通常是托管该模块的 Git 仓库 URL），并且有一个 `go.mod` 文件来记录模块的元数据及其依赖关系。使用模块可以帮助解决 Go 项目的依赖管理和版本控制问题。
## go mod init go-dev` 是 Go 语言中用于初始化一个新的模块的命令。具体来说：

    - `go mod`：是 Go 1.11 引入的模块支持工具，用来管理项目依赖。它替代了旧版本的 Go 中使用的 `GOPATH` 工作空间和 `vendor` 文件夹来处理依赖关系。
    - `init`：是 `go mod` 的子命令之一，用于初始化一个新的 Go 模块，在当前目录下创建一个 `go.mod` 文件。
    - `go-dev`：这是你为你的 Go 模块指定的模块路径名称。它可以是你项目的唯一标识符（例如，基于托管代码的Git仓库URL），也可以是一个描述性的名称，取决于你如何组织和分发你的代码。

    执行 `go mod init go-dev` 命令后，Go 会在当前目录创建一个名为 `go.mod` 的文件，该文件会包含类似如下的内容：

    ```
    module go-dev

    go 1.x
    ```

    这行声明表示这个 Go 模块的名称是 `go-dev`，并且指定了模块所兼容的最小 Go 版本（这里的 `1.x` 会根据你的Go版本自动填充）。

    `go.mod` 文件对于追踪项目直接依赖的其他模块及其版本非常重要。当你添加新的依赖时，这些信息会被记录在这个文件中，并且 Go 工具链会使用这个文件来确保所有开发者以及构建环境都使用相同的依赖版本。

    如果你打算分享或发布你的模块，推荐使用更具体的模块路径，比如包含你在代码托管平台上的用户名和仓库名，例如 `github.com/yourusername/go-dev`。这样可以确保模块路径的唯一性，并使得其他人更容易导入和使用你的模块。

## Go 编译器通过 go.mod 文件来确定代码所在的模块。
当在一个项目中运行 Go 命令（如 go build、go run 或 go test）时，编译器会从当前工作目录开始查找 go.mod 文件，沿着文件系统的父目录逐级向上搜索，直到找到一个 go.mod 文件为止。这个过程被称为“模块根”的发现。一旦找到了 go.mod 文件，该文件所在目录就被认为是模块的根目录，并且所有位于此目录及其子目录下的 Go 包都被认为是属于同一个模块。  

## `go.mod `文件定义了 Go 模块的`元数据和依赖项`，而 `go.sum `文件记录了`模块依赖项的版本及其校验和`，确保依赖的完整性与可重复构建。`go mod tidy `是 Go 模块系统中的一个命令，用于管理和清理 go.mod 和 go.sum 文件。

## 在 Go 语言中，`包（package）和文件夹（directory`之间的关系通常是`1:1的映射`。
这意味着一个包通常存在于一个单独的文件夹下，该文件夹下的所有源代码文件都属于同一个包。这种组织方式有助于保持代码的清晰性和可维护性，并且是 Go 社区广泛接受的标准。每个 Go 源文件的第一行应该是 package 声明，用来指定该文件所属的包。

## 在 Go 语言中，只有` main 包`可以包含一个名为 main 的函数，并且这个 main 函数是程序的入口点。其他包不可以定义自己的 main 函数。如果尝试在一个非 main 包中定义 main 函数，编译器会报错，因为这违反了 Go 语言的规定。

## `go build` 更适合于生产环境部署和构建过程，你会得到一个可执行文件，你可以随后运行它；而 `go run` 则更适合于开发和测试阶段，你会编译并立即运行这个程序，你会看到输出但不会有持久的二进制文件留下。

## Go 编译器会递归地编译所有依赖的包。这意味着不仅 hello.go 会被编译，而且它所导入的所有包（包括你自己写的包）以及这些包所依赖的任何其他包也会被编译。这个过程确保了整个依赖链都被正确处理，并且生成的可执行文件包含了运行所需的所有代码。

## `go get` :
如果你的项目已经启用了 Go Modules（通过在项目根目录下创建了 go.mod 文件），那么 go get -u github.com/gin-gonic/gin 的行为如下：
1. 下载到本地模块缓存：首先，Go 会将 gin 包下载到本地模块缓存中，默认路径通常是 $GOPATH/pkg/mod/ 或者如果设置了 GOMODCACHE 环境变量，则会下载到该路径指定的位置。
2. 更新 go.mod 和 go.sum 文件：


  // 定义一个字符串到整数的映射
    // map[string]int{} 更加直观，适用于初始化空 map 或者带有一些初始值的 map。
    // make(map[string]int) 更灵活，特别是在你需要预分配一定容量时非常有用。
    在Go语言中，定义并初始化一个slice有几种不同的写法。以下是常见的几种方式：

1. 使用字面量直接定义并初始化：
   ```go
   s := []int{1, 2, 3}
   ```
   这种方式直接定义了一个类型为`[]int`的slice，并初始化了它的元素。

2. 使用内置的`make`函数：
   ```go
   s := make([]int, 5) // 定义了一个长度为5，初始值为0的slice
   ```
   或者指定容量（capacity）：
   ```go
   s := make([]int, 3, 5) // 定义了一个长度为3，容量为5的slice
   ```

3. 从现有的数组或slice创建新的slice：
   ```go
   arr := [5]int{1, 2, 3, 4, 5}
   s := arr[1:4] // 创建一个新的slice，包含arr的第2个到第4个元素（索引从0开始）
   ```

4. 使用`nil`初始化slice变量，然后在后续代码中动态添加元素：
   ```go
   var s []int // 初始化为nil slice
   s = append(s, 1) // 动态添加元素
   ```

这些方法分别适用于不同的场景，根据实际需要选择合适的方式来定义和初始化slice。


// 在Gin框架中，c.JSON方法用于向客户端返回JSON格式的响应。这个方法通常有两个参数：
                // 状态码（int）：表示HTTP响应的状态码，比如200表示成功，400表示客户端错误，500表示服务器内部错误等。
                // 数据（interface{}）：需要序列化为JSON格式并发送给客户端的数据，它可以是Go语言中的任何类型，如map、slice、struct等。

// Find vs First: Find用于查询多条记录并将它们存储在一个切片中，而First则用于获取第一条匹配的记录并填充到提供的结构体实例中。因此，Find可以接受一个未初始化的切片指针，因为它会自动管理切片的大小和容量。相比之下，First需要一个具体的对象来填充数据，这意味着你需要提前为这个对象分配内存。

## string
### 遍历
在Go语言中，处理字符串时经常会遇到两种不同的循环方式来遍历字符串中的字符。这两种方式主要区别在于它们如何处理和解释字符串中的字符编码，特别是涉及到非ASCII字符（比如中文）的时候。

#### 通过`len`函数和索引访问

```go
s := "你好 golang"
for i := 0; i < len(s); i++ {
    fmt.Printf("%v(%c)\t", s[i], s[i])
}
```

这段代码尝试通过遍历字符串的每个字节并打印出来。这里使用了`len(s)`来获取字符串的字节长度，并直接用索引访问字符串中的每个字节。对于纯ASCII字符（如空格和"golang"部分），这种方式工作正常。但是，对于像"你好"这样的UTF-8编码的非ASCII字符，这种方法会导致问题。这是因为UTF-8编码使用一个以上的字节来表示一个字符。例如，在UTF-8中，一个中文字符通常由3个字节组成。当你试图将这些字节单独作为字符打印时（使用`%c`格式化符），你不会得到正确的字符，而是得到其组成部分的不可读字节值。

#### 使用`range`进行遍历

```go
for index, v := range s {
    fmt.Println(index, v)
}
```

使用`range`遍历字符串是一种更推荐的方式，特别是当字符串包含非ASCII字符时。`range`能够正确识别UTF-8编码，并为每个字符返回其对应的rune值（即Unicode码点）。这里的`index`是字符（而不是字节）在字符串中的起始位置，`v`是该字符的Unicode码点。这样，无论字符是ASCII还是任何其他Unicode字符，都能被正确地识别和处理。这种方式适用于所有支持的字符集，包括但不限于中文、日文、韩文等多字节字符集。

总结来说，如果你需要逐字符处理字符串，并且字符串可能包含非ASCII字符（如中文），你应该使用`range`来遍历字符串。这确保了每个字符都被正确解析和处理，避免了由于UTF-8编码带来的潜在问题。而直接通过索引访问字符串并使用`len`函数来控制循环，则更适合于仅包含ASCII字符或你需要特别操作底层字节的情况。

### 字符数
如果你想获取字符串 `s` 中字符的数量（即字符的个数，而不是字节数），你应该使用能正确处理UTF-8编码的方法。在Go语言中，可以通过将字符串转换为`rune`类型来实现这一点，因为`rune`代表单个Unicode字符，并且可以正确处理多字节字符如中文、日文等。

下面是一个示例代码，演示了如何计算字符串中的字符数：

```go
s := "你好 golang"
// 将字符串转换为rune切片
runeSlice := []rune(s)
// 获取字符数
charCount := len(runeSlice)
fmt.Println("字符数:", charCount)
```

在这个例子中，我们首先通过`[]rune(s)`将字符串`s`转换为一个`rune`类型的切片。这样做的好处是它可以正确处理包含多字节字符的字符串。然后，我们使用`len()`函数对这个`rune`切片求长度，得到的结果就是字符串中字符的实际数量，无论这些字符是否由多个字节组成。

这种方法适用于任何包含Unicode字符的字符串，确保你能够准确地统计出字符串中的字符数量。

### Unicode
Unicode是一种字符编码标准，旨在为全球所有的书写系统中的每一个字符提供一个唯一的数字编号，不论平台、程序或语言。它的设计目标是支持世界上所有现代和古代的书面语言，并且能够随着新字符的需要而扩展。

以下是关于Unicode的一些关键点：

1. **唯一标识符**：每个字符在Unicode标准中都有一个唯一的标识符，称为代码点（Code Point）。例如，字母“A”的Unicode代码点是U+0041。

2. **统一性**：Unicode尝试将所有字符集统一到一个单一的标准中，从而解决不同字符编码之间的兼容性和转换问题。这意味着无论你处理的是拉丁文、汉字、阿拉伯文还是其他任何文字，都可以使用同一套编码规则。

3. **广泛覆盖**：Unicode涵盖了几乎所有现代使用的书写系统以及许多历史上的书写系统。它不仅包括了文本字符，还包括了标点符号、技术符号、数学符号、甚至表情符号等。

4. **编码形式**：虽然Unicode给每个字符分配了一个理论上的唯一数字，但这些数字需要以某种方式存储在计算机上。因此，Unicode定义了几种不同的编码形式，包括UTF-8、UTF-16和UTF-32。其中：
   - **UTF-8**：一种变长编码，可以使用1到4个字节表示一个字符。它向后兼容ASCII。
   - **UTF-16**：使用16位编码单元，对于一些字符可能需要两个编码单元（即所谓的代理对）。
   - **UTF-32**：固定长度编码，总是使用32位来表示一个字符。

5. **跨平台和跨语言支持**：由于其全面性和灵活性，Unicode成为了互联网、操作系统和编程语言中最常用的字符编码之一，确保了数据可以在不同的环境之间无缝传输和处理。

通过采用Unicode，开发者可以创建更加国际化和本地化的软件产品，同时简化了多语言文本处理的任务。

### 修改字符串

要修改字符串，需要先将其转换成[]rune 或 []byte类型，完成后在转换成string，无论哪种转换都会重新分配内存，并复制字节数组

转换为 []byte 类型

```go
// 字符串转换
s1 := "big"
byteS1 := []byte(s1)
byteS1[0] = 'p'
fmt.Println(string(byteS1))
```

转换为rune类型

```go
// rune类型
s2 := "你好golang"
byteS2 := []rune(s2)
byteS2[0] = '我'
fmt.Println(string(byteS2))
```