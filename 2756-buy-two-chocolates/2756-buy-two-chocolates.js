/**
 * @param {number[]} prices
 * @param {number} money
 * @return {number}
 */
var buyChoco = function(prices, money) {
    let possible=[]
    for(let i=0;i<prices.length;i++){
        for(let j=i+1;j<prices.length;j++){
            if(prices[i]+prices[j]<=money){
                possible.push([prices[i],prices[j]]);
            }
        }
    }
    let smallest=100000000000000000000000000000000000000000000000000;
    let small=[]
    for(let i of possible ){
        if((i[0]+i[1])<=smallest){
            smallest=(i[0]+i[1])
            small=i;
        }
    }
    return small.length==0?money:(money-(small[0]+small[1]))

};