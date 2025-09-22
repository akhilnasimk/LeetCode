/**
 * @param {number[]} nums
 * @return {number}
 */
var maxFrequencyElements = function(nums) {
    let map= new Map()
    for (let a of nums ){
        map.has(a)?map.set(a,map.get(a)+1):map.set(a,1)
    }
    console.log(map)
    let hfr=0;
    let h=0
    for( let [key,value] of map){
        if(value>h){
            hfr=value;
            h=value;
        }else if(value==h){
            hfr+=value;
        }
        else{
            hfr=hfr
        }
    }
    return (hfr)
};