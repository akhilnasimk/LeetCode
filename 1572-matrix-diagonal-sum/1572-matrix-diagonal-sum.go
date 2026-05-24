func diagonalSum(mat [][]int) int {
   n:=len(mat)
   sum:=0
   for row,val:=range mat{
        for col,v:= range val{
            if row==col || row+col+1==n{
                sum+=v
            }
        }
   } 
   return sum 
}