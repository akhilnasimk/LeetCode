/**
 * @param {string} boxes
 * @return {number[]}
 */
var minOperations = function(boxes) {
    let res=[]
    for(let j=0;j<boxes.length;j++){
        let sum=0;
        for(let i=0;i<boxes.length;i++){
            if(boxes[i]==1){
                sum+=Math.abs(i-j)
            }
        }
        res.push(sum)
    }
    return res
};