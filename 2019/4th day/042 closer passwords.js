let min = 158126;
let max = 624574;
let sum = 0;
for (let i=min;i<max;i++) if (dC(i) && noDec(i)) sum++;
function dC(num){
    let n = num.toString().split('');
    for (i=0;i<5;i++) if (n[i] == n[i+1]) if ((n[i-1] == undefined || n[i-1] != n[i]) && (n[i+2] == undefined || n[i+2] != n[i])) return true;
}
function noDec(num){
    let n = num.toString().split('').reverse();
    for (j=0;j<5;j++) if (n[j] < n[j+1]) return false;
    return true;
}
console.log(sum);