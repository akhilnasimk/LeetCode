/**
 * @param {string} licensePlate
 * @param {string[]} words
 * @return {string}
 */
var shortestCompletingWord = function(licensePlate, words) {
    let result = licensePlate.replace(/[^a-zA-Z]/g, "").toLowerCase(); 
    let short=null
    // console.log(result)
    let LMap= new Map()
    for(let a of result){
        LMap.has(a)?LMap.set(a,LMap.get(a)+1):LMap.set(a,1)
    }
    // console.log(LMap)

    for(let a of words ){
        let newMap= new Map()
        for(let b of a ){
            newMap.has(b)?newMap.set(b,newMap.get(b)+1):newMap.set(b,1);
        }
        console.log(newMap)
        // if(newMap.size !== LMap.size )continue 
        let match=true 
        for (let [key, val] of LMap) {
            if (!newMap.has(key) || newMap.get(key) < val) {
                match = false;
                break;
         }
            }
        if(match){
            if(short==null || a.length < short.length){
                short=a
            }
        }
    }
    return  short 
};