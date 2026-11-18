/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
    const count = new Map();
    let result = 0;
    let maxCount = 0;

    for (const n of nums) {
        count.set(n, (count.get(n) || 0) + 1);

        if (count.get(n) > maxCount) {
            result = n;
            maxCount = count.get(n);
        }
    }

    return result;
};