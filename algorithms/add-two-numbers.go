package main

import (
	"strconv"
	"fmt"
)

func main() {

	n1 := &ListNode{
		Val:0,
		Next:&ListNode{
			Val:0,
			Next:&ListNode{
				Val:0,
			},
		},
	}

	n2 := &ListNode{
		Val:0,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:1,
			},
		},
	}

	result := addTwoNumbers(n1, n2)

	for  e := result; e != nil; e = e.Next {
		fmt.Println(e)
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ns1 []int
	for n := l1; n != nil; n = n.Next {
		ns1 = append(ns1, n.Val)
	}
	num1 := buildnum(ns1)

	var ns2 []int
	for n := l2; n != nil; n = n.Next {
		ns2 = append(ns2, n.Val)
	}
	num2 := buildnum(ns2)

	sum := num1 + num2

	return buildslice(sum)
}

func buildslice(num int) *ListNode {
	head := &ListNode{}
	list := head
	ns := strconv.Itoa(num)
	for i := len(ns)-1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(ns[i]))
		list.Val = n
		if i > 0 {
			list.Next = &ListNode{}
			list = list.Next
		}
	}
	return head
}

func buildnum(numslice []int) int {
	var num int
	for i := len(numslice)-1; i >=0; i-- {
		if numslice[i] == 0 {
			num *= 10
			continue
		}

		num *= 10
		num += numslice[i]
	}

	return num
}
