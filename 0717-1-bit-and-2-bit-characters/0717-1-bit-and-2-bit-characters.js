/**
 * @param {number[]} bits
 * @return {boolean}
 */
var isOneBitCharacter = function(bits) {
    let i =0
    let res=[]
    while(i<bits.length)
    {
        if(bits[i]==0){
            res.push(0)
        }else if(bits[i]==1 && bits[i+1]!==undefined ){
            res.push([bits[i],bits[i+1]])
            i++
        }
        i++
    }
    // console.log(res)
    if(res[res.length-1]===0){
        return true 
    }
    return false 
};