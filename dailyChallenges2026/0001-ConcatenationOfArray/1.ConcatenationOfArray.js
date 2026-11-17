/**
 * @param {number[]} nums
 * @return {number[]}
 */
var getConcatenation = function(nums) {
    var result = []

    for(let i=0; i<2; i++) {
        for (let n of nums) {
            result.push(n);
        }
    }
    return result;
};