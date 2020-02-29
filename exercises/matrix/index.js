// --- Directions
// Write a function that accepts an integer N
// and returns a NxN spiral matrix.
// --- Examples
//   matrix(2)
//     [[1, 2],
//     [4, 3]]
//   matrix(3)
//     [[1, 2, 3],
//     [8, 9, 4],
//     [7, 6, 5]]
//  matrix(4)
//     [[1,   2,  3, 4],
//     [12, 13, 14, 5],
//     [11, 16, 15, 6],
//     [10,  9,  8, 7]]

function matrix(n, v = 0, x = 0, y = 0, dx = 0, dy = 1, mx = []) {
    if (v > n * n) {
        console.log(mx);
        return mx;
    }
    if (v === 0) {
        for (let i = 0; i < n; i++) {
            mx.push([]);
            for (let j = 0; j < n; j++) {
                mx[i].push(0);
            }
        }
        return matrix(n, v + 1, x, y, dx, dy, mx);
    } else {
        mx[x][y] = v;
        if (dx != 0 && (x + dx == n || x + dx < 0 || mx[x + dx][y] > 0)) {
            dy = -dx;
            dx = 0;
        } else if (dy != 0 && (y + dy == n || y + dy < 0 || mx[x][y + dy] > 0)) {
            dx = dy;
            dy = 0;
        }
        return matrix(n, v + 1, x + dx, y + dy, dx, dy, mx);
    }
}

module.exports = matrix;