/**
 * @param {string[]} operations
 * @return {number}
 */
var finalValueAfterOperations = function(operations) {
    let value=0;
    for(let a of operations){
        if(a=="X++" || a=="++X"){
            value+=1;
        }
        else{
            value-=1;
        }
    }
    return value
};