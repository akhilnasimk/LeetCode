func numWaterBottles(numBottles int, numExchange int) int {
    total :=numBottles 
    reminder:=0
    for (numBottles)>=numExchange{
        total+=numBottles/numExchange;
        fmt.Println(numBottles/numExchange)
        reminder= numBottles % numExchange
        numBottles= (numBottles/numExchange) + reminder 
    }
    return total 
}