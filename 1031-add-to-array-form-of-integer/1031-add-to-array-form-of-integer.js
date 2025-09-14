/**
 * @param {number[]} num
 * @param {number} k
 * @return {number[]}
 */
var addToArrayForm = function(num, k) {
     let sum = "";

    for (let a of num) {
        sum += a;
    }
    console.log("Original array as string:", sum);

    sum = BigInt(sum) + BigInt(k);
    console.log("After addition:", sum.toString());

    let res = [];
    for (let a of sum.toString()) {
        res.push(Number(a));
    }

    return res;
};