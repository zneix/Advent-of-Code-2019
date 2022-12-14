const { compareLists } = require("./util.js");
const input = require("fs").readFileSync("input", "utf-8").trim();
const pairs = input.split("\n\n");

console.log(
	pairs.reduce((total, curr, i) => {
		const pair = curr.split("\n");
		const left = JSON.parse(pair[0]);
		const right = JSON.parse(pair[1]);

		return total += compareLists(left, right) ? i+1 : 0;
	}, 0)
);
