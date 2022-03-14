1.多核CPU场景下，cache如何保持一致、不冲突？

处理器上有一套完整的协议，来保证Cache一致性。比较经典的Cache一致性协议当属MESI协议
M(Modified) :这行数据有效，数据被修改了，和内存中的数据不一致，数据只存在于本Cache中。
E(Exclusive):这行数据有效，数据和内存中的数据一致，数据只存在于本Cache中。
S(Shared):这行数据有效，数据和内存中的数据一致，数据存在于很多Cache中。
I(Invalid):这行数据无效。

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
