func maxBottlesDrunk(numBottles int, numExchange int) int {
    drinked:=numBottles
    emty:=numBottles 
    for emty>=numExchange {
        emty=(emty-numExchange)+1
        drinked++
        numExchange++
    }
    return drinked
}