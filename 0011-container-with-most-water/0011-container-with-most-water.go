func maxArea(height []int) int {
    max:=0
    left:=0
    right:=len(height)-1
    for left<right {
        min:= Min(height[left],height[right])
        Area:=min*(right-left)
        if Area > max{
            max=Area
        }
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return max
}

func Min(a,b int)int{
    if(a>b){
        return b
    }
    return a
}