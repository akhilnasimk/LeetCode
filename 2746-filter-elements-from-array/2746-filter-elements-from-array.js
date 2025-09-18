/**
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 */
var filter = function(arr, fn) {
    let filteredArr=[]
    for(let i=0;i<arr.length;i++){
        let bool=fn(arr[i],i);
        if(bool){
            filteredArr.push(arr[i])
        }
    }
    return filteredArr
    
};