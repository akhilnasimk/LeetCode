/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    num:=[]int{}
    cur:= head 
    for cur!=nil{
        num=append(num,cur.Val)
        cur=cur.Next 
    }
    left,right:=0,len(num)-1
    for left<=right{
        if num[left]!=num[right]{
            return false 
        }
        left++
        right--
    }
    // fmt.Println(num)``
    return true 
}