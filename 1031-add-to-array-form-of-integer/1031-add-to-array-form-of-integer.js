/**
 * @param {number[]} num
 * @param {number} k
 * @return {number[]}
 */
var addToArrayForm = function(num, k) {
    let i = num.length - 1;
    let carry = k;

    while (i >= 0 || carry > 0) {
        if (i >= 0) {
            carry += num[i];      
            num[i] = carry % 10;  
            carry = Math.floor(carry / 10);  
            i--;
        } else {
            num.unshift(carry % 10);
            carry = Math.floor(carry / 10);
        }
    }

    return num;
};