// --- Directions
// Check to see if two provided strings are anagrams of eachother.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case
// --- Examples
//   anagrams('rail safety', 'fairy tales') --> True
//   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//   anagrams('Hi there', 'Bye there') --> False

function anagrams(stringA, stringB) {
    stringA.split(" ").join("").toLowerCase();
    stringB.split(" ").join("").toLowerCase();
    let map = {};
    for (let c of stringA) {
        if (!map[c]) {
            map[c] = 1;
        } else {
            map[c]++;
        }
    }
    for (let c of stringB) {
        console.log(c + " " + map[c]);
        if (!map[c]) {
            return false;
        }
        map[c]--;
        if (map[c] < 0) {
            return false;
        }
    }
    console.log(map);
    return true;
}

console.log(anagrams("A tree, a life, a bench", "A tree, a fence, a yard"));
module.exports = anagrams;
