package main

//使用位运算，计算加，减，乘，除
//a+b 溢出，结果也是溢出
func BitAdd(a, b int) int {
	//temp := a ^ b           //亦或-也是无进位加法
	//addtion := (a & b) << 1 //与操作，计算出的1是进位的位置,左移1位，得到进位的值

	//temp + addtion其实就是最终的a+b
	//重复上述操作，知道addtion，是0，那么结果就是temp

	temp := a ^ b
	addtion := (a & b) << 1
	for addtion != 0 {
		tempA := temp
		temp = temp ^ addtion
		addtion = (tempA & addtion) << 1
	}
	return temp
}

//减法， a-b = a+ 相反数（b）
func BitMinus(a, b int) int {
	return BitAdd(a, BitNegNumber(b))
}

//求一个数的相反数
//当前数取反，加1
//计算机中没有减法，负数表示成，反码+1的形式（补码）
func BitNegNumber(a int) int {
	return ^a + 1
}

//模拟二进制乘法
func BitMulti(a, b int) int {
	aIsNeg := isNeg(a)
	if aIsNeg {
		a = -a
	}
	bIsNeg := isNeg(b)
	if bIsNeg {
		b = -b
	}
	result := 0
	for b != 0 {
		if (b & 1) != 0 { //如果是0，得到的结果是0，跳过
			result = BitAdd(result, a)
		}
		a = a << 1 //a 左移
		b = b >> 1 //b 右移
	}
	if aIsNeg && !bIsNeg {
		return -result
	}
	if !aIsNeg && bIsNeg {
		return -result
	}
	return result

}

func BitDiv(a, b int) int {
	aIsNeg := isNeg(a)
	if aIsNeg {
		a = -a
	}
	bIsNeg := isNeg(b)
	if bIsNeg {
		b = -b
	}
	res := 0
	//a/b  2^31右移动，得到刚好比b大的数（）
	for i := 31; i >= 0; i = BitMinus(i, 1) {
		if a>>i >= b {
			res |= (1 << i)
			a = BitMinus(a, b<<i)
		}
	}
	if aIsNeg && !bIsNeg {
		return -res
	}
	if !aIsNeg && bIsNeg {
		return -res
	}
	return res
}

//判断a是不是负数
func isNeg(a int) bool {
	return a < 0
}
