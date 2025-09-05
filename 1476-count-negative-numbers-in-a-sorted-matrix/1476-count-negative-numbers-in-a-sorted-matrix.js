/**
 * @param {number[][]} grid
 * @return {number}
 */
var countNegatives = function(grid) {
    let neg=0;
    for(let i of grid){
        for(let k of i ){
            if(k<0){
                neg++
            }
        }
    }
    return neg
};