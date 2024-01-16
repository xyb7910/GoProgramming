package thefirstday

import "strconv"

// 1.两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

// 2.两数相加

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	head := &ListNode{Val: 0, Next: nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		current.Next = &ListNode{Val: (x + y + carry) % 10, Next: nil}
		current = current.Next

		carry = (x + y + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		current.Next = &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
	}
	return head.Next
}

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	//滑动窗口 返回长度即可
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

// 7. 整数反转
func reverse(x int) int {
	temp := 0
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
	}
	if temp > 1<<31-1 || temp < -(1<<31) {
		return 0
	}
	return temp
}

// 9. 回文数
func isPalindrome(x int) bool {
	//先进行特殊处理
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	str := strconv.Itoa(x)
	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
