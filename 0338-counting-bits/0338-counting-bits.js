/**
 * @param {number} n
 * @return {number[]}
 */
var countBits = function(n) {
    let ans=[]
    for(let a =0;a<=n;a++){
        let bain="";
        let temp=a
        let count=0
        while(temp>0){
            bain+=temp%2;
            temp=Math.floor(temp/2)
        }
        for(let  a of bain ){
            if(a==1){
                count++
            }
        }
        console.log(bain)
        ans.push(Number(count))
    }
    return (ans)
};