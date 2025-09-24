/**
 * @param {number[]} nums
 * @return {number}
 */
var missingNumber = function(nums) {
    let actual =0
    let numlen=0
    for(let i=1;i<=nums.length;i++){
        actual+=i
        numlen+=nums[i-1]
    }
    return actual-numlen
};