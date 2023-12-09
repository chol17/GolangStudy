package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 打印1~100 之间所有是9 的整数的个数及总和
	sum := 0
	num := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			sum += i
			num++
		}
	}
	fmt.Printf("整除的个数为：%v\n", num)
	fmt.Printf("整除的总和为：%v\n", sum)

	//随机生成1~100的一个数，直到生成了99这个数，看看一共用了多少次
	// 我们为了生存一个随机数，还需要给 rand 设置一个种子
	// time.Now().Unix() 表示从1970-01-01 00:00:00 到现在的时间戳值
	rand.Seed(time.Now().Unix()) // 因为时间戳一直在变。因此根据时间戳生成的随机种子也就一直在变，从而生成的随机数根据时间戳绑定，就一直可以变化，这样就达到了获取随机数的效果
	//rand.Seed(time.Now().UnixNano()) // 纳秒
	rand_num := rand.Intn(100) + 1
	count := 0
	for {
		if rand_num == 99 {
			fmt.Printf("一共用了:%d", count)
			break
		} else {
			count++
			rand_num = rand.Intn(100) + 1
		}
	}
}
