//链表

package main

type SingalNode struct {
	Next  *SingalNode
	Value int
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
