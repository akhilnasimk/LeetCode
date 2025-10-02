func maxBottlesDrunk(numBottles int, numExchange int) int {
    if numBottles<numExchange{
        return numBottles 
    }
    drinked:=numBottles
    emty:=numBottles 
    // remi:=0
    for emty>=numExchange {
        emty=(emty-numExchange)+1
        drinked++
        numExchange++
    }
    return drinked
}