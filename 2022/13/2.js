const { compareLists } = require("./util.js");
const input = require("fs").readFileSync(process.argv[2] ?? "input", "utf-8");
const packets = input.split("\n").filter(Boolean).map(str => JSON.parse(str));

// inject special packets
packets.push([[2]], [[6]]);

// measure indices of special packets (they're one-indexed) by checking how many packets are smaller than these
let specialIndex = [1, 1];
for (let i = 0; i < packets.length; i++) {
	if (compareLists(packets[i], [[2]])) {
		specialIndex[0]++;
		specialIndex[1]++;
	} else if (compareLists(packets[i], [[6]])) {
		specialIndex[1]++;
	}
}

// get multiplied, indices of special packets
console.log(specialIndex[0] * specialIndex[1]);
