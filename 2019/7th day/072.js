let oarr = [3,8,1001,8,10,8,105,1,0,0,21,46,67,76,101,118,199,280,361,442,99999,3,9,1002,9,4,9,1001,9,2,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,1001,9,3,9,102,2,9,9,1001,9,2,9,1002,9,3,9,4,9,99,3,9,101,3,9,9,4,9,99,3,9,1001,9,2,9,1002,9,5,9,101,5,9,9,1002,9,4,9,101,5,9,9,4,9,99,3,9,102,2,9,9,1001,9,5,9,102,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,99];
let res;
let permuts = [];
for (let a=5;a<10;a++){for (let b=5;b<10;b++){for (let c=5;c<10;c++){for (let d=5;d<10;d++){for (let e=5;e<10;e++){
    if (a != b && a != c && a != d && a != e && b != c && b != d && b != e && c != d && c != e && d != e) permuts.push([a, b, c, d, e]);
}}}}} //yeah, cool way to do it, eh?
let thrusters = [];
permuts.forEach(perm => {
    end = false;
    indexes = [0,0,0,0,0];
    nexts = [[],[],[],[],[]];
    arrays = [[...oarr],[...oarr],[...oarr],[...oarr],[...oarr]];
    for (let i=0;i<5;i++) nexts[i].push(perm[i]);
    nexts[0].push(0);
    while (!end) for (let j=0;j<5;j++) thruster = intcode(arrays[j], indexes[j], j, nexts[j], nexts[(j+1)%5]);
    thrusters.push(thruster);
});
console.log(Math.max(...thrusters));
function intcode(arr, point, i, input, next){
    for (let asd=point;asd<arr.length;){
        let opc = arr[asd]%100;
        let x1 = arr[asd].toString().split('').reverse()[2]==1?arr[asd+1]:arr[arr[asd+1]];
        let x2 = arr[asd].toString().split('').reverse()[3]==1?arr[asd+2]:arr[arr[asd+2]];
        switch(opc){
            case 1:
                arr[arr[asd+3]] = x1 + x2;
                asd += 4;
                break;
            case 2:
                arr[arr[asd+3]] = x1 * x2;
                asd += 4;
                break;
            case 3:
                if (!input.length) return indexes[i] = asd;
                arr[arr[asd+1]] = input.shift();
                asd += 2;
                break;
            case 4:
                res = x1;
                indexes[i] = asd;
                next.push(res);
                asd += 2;
                break;
            case 5:
                if (x1 != 0) asd = x2;
                else asd += 3;
                break;
            case 6:
                if (x1 == 0) asd = x2;
                else asd += 3;
                break;
            case 7:
                arr[arr[asd+3]] = (x1 < x2)?1:0;
                asd += 4;
                break;
            case 8:
                arr[arr[asd+3]] = (x1 == x2)?1:0;
                asd += 4;
                break;
            case 99:
                next.push(res);
                if (i==4) end = true;
                return res;
        }
    }
}