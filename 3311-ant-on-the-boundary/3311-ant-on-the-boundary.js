/**
 * @param {number[]} nums
 * @return {number}
 */
var returnToBoundaryCount = function(nums) {
    let num=0
    let sum=0;
    for(let a of nums){
        sum+=a;
        if(sum==0){
            num++;
        }
    }
    return num
};