/**
 * @param {string} text
 * @param {string} brokenLetters
 * @return {number}
 */
var canBeTypedWords = function(text, brokenLetters) {
    text=text.split(" ");
    let count=0;
    // console.log(text)
    brokenLetters=brokenLetters.split("")
    for(let a of text){
        let set=new Set(a)
        let flag=true;
        for(let b of brokenLetters){
            if(set.has(b)){
                flag=false;
            }
            else{
                flag=flag
            }
        }
        if(flag){
            count++
        }
    }

    return(count)
    
};