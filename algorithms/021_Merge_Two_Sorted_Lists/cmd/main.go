package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoSortedLists(list1, list2 *ListNode) *ListNode {

	mergeList := &ListNode{}
	temp := mergeList

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}

		temp = temp.Next
	}

	if list1 == nil {
		temp.Next = list2
	}
	if list2 == nil {
		temp.Next = list1
	}

	return mergeList.Next
}
