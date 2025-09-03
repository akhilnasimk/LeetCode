/**
 * @param {number} n
 * @return {boolean}
 */
var isThree = function(n) {
    let div=1;
    for(let i=1;i<=n/2;i++){
        if(n%i==0){
            div++
        }
    }
    return div==3?true:false;
};