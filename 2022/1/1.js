let input = require("fs").readFileSync("input", "utf-8").slice(0, -1);
let elves = input.split("\n\n");

let each = elves
	.map((x) => x.split("\n").map(Number).reduce((sum, curr) => sum += curr, 0))
	.sort((a, b) => b - a);

console.log(each[0]);
console.log(each[0] + each[1] + each[2]);
