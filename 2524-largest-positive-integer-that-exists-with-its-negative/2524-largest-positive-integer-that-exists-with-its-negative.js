/**
 * @param {number[]} nums
 * @return {number}
 */
var findMaxK = function(nums) {
    let largest=0
    let check=new Set(nums);
    for(let a of nums){
        if(a>largest&&check.has(-a)){
            largest=a;
        }
    }
    return largest!==0?largest:-1
};