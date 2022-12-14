let img = require('fs').readFileSync('input').toString();
let phases = [];
for (let i=0;i<img.length;i+=150){
    phases.push(img.substring(i, 150+i));
}
let zeroes = [];
phases.forEach(key => {
    let count = 0;
    for (let i=0;i<key.length;i++) if (key.charAt(i)==0) count++;
    zeroes.push(count);
});
let rstr = phases[zeroes.indexOf(Math.min(...zeroes))]; //right-version string ♂
let ones = count(rstr, 1);
let twos = count(rstr, 2);
function count(str, num){
    let cnt = 0;
    for (let i=0;i<str.length;i++) if (str.charAt(i)==num) cnt++;
    return cnt;
}
console.log(ones*twos);