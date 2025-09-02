/**
 * @param {number} x
 * @return {number}
 */
var sumOfTheDigitsOfHarshadNumber = function(x) {
    let str=String(x).split("");
    console.log(str);
    let sum=0;
    for(let a of str){
        sum+=Number(a);
    }
    console.log(sum);
    if(x%sum==0){
        return sum
    }
    else{
        return -1;
    }
};