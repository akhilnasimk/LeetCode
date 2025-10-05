func numberGame(nums []int) []int {
   res:=[]int{}
//    rand:=101
// fmt.Println(len(res),len(nums))
   for len(res)!=len(nums) {
        // fmt.Println("hye")
        smallest := [1][2]int{{101, 0}}
        second := [1][2]int{{101, 0}}
        for ind,val:=range nums{
            // fmt.Println(val,smallest[0][0])
            if val<smallest[0][0]{
                // fmt.Println("hi")
                second[0][0]= smallest[0][0]
                second[0][1]=smallest[0][1]
                smallest[0][0]=val 
                smallest[0][1]=ind
            }else if val<second[0][0]{
                // fmt.Println("hi1")
                second[0][0]= val 
                second[0][1]=ind
            }
        }
        // fmt.Println(second[0][0],smallest[0][1])
        nums[second[0][1]]=100
        nums[smallest[0][1]]=100
        res=append(res,second[0][0],smallest[0][0])
   }
   return res 
}