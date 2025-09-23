/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    let num =BigInt(digits.join(""))
    num=num+1n
    let res=[]
    while(num>0n){
        res.push(Number(num%10n))
        num = num/10n
    }
    return (res.reverse())
};