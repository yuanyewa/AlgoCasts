// --- Directions
// Given the root node of a tree, return
// an array where each element is the width
// of the tree at each level.
// --- Example
// Given:
//     0
//   / |  \
// 1   2   3
// |       |
// 4       5
// Answer: [1, 3, 2]

function levelWidth(root) {
    if (!root) {
        return [];
    }
    let res = [0];
    let arr = [root, null];
    while (arr.length > 1) {
        const node = arr.shift();
        if (node === null) {
            res.push(0);
            arr.push(node);
        } else {
            res[res.length - 1]++;
            arr.push(...node.children);
        }
    }
    return res;
}

module.exports = levelWidth;