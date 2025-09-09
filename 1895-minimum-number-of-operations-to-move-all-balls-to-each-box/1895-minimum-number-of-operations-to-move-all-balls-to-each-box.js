/**
 * @param {string} boxes
 * @return {number[]}
 */
var minOperations = function(boxes) {
    let res=[]
    let index=0;
    while(res.length !=boxes.length){
        let sum=0;
        for(let i=0;i<boxes.length;i++){
            if(boxes[i]==1){
                sum+=Math.abs(i-index)
            }
        }
        res.push(sum)
        index++
    }
    return res
};