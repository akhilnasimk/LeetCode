/**
 * @param {string} s
 * @param {string} t
 * @return {character}
 */
var findTheDifference = function(s, t) {
    const stringMap=new Map()
    for(let a of s){
        stringMap.has(a)?stringMap.set(a,stringMap.get(a)+1):stringMap.set(a,1);
    }
    console.log(stringMap);
    let tMap=new Map()
    for(let a of t){
        tMap.has(a)?tMap.set(a,tMap.get(a)+1):tMap.set(a,1);
    }
    console.log(tMap)

    for(let a of t){
        if(stringMap.has(a)){
            if(tMap.get(a)>stringMap.get(a)){
                return a
            }
        }
        else{
            return a
        }
    }
};