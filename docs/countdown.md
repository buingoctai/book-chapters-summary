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
+ coordination and synchronization with channels
+ select statement for multichannel communication

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
