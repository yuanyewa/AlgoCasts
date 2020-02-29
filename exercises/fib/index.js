// --- Directions
// Print out the n-th entry in the fibonacci series.
// The fibonacci series is an ordering of numbers where
// each number is the sum of the preceeding two.
// For example, the sequence
//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
// forms the first ten entries of the fibonacci series.
// Example:
//   fib(4) === 3

// function fib(n) {
//     let res = [0, 1];
//     for (let i = 2; i <= n; i++) {
//         res[i] = res[i - 1] + res[i - 2];
//     }
//     return res[n];
// }

// function fib(n) {
//     if (n == 0 || n == 1) {
//         return n;
//     }
//     return fib(n - 1) + fib(n - 2);
// }

// // function memorize(fn) {
// //     const cache = {};
// //     return function(...args) {
// //         if (cache[args]) {
// //             return cache[args];
// //         }
// //         const res = fn.apply(this, args);
// //         cache[args] = res;
// //         return res;
// //     }
// // }
// function memorize(fn) {
//     const cache = {};
//     return function(n) {
//         if (cache[n]) {
//             return cache[n];
//         }
//         const res = fn(n);
//         cache[n] = res;
//         return res;
//     }
// }

// fib = memorize(fib);

cache = [];

function fib(n) {
    if (n == 0 || n == 1) {
        return n;
    }
    if (cache[n]) {
        return cache[n];
    }
    const res = fib(n - 1) + fib(n - 2);
    cache[n] = res;
    return res;
}

module.exports = fib;