/**
 * @param {string} command
 * @return {string}
 */
var interpret = function(command) {
    command=command.replace(/\(\)/g,"o");
    command=command.replace(/\(\a\l\)/g,"al")
    return (command)
};