/**
 * @param {number[]} nums
 * @return {number}
 */
var zeroFilledSubarray = function(nums) {
    let totalCount = 0;
    let zeroCount = 0;

    for (let num of nums) {
        if (num === 0) {
            zeroCount++;
            totalCount += zeroCount;
        } else {
            zeroCount = 0; 
        }
    }

    return totalCount;
};