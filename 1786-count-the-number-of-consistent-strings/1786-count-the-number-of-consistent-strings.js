/**
 * @param {string} allowed
 * @param {string[]} words
 * @return {number}
 */
var countConsistentStrings = function(allowed, words) {
    let count=0
    let map=new Map()
    for(let a of allowed ){
        map.has(a)?map.set(a,map.get(a)+1):map.set(a,1)
    }
    for(let a of words){
        let flag=true;
        for(let b of a ){
            if(map.has(b)){
                falg=flag
            }
            else{
                flag=false
            }
        }
        if(flag){
            count++
        }
    }
    return count
};