/**
 * @param {number[]} nums
 * @return {number[]}
 */
var decompressRLElist = function(nums) {
    let res=[]
    for(let j=0;j<nums.length;j+=2){
        let freq=nums[j]
        let number=nums[j+1];
        let temp=new Array(freq)
        temp.fill(number)
        res.push(temp)
    }
    return res.flat(Infinity)
};