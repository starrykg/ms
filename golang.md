1.
package main

import "fmt"

func main(){
defer_call()
}

func defer_call(){
defer func(){fmt.Println("打印前")}()
defer func(){fmt.Println("打印中")}()
defer func(){fmt.Println("打印后")}()

panic("触发异常")
}
/*
打印后
打印中
打印前
panic: 触发异常
 */

2.
package main

import "fmt"

func main(){
	var a uint = 1
	var b uint = 2
	fmt.Println(a-b)
}

/*
32 位系统 2^32 -1 
64 位系统 2^64 -1
 */
 
3.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//rune:int32的别名，几乎在所有方面等同于int32
	var str = "hello 你好"
	//golang中string底层是通过byte数组实现的，直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))
	//以下两种都可以得到str的字符串长度
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))
}

4.channel
概念:无缓冲的channel，必须先有一个消费者，才能往channel发数据。无缓冲的channel没有缓冲区，没有channel消费，发送者处于阻塞状态，导致程序死锁。
实现:
数据：
向一个关闭的channel里，读数据。
1. 带缓存的channel
缓存里还有数据
返回 数据, true

缓存里没有数据
返回 数据类型未初始化零值, false

2. 不带缓存的channel
缓存里没有数据
返回 数据类型未初始化零值, false

缓存里有数据
执行不到 从缓存里读数据就报错。因为数据没人读就阻塞在哪里了

5.mutex
悲观锁，读写都互斥，同一时间只有一个groutine能拿到
go-mutex的模式：
1. 正常模式：如果 waiter 获取不到锁的时间超过阈值 1 毫秒，那么，这个 Mutex 就进入到了饥饿模式
2. 饥饿模式：如果拥有 Mutex 的 waiter 发现下面两种情况的其中之一，它就会把这个 Mutex 转换成正常模式:
       此 waiter 已经是队列中的最后一个 waiter 了，没有其它的等待锁的 goroutine 了；
       此 waiter 的等待时间小于 1 毫秒。
rwmutex：读操作不产生互斥

6.gpm模型
grouting
processor：中间的一个上下问调度，每一个p有本地的grouting队列，优先分配短的grouting，还会有全局的grouting，全局也满了就会进入抢占模式
machine

7.垃圾回收
三色标记法：
标记阶段，可能有stop the word，所有的grouting处于暂停状态，gc的grouting运行，把没有被引用的要被回收的打上标记。
然后程序继续运行，垃圾回收线程会被清理掉。（go的小优化把大量的片段申请分成一小段一小段，缩短每次的stop the word时间段）
回收机制：手动掉gc函数

8.内存分配
依赖Tcmalloc，利用不同的大小采用不同的分配器
channel内存分配在堆上

9.内存溢出：忘记关闭channel
内存逃逸：在方法内把局部变量指针返回 局部变量原本应该在栈中分配，在栈中回收。但是由于返回时被外部引用，因此其生命周期大于栈。


       
