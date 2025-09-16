/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let map=new Map()
    for(let a of nums){
        map.has(a)?map.set(a,map.get(a)+1):map.set(a,1);
    }
    for(let [key,val] of map){
        if(val>2){
            let count=val-2
            while(count>0){
                let index=nums.indexOf(key)
                nums.splice(index,1)
                count--
            }
        }
    }
};