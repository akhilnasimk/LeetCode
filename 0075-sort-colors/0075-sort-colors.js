/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function(nums) {
    for(let i=0;i<nums.length;i++){
        for(let k=i+1;k<nums.length;k++){
            if(nums[i]>nums[k]){
                [nums[i],nums[k]]=[nums[k],nums[i]];
            }
        }
    }
    return nums;
};

// [0,0,2,2,1,1]
// []