package main

import (
	"fmt"
	"sort"
)

/*
给定两个数组，编写一个函数来计算它们的交集。
示例1 输入: nums1 = [1,2,2,1], nums2 = [2,2] 输出: [2,2]
示例2 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4] 输出: [4,9]

映射题（map映射）
我们需找出两个数组的交集元素，同时应与两个数组中出现的次数一致。
这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>
*/

func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		//遍历nums1，初始化map
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		//如果元素相同，将其存入nums2中，并将出现次数减1
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}

/*
进阶-数组是有序的
用双指针解法
*/
func intersect2(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums3 := intersect(nums1, nums2)
	fmt.Println("intersect nums3 = ", nums3)

	nums4 := []int{1, 2, 3, 4, 4, 13}
	nums5 := []int{1, 2, 3, 9, 10}
	nums6 := intersect2(nums4, nums5)
	fmt.Println("intersect2 nums6 = ", nums6)
}
