// --- Directions
// Write a function that accepts a positive number N.
// The function should console log a pyramid shape
// with N levels using the # character.  Make sure the
// pyramid has spaces on both the left *and* right hand sides
// --- Examples
//   pyramid(1)
//       '#'
//   pyramid(2)
//       ' # '
//       '###'
//   pyramid(3)
//       '  #  '
//       ' ### '
//       '#####'

function pyramid(n, row = 0, str = '') {
    if (row === n) {
        return;
    }
    if (str.length === 2 * n - 1) {
        console.log(str);
        pyramid(n, row + 1, '');
    } else {
        if (str.length < n - 1 - row || str.length > n - 1 + row) {
            str += ' ';
        } else {
            str += '#';
        }
        pyramid(n, row, str);
    }
}

module.exports = pyramid;