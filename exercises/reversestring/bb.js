function reverse(str) {
    // return str.split('').reverse().join('');
    // for (var i = 0; i < str.length; i++)
    // let reversed = '';
    // for (let char of str) {
    //     reversed = char + reversed;
    // }
    // return reversed;
    str = str.split('');
    for (let i = 0, j = str.length - 1; i < j; i++, j--) {
        let tmp = str[i];
        str[i] = str[j];
        str[j] = tmp;
    }
    return str.join('');
    console.log("hihi");
    // return str.split('').reduce((rev, char) => char + rev, "");
}
console.log(reverse("abcd"));