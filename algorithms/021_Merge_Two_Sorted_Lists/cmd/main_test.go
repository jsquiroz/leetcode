package main

import (
	"log"
	"testing"
)

type values struct {
	letter       string
	expectResult int
}

func TestMain(t *testing.T) {
	uno_1 := ListNode{}
	dos_1 := ListNode{}
	tres_1 := ListNode{}

	uno_2 := ListNode{}
	dos_2 := ListNode{}

	uno_1.Val = 1
	uno_1.Next = &dos_1

	dos_1.Val = 2
	dos_1.Next = &tres_1

	tres_1.Val = 3

	uno_2.Val = 4
	uno_2.Next = &dos_2

	dos_2.Val = 5

	nodo := MergeTwoSortedLists(&uno_1, &uno_2)

	for nodo != nil {
		log.Println(nodo.Val)
		nodo = nodo.Next
	}
}
