/**
 * @param {number} n
 * @return {number}
 */
var smallestEvenMultiple = function(n) {
    let i=1;
    let s=0;
    while(true){
        if(i%2==0&&i%n==0){
            s=i;
            break;
        }
        else{
            i++;
        }
    }
    return s
};