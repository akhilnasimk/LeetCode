func subsets(nums []int) [][]int {
    var backtrack func(int)
    res:=[][]int{}

    subset:=[]int{}

    backtrack=func (i int ){
        if i>=len(nums){
            res = append(res, append([]int{}, subset...)) 
            return 
        }


        subset=append(subset,nums[i])
        backtrack(i+1)

        subset=subset[:len(subset)-1]
        backtrack(i+1)
    }

    backtrack(0)
    return res
}