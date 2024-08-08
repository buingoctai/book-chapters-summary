## Discover `countdown logic` & learn the relevant knownledges
### Logic:



### Concepts:
- goroutines

+ each concurrently executing activity
+ distinct approach compared to traditional thread based models
+ lightweight, manged by go runtime

- channel
+ used to pass synchronized data between goroutines

![Untitled](/docs/goroutines.jpg)

=> `it helps understand how go runtime manage goruntines effectly [internal]`

patterns
+ parallel execution with gorountines
```
func main() {
    ch := make(chan int)

    go compute(ch, 1)
    go compute(ch, 2)

    result1 := <-ch
    result2 := <-ch

    fmt.Println("Result 1:", result1)
    fmt.Println("Result 2:", result2)
}

func compute(ch chan int, num int) {
    result := num * 2
    ch <- result
}
```
+ coordination and synchronization with channels
```
func main() {
    ch := make(chan string)

    go producer(ch)
    go consumer(ch)

    time.Sleep(time.Second)
}

func producer(ch chan string) {
    ch <- "Data"
    close(ch)
}

func consumer(ch chan string) {
    for data := range ch {
        fmt.Println("Received:", data)
    }
}
```
+ select statement for multichannel communication
```
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(time.Second)
        ch1 <- "Hello"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "World"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}

```

=> `it helps know how to receive responses from goroutines via channels, enable adapt with popular use cases`

- common concurrency challenges
+ deadlocks
+ channel misuse
+ race conditions: occur when multiple goroutines access shared data concurrently without proper synchronization => use sync package.
+ resource leaks => use defer package.
+ complex concurrency patterns.


- real world applications
+ web servers and microservices
+ distributed systems
+ database access
+ concurrency patterns in data processing
+ high performance networking
+multimedia and real time applications
+ fintech and trading platforms




### Workflow:
