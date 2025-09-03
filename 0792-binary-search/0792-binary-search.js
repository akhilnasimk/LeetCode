/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    let start=0;
    let end=nums.length-1;
    while(end>=start){
        let mid=Math.floor((start+end)/2);
        // console.log(nums[mid]);
        if(target===nums[mid]){
            return nums.indexOf(nums[mid]);
        }
        else if(target>nums[mid]){
            start=mid+1;
        }
        else{
            end=mid-1;
        }
    }
    return -1;
};