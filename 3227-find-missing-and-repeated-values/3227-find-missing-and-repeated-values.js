/**
 * @param {number[][]} grid
 * @return {number[]}
 */
var findMissingAndRepeatedValues = function(grid) {
    let res=[]
    grid=grid.flat(Infinity).sort((a,b)=>a-b)
    for(let a=0;a<grid.length;a++){
        if(!grid.includes(a+1)){
            res[1]=a+1;
        }
        if(grid[a]==grid[a+1]){
            res[0]=grid[a]
        }
    }
    return res

};