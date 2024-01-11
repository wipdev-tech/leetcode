package addtwonumbers

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return numberToList(listToNumber(l1) + listToNumber(l2))
}

func numberToList(num int) *ListNode {
	numStr := fmt.Sprint(num)
	nums := strings.Split(numStr, "")
	slices.Reverse(nums)

	lst := &ListNode{}
	cur := lst
	for i, n := range nums {
		nInt, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		cur.Val = nInt
		if i < len(nums)-1 {
			cur.Next = &ListNode{}
		}
		cur = cur.Next
	}

	return lst
}

func listToNumber(lst *ListNode) int {
	nums := []string{}
	cur := lst
	for cur != nil {
		nums = append(nums, fmt.Sprint(cur.Val))
		cur = cur.Next
	}

	slices.Reverse(nums)
	numStr := strings.Join(nums, "")
	numInt, err := strconv.Atoi(string(numStr))
	if err != nil {
		panic(err)
	}

	return numInt
}
