/**
 * @param {string} boxes
 * @return {number[]}
 */
var minOperations = function(boxes) {
    let res=[]
    let sum=0;
    for(let j=0;j<boxes.length;j++){
        for(let i=0;i<boxes.length;i++){
            if(boxes[i]==1){
                sum+=Math.abs(i-j)
            }
        }
        res.push(sum)
        sum=0;
    }
    return res
};