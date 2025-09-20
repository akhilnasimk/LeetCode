/**
 * @param {number[]} nums
 * @return {number}
 */
var thirdMax = function(nums) {
    let first=-Infinity
    let second=-Infinity
    let Third=-Infinity
    for(let a of nums){
         if (a === first || a === second || a === Third) continue;

        if (a > first) {
            third = second;
            second = first;
            first = a;
        } else if (a > second) {
            third = second;
            second = a;
        } else if (a > third) {
            third = a;
        }
    }
    return third==-Infinity? first : third
};