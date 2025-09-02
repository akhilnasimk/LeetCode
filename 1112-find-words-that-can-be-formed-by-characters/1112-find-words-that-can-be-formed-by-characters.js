var countCharacters = function(words, chars) {
    let leng = 0;
    let charCount = {};

    // count chars frequency
    for (let a of chars) {
        charCount[a] = (charCount[a] || 0) + 1;
    }

    for (let word of words) {
        let temp = {};
        let flag = true;

        // count word frequency
        for (let b of word) {
            temp[b] = (temp[b] || 0) + 1;
        }

        // compare only unique letters in word
        for (let key in temp) {
            if (!charCount[key] || temp[key] > charCount[key]) {
                flag = false;
                break;   // stop early if not valid
            }
        }

        if (flag) leng += word.length;
    }
    return leng;
};
