/**
 * @param {Object|Array} obj
 * @return {boolean}
 */
var isEmpty = function(obj) {
    let count=0
    for(let a in obj){
        count++
    }
    return count==0? true :false
};