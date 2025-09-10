/**
 * @param {number} n
 * @return {number}
 */
var alternateDigitSum = function(n) {
    let str=n+"";
    let sum=0
    for(let i=0;i<str.length;i++){
        if((i+1)%2!==0){
            console.log( "first"+str[i])
            sum+=Number(str[i])
        }else{
            console.log("second"+str[i])
            sum-=Number(str[i])
        }
    }
    return sum
};