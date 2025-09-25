func findWordsContaining(words []string, x byte) []int {
    count :=[]int {}
   for ind,val:=range words{
       for _,v:=range val{
            if v==rune(x) {
                count=append(count,ind)
                break;
            }
       }

   }
   return count

}