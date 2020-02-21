// --- Directions
// Given a string, return a new string with the reversed
// order of characters
// --- Examples
//   reverse('apple') === 'leppa'
//   reverse('hello') === 'olleh'
//   reverse('Greetings!') === '!sgniteerG'

function reverse(str) {
    // return str.split('').reverse().join('');
    // for (var i = 0; i < str.length; i++)
    // let reversed = '';
    // for (let char of str) {
    //     reversed = char + reversed;
    // }
    // return reversed;
    // str = str.split('');
    // for (let i = 0, j = str.length - 1; i < j; i++, j--) {
    //     let tmp = str[i];
    //     str[i] = str[j];
    //     str[j] = tmp;
    // }
    // return str.join('');
    return str.split('').reduce((rev, char) => char + rev, "");
}

reverse("abced");
module.exports = reverse;
