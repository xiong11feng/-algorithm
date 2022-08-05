package main

//异或的性质:
// A^A = 0
// 0^A = A
// A^B = B^A
// A^B^C = A^(B^C)
// 此算法就是根据异或的性质，将数组中i，j元素互换位置，
//  i != j
func Xor_ExchangeItem(items []int, i, j int) {
	items[i] = items[i] ^ items[j]
	items[j] = items[i] ^ items[j]
	items[i] = items[i] ^ items[j]
}

//题目：int数组中，只有一个数字是奇数个，其余的所有数字都是偶数个，找出这个数字
//原因：A^A = 0
//     A^A^A = 0^A =A
//结论：奇数个A异或为A，偶数个A异或为0
func Xor_FindOddDights(items []int) int {
	xor := 0
	for i := 0; i < len(items); i++ {
		xor ^= int(items[i])
	}
	return int(xor)
}

//题目：int数组中，只有两个数字是奇数个，其余的所有数字都是偶数个，找出这个数字
//解题步骤：
//1.数组中所有数异或之后，得到是那两个奇数A和B的异或值 A^B = Z
//2.A^B一定不为0，否则，A=B。所以一定某一个bit上，A B在这一bit上不同（一个是0一个是1）
//3.假设最后1个bit，A是0，B是1.
//4.用所以根据最后1bit是0还是1，可以将数据分成两组。X，Y，Y组最后一位是1。
//5.Y组中包含的数据一定有B，且不包含A。
//6.Y组相当于 Xor_FindOddDights 的问题，只有B是奇数个，其余均是偶数个
//7.Y组中所有数异或，即可得到B，之后，通过 Z^B = A
func Xor_FindTwoOddDights(items []int) (a, b int) {
	z := int(0)
	for i := 0; i < len(items); i++ {
		z ^= items[i]
	}
	//z = a^b

	//找到z的最右的1
	//例如 z  = 1100110000
	// ^z = 0011001111 (z取返)
	//^z+1= 0011010000
	//1100110000 & 0011010000 = 0000010000
	rigth_1 := z & (^z + 1)

	for i := 0; i < len(items); i++ {
		if items[i]&rigth_1 == 0 {
			a ^= items[i]
		}
	}

	b = a ^ z

	return
}
