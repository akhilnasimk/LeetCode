/**
 * @param {string[]} sentences
 * @return {number}
 */
var mostWordsFound = function(sentences) {
    let max=0
    for(let a of sentences){
        let newAr=a.split(" ")
        if(newAr.length>max){
            max=newAr.length
        }
    }
    return max
};