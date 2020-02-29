// --- Directions
// Write a function that accepts a positive number N.
// The function should console log a step shape
// with N levels using the # character.  Make sure the
// step has spaces on the right hand side!
// --- Examples
//   steps(2)
//       '# '
//       '##'
//   steps(3)
//       '#  '
//       '## '
//       '###'
//   steps(4)
//       '#   '
//       '##  '
//       '### '
//       '####'

// function steps(n) {
//     for (let r = 0; r < n; r++) {
//         let str = '';
//         for (let c = 0; c < n; c++) {
//             if (c <= r) {
//                 str += "#";
//             } else {
//                 str += " ";
//             }
//         }
//         console.log(str);
//     }
// }
function steps(n, r = 0, s = "") {
    if (r === n) {
        return;
    }
    if (s.length === n) {
        console.log(s);
        steps(n, r + 1);
    } else {
        if (s.length <= r) {
            s += "#";
        } else {
            s += " ";
        }
        steps(n, r, s);
    }
}

module.exports = steps;