let img = require('fs').readFileSync('input').toString();
let phases = [];
for (let i=0;i<img.length;i+=150){
    phases.push(img.substring(i, 150+i));
}
function forsenCD(image, next){ //transparency master
    let imageArr = image.split('');
    for (let k=0;k<image.length;k++) if (imageArr[k] == 2) imageArr[k] = next.charAt(k);
    return imageArr.join('');
}
let image = phases[0];
for (let j=1;j<phases.length;j++){
    next = phases[j];
    image = forsenCD(image, next);
}
for (let k=0;k<image.length;k+=75){
    image = image.replace(/0/g, '   ');
    image = image.replace(/1/g, '███');
    console.log(image.substring(k, k+75));
}