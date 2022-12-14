const input = require("fs").readFileSync("input", "utf-8");

let monkeys = input.split("\n\n").map((monkeyString, i) => {
	let attributes = monkeyString.split("\n");
	let operationStr = attributes[2].slice(attributes[2].indexOf("old") + 4);

	return {
		index: i,
		inspectCount: 0,
		items: attributes[1].slice(attributes[1].indexOf(":") + 2).split(", ").map(x => parseInt(x)),
		operation: {
			sign: operationStr[0],
			value: operationStr.slice(2),
		},
		test: parseInt(attributes[3].slice(attributes[3].lastIndexOf(" ") + 1)),
		ifTrue: parseInt(attributes[4].slice(attributes[4].lastIndexOf(" ") + 1)),
		ifFalse: parseInt(attributes[5].slice(attributes[5].lastIndexOf(" ") + 1)),
	}
});

const magicModuloShit = monkeys.reduce((total, monkey) => total *= monkey.test, 1);
const ROUND_COUNT = 10_000; // part 1: 20; part 2: 10_000

for (let r = 0; r < ROUND_COUNT; r++) {
	//console.log(`Round ${r+1}... fight!`);
	for (m of monkeys) {
		//console.log(`\nGo, #${m.index}!\n`, m.items);
		while (m.items.length > 0) {
			m.inspectCount++;
			let item = m.items.shift();
			// increase the worry level
			if (m.operation.sign === "+") {
				item += m.operation.value === "old" ? item : parseInt(m.operation.value);
			} else if (m.operation.sign === "*") {
				item *= m.operation.value === "old" ? item : parseInt(m.operation.value);
			}
			// keep worry level manageable
			//item = Math.floor(item / 3); // part 1
			item %= magicModuloShit; // part 2
			// give the item to one of the other monkeys
			monkeys[(item % m.test) === 0 ? m.ifTrue : m.ifFalse].items.push(item);
		}
	}
}
monkeys.sort((a, b) => b.inspectCount - a.inspectCount);
//console.log(monkeys.map(x => {return {index: x.index, i: x.inspectCount}}));
console.log((monkeys[0].inspectCount * monkeys[1].inspectCount));
