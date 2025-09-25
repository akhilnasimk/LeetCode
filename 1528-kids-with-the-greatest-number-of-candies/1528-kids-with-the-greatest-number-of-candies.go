func kidsWithCandies(candies []int, extraCandies int) []bool {
    max:=0
    for _,val:=range candies{
        if val>max{
            max=val
        }

    }
    res:=[]bool{}
    for _,value :=range candies{
        if(value+extraCandies>=max){
            res=append(res,true)
        }else{
            res=append(res,false)
        }
    }
    return res
}