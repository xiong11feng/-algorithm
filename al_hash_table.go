//哈希表和哈希函数
//哈希函数，输入无穷，输出有限（2N次方-1）.可能会碰撞
//输出的分布是概率大体相同
//相同输入，输出也相同
package main

//【不隆过滤器】
// url黑名单，假设url有几十亿
// 爬虫去重
// hashset，会导致内存过高
//不隆过滤器，使用内存很少，允许一定程度的失误率
//位图 [10]int64,可以代表表示64*10bit的数组
//178位信息 如何获取？
//178/64 ,找到应该在第几个数
//178%64, 找到第几位
//不隆过滤器，使用位图实现
//m位的位图，每个黑名单url，使用k种hash函数，值对m求摸，找到对应的位图，变成1，理论会产生k个1
//判断url在不在黑名单中时，对url进行k次hash函数，找到对应位置，如果全是1，则认为是黑名单
//此算法不会错过一个黑名单，但是可能会把某些url错误的当成黑名单
//https://github.com/bits-and-blooms/bloom golang的实现

//【一致性哈希】
//https://ably.com/blog/implementing-efficient-consistent-hashing
//https://github.com/kkdai/consistent golang 实现
//虚拟节点，提高节点环上的均一性
