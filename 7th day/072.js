let data = [3,8,1001,8,10,8,105,1,0,0,21,46,67,76,101,118,199,280,361,442,99999,3,9,1002,9,4,9,1001,9,2,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,1001,9,3,9,102,2,9,9,1001,9,2,9,1002,9,3,9,4,9,99,3,9,101,3,9,9,4,9,99,3,9,1001,9,2,9,1002,9,5,9,101,5,9,9,1002,9,4,9,101,5,9,9,4,9,99,3,9,102,2,9,9,1001,9,5,9,102,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,99];
let output;

let combinations = getAnagrams();
let vals = [];
combinations.forEach(arr => {
    finished = false;
    indexes = [0,0,0,0,0];
    let queues = [[],[],[],[],[]];
    let arrays = [[...data],[...data],[...data],[...data],[...data]];

    for(var i =0;i<5;i++) queues[i].push(arr[i]);
    queues[0].push(0);

    while(!finished){
        for(let j=0;j<5;j++) signal = getSignal(arrays[j],indexes[j],j,queues[j],queues[(j+1)%5]);
    }
    vals.push(signal);
});
console.log(Math.max(...vals));

function getSignal(data, iP, id, queue, dest){ // Amp's array, indexPointer, Amp's id, Amp's queue list, Amp's output receiver 
    for(let i=iP;i<data.length;i++){
        let instruction = data[i]%100;
        let x1 = parseInt(data[i].toString().split('').reverse()[2])==1?data[i+1]:data[data[i+1]];
        let x2 = parseInt(data[i].toString().split('').reverse()[3])==1?data[i+2]:data[data[i+2]];
        switch(instruction){
            case 1:
                data[data[i+3]] = x1 + x2
                i+=3;
                break;
            case 2:
                data[data[i+3]] = x1 * x2
                i+=3;
                break;
            case 3:
                // console.log(queue)
                if (queue.length<1) return indexes[id] += 2;
                data[data[i+1]] = queue.shift();
                i+=1;
                break;
            case 4:
                output = x1;
                indexes[id] = i;
                dest.push(output);
                i+=1;
                break;
            case 5:
                x1 !== 0 ? i = x2 - 1 : undefined;          
                break;
            case 6:
                x1 == 0 ? i = x2 - 1 : undefined;          
                break;
            case 7:
                x1 < x2 ? data[data[i+3]] = 1 : data[data[i+3]] = 0;
                i+=3;
                break;
            case 8:
                x1 == x2 ? data[data[i+3]] = 1 : data[data[i+3]] = 0;
                i+=3;
                break;
            case 99:
                dest.push(output)
                id==4 ? finished = true : undefined;
                return output
        }
    }
}
function getAnagrams() {
    let arr = [];
    let sign = 5;
    for (let a=sign;a<sign+5;a++){for (let b=sign;b<sign+5;b++){for (let c=sign;c<sign+5;c++){for (let d=sign;d<sign+5;d++){for (let e=sign;e<sign+5;e++){
        if (a != b && a != c && a != d && a != e && b != c && b != d && b != e && c != d && c != e && d != e) arr.push([a, b, c, d, e]);
    }}}}} //yeah, cool way to do it, eh?
    return arr;
}