//链表

package main

type SingalNode struct {
	Next  *SingalNode
	Value int
	Rand  *SingalNode //另一个指针，指向一个随机的节点
}

//找到链表的中点
func FindMidNode(head *SingalNode) *SingalNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func ReversalNode(head *SingalNode) *SingalNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	last := head.Next
	first.Next = nil
	for last != nil {
		temp := first
		first = last
		last = last.Next
		first.Next = temp
	}
	return first
}

//判断一个链表是否是回文结构
//例如 1->2->3->2->1 是回文
//1->2->2->1 是回文
//1->20->3->20->1 是回文
//解法1（需要栈）：
//链表右半部分入栈，
//出栈和头指针比较，如果相等，指针右移，继续出栈比较
//有不想等，那么不是回文，
//如果栈无数据，是回文
//快慢指针，找出链表中点

//解法2
//将后半部分指针逆序
//从头和尾向中间前进
func IsSymmetryLinkedList(head *SingalNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := FindMidNode(head)

	first := head
	last := ReversalNode(mid)
	temp := last
	result := true
	for temp != nil {
		if first.Value != temp.Value {
			result = false
			break
		}
		temp = temp.Next
		first = first.Next
	}

	ReversalNode(last)

	return result
}

//单链表，小于某个值在左边，等于某个值在中间，大于某个值在右边
//方法一：链表节点放在数组中，进行分区操作（空间复杂度较大）
//方法二：使用三对头尾指针，小于a，等于a，大于a；遍历指针之后，再将三对头尾指针相连。（此方法可以保证相对位置不变）
func ListPartition(head *SingalNode, a int) *SingalNode {
	var minHead *SingalNode
	var minTail *SingalNode
	var eqHead *SingalNode
	var eqTail *SingalNode
	var maxHead *SingalNode
	var maxTail *SingalNode

	for head != nil {
		next := head.Next
		if head.Value < a {
			if minHead == nil {
				minHead = head
				minTail = head
			} else {
				minTail.Next = head
				minTail = head
			}
		} else if head.Value == a {
			if eqHead == nil {
				eqHead = head
				eqTail = head
			} else {
				eqTail.Next = head
				eqTail = head
			}
		} else {
			if maxHead == nil {
				maxHead = head
				maxTail = head
			} else {
				maxTail.Next = head
				maxTail = head
			}
		}
		head = next
	}
	//如果有小于区域
	if minTail != nil {

		head = minHead
		if eqHead != nil {
			minTail.Next = eqHead
			if maxHead != nil {
				eqTail.Next = maxHead
			}
		} else if maxHead != nil {
			minTail.Next = maxHead
		}
	} else if eqHead != nil {
		head = eqHead
		if maxHead != nil {
			eqTail.Next = maxHead
		}
	} else if maxHead != nil {
		head = maxHead
	}
	return head
}

//复制带有random节点的链表
//方法一：copy所有节点，通过map，存储对应关系（key，value）。根据key节点的指向，将value节点执行对应的克隆节点
//方法二：将链表直接复制，a->b->c => a->a'->b->b'->c->c'; 再遍历，将随机节点指向，在x'上分配，最后返回x'
func CopyListWithRandomNode(head *SingalNode) *SingalNode {
	temp := head

	if head == nil {
		return nil
	}

	//1. a->b->c => a->a'->b->b'->c->c'
	for temp != nil {
		newNode := CopyNodeValue(temp)
		next := temp.Next
		temp.Next = newNode
		newNode.Next = next
		temp = next
	}
	//再copy随机节点
	temp = head
	for temp != nil {
		next := temp.Next
		if temp.Rand != nil {
			next.Rand = temp.Rand.Next
		} else {
			next.Rand = nil
		}
		temp = next.Next
	}

	tailOld := head
	res := head.Next
	tail := res
	temp = res.Next
	for temp != nil {
		tailOld.Next = temp
		tailOld = temp
		tail.Next = temp.Next
		tail = temp.Next
		temp = tail.Next
	}
	tailOld.Next = nil
	return res
}

//仅仅拷贝值
func CopyNodeValue(node *SingalNode) *SingalNode {
	return &SingalNode{Value: node.Value}
}

func CompareNodeList(n1, n2 *SingalNode) bool {
	temp1 := n1
	temp2 := n2
	for temp1 != nil && temp2 != nil {
		if temp1.Value != temp2.Value {
			return false
		}
		if temp1.Next == nil && temp2.Next != nil {
			return false
		}
		if temp1.Next != nil && temp2.Next == nil {
			return false
		}
		if temp1.Next != nil && temp2.Next != nil {
			if temp1.Next.Value != temp2.Next.Value {
				return false
			}
		}
		if temp1.Rand == nil && temp2.Rand != nil {
			return false
		}
		if temp1.Rand != nil && temp2.Rand == nil {
			return false
		}
		if temp1.Rand != nil && temp2.Rand != nil {
			if temp1.Rand.Value != temp2.Rand.Value {
				return false
			}
		}
		temp1 = temp1.Next
		temp2 = temp2.Next
	}
	if temp1 != nil || temp2 != nil {
		return false
	}
	return true
}

//【题目】给定两个可能有环，也可能无环的单链表，头节点head1，head2，请实现一个函数，如果两个链表相交，请返回相交的第一个节点，如果不想交，请返回null。
//【要求】如果两个链表的长度之和是N，时间复杂度是O（N），额外空间复杂度达到O（1）
func GetIntersectNode(node1, node2 *SingalNode) *SingalNode {
	if node1 == nil || node2 == nil {
		return nil
	}
	c1 := GetCircleListNode(node1)
	c2 := GetCircleListNode(node2)

	if c1 == nil && c2 == nil {
		return noCircle(node1, node2)
	} else if c1 != nil && c2 != nil {
		return bothCircle(node1, node2, c1, c2)
	}
	return nil
}

//两个链表没有环
func noCircle(node1, node2 *SingalNode) *SingalNode {
	if node1 == nil || node2 == nil {
		return nil
	}
	//相交的话，只有一种可能，相交后的所有节点都相同，如果相交，最后一个节点是相同的
	temp1 := node1
	temp2 := node2
	len1 := 1
	len2 := 1
	for temp1.Next != nil {
		temp1 = temp1.Next
		len1++
	}
	for temp2.Next != nil {
		temp2 = temp2.Next
		len2++
	}

	//最后一个节点不同，表示不相交
	if temp1 != temp2 {
		return nil
	}
	temp1 = node1
	temp2 = node2
	if len1 > len2 {
		for len1 != len2 {
			temp1 = temp1.Next
			len1--
		}
	} else if len1 < len2 {
		for len1 != len2 {
			temp2 = temp2.Next
			len2--
		}
	}
	for temp1 != temp2 {
		temp1 = temp1.Next
		temp2 = temp2.Next
	}
	return temp1
}

//两个链表都有环
func bothCircle(node1, node2, circle1, circle2 *SingalNode) *SingalNode {
	//链表的入环点相同, 可以将入环点作为末尾点，从而转换成无环的问题
	if circle1 == circle2 {
		temp1 := circle1.Next
		temp2 := circle2.Next
		circle1.Next = nil
		circle2.Next = nil
		res := noCircle(node1, node2)
		circle1.Next = temp1
		circle2.Next = temp2
		return res
	} else { //一定在环上有两个交点
		temp := circle1.Next
		for temp != circle1 {
			if temp == circle2 {
				return temp
			}
			temp = temp.Next
		}
		return nil
	}
}

//判断链表有环无环，返回第一个入环节点，否则返回nil
//使用快慢指针，空间复杂度最小
//1.slow指针和fast指针，如果有环必相遇
//2.相遇后快fast指针回到头，slow指针不动，同时以速度1前进，一定会在入环点相遇
//大体的证明过程：https://blog.csdn.net/donghuaan/article/details/78988987
//头到入环口距离k，此时快指针走了2k，其中后一个k走了n圈+delta ，所以 k = nR + delta
//因为快指针距离slow指针式delta，所以在环上相遇，要多走 R-delta，所以快指针走2*（R-delta），所以相遇位置事2*（R-delta）+ delta = 2R-delta
//k = nR + delta, k+delta = (2+n)R 所以在入口相遇
func GetCircleListNode(node *SingalNode) *SingalNode {
	head := node
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast := head.Next.Next
	slow := head.Next
	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
