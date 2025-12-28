func countNegatives(grid [][]int) int {
    count:=0
    for _,Val:=range grid{
        i:=len(Val)-1
        for Val[i]<0{
            count++
            i--
            if i==-1{
                break 
            }
        }
    }

    return count 
}