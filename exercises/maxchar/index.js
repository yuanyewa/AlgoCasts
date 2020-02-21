// --- Directions
// Given a string, return the character that is most
// commonly used in the string.
// --- Examples
// maxChar("abcccccccd") === "c"
// maxChar("apple 1231111") === "1"

function maxChar(str) {
    let map = {};
    for (let char of str.split('')) {
        if (!map[char]) {
            map[char] = 1;
        } else {
            map[char]++;
        }
    }
    let mx = 0;
    let c = '';
    for (let char in map) {
        if (map[char] > mx) {
            mx = map[char];
            c = char;
        }
    }
    console.log(map);
    console.log(mx);
    return c;
}

module.exports = maxChar;
