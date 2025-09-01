/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    let counts=new Set()
    for(let a of nums){
        if(counts.has(a)){
            return true;
        }else{
            counts.add(a)
        }
    }
    return false
};