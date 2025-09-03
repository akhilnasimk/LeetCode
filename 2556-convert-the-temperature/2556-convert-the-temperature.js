/**
 * @param {number} celsius
 * @return {number[]}
 */
var convertTemperature = function(celsius) {
    let ar=[];
    ar.push(celsius+273.15);
    ar.push((celsius*(9/5))+32);
    return ar
};