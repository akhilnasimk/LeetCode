/**
 * @param {string[]} words
 * @param {string} chars
 * @return {number}
 */
var countCharacters = function(words, chars) {
    let leng=0;
    let charCount={}
    for(let a of chars ){
        if(charCount[a]){
            charCount[a]=charCount[a]+1
        }
        else{
            charCount[a]=1;
        }
    }
    console.log(charCount);
    for(let a of words){
        let temp={}
        let flag=true;
        for(let b of a ){
            if(temp[b]){
                temp[b]=temp[b]+1
            }
            else{
                temp[b]=1;
            }
        }
        console.log(temp)
        for(let j of a){
            if(charCount[j]>=temp[j]){
                flag=flag
            }
            else{
                flag=false
            }
        }
        if(flag){
            leng+=a.length
        }
    }
    return leng
};