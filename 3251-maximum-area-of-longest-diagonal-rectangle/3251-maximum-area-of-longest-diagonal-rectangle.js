/**
 * @param {number[][]} dimensions
 * @return {number}
 */
var areaOfMaxDiagonal = function(dimensions) {
    let Hdaig=0;
    let res=0
    for(let i=0;i<dimensions.length ;i++){
        let diag=Math.sqrt(dimensions[i][0]**2 + dimensions[i][1]**2)
        let area= dimensions[i][0] * dimensions[i][1]
        if(diag>Hdaig){
            Hdaig=diag;
            res=area;
        }else if (diag === Hdaig && area > res) {
            res = area;
        }
    }
    return (res)
};