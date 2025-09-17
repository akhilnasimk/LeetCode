/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function(strs) {
    let map=new Map()
    for(let a of strs){
        let sorted=a.split("").sort().join("")
        if(!map.has(sorted)){
              map.set(sorted,[])
        }
        map.get(sorted).push(a)
        
       
        
        
    }
    return (Array.from(map.values()))
};