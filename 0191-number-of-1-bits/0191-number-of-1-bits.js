/**
 * @param {number} n
 * @return {number}
 */
var hammingWeight = function(n) {
    let a=""
    while(n>0){
        a+=n%2
        n=Math.floor(n/2)
    }
    console.log(a)
    let count=0
    for(let b of a ){
        if(b=="1"){
            count++
        }
    }
    return count 
};