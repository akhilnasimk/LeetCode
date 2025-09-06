/**
 * @param {string} address
 * @return {string}
 */
var defangIPaddr = function(address) {
    let res=""
    for(let i of address){
        if(i=="."){
            res+=`[${i}]`
        }
        else{
            res+=i;
        }
    }
    return res
};