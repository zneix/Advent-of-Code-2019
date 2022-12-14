const input = require("fs").readFileSync("input", "utf-8").slice(0, -1);

//foe, me, definition, points
//A = X = Rock (0)
//B = Y = Paper (1)
//C = Z = Scissors (2)

// if wins object under given index contains opponent's value, it means you've won
const wins = {
	A: "C",
	B: "A",
	C: "B",
};

// XXX: Unnecessary, but idc
const losses = {
	C: "A",
	A: "B",
	B: "C",
};

const score = {
	A: 1,
	B: 2,
	C: 3,
};

let sum = 0;
for (const game of input.split("\n").map(x => x.split(" "))) {
	//if (game[1] !== "X") continue;
	//sum += score[game[1]];
	switch (game[1]) {
		case "Z": // win!
			//let xd = Object.values(wins).find(v => v === game[0]);
			let xd = losses[game[0]];
			sum += 6 + score[xd];
			console.log(game, "win", xd, score[xd]);
			break;
		case "Y": // draw!
			let b = game[0];
			sum += 3 + score[b];
			console.log(game, "draw", b, score[b]);
			break;
		case "X": // lose!
			let a = wins[game[0]];
			sum += 0 + score[a];
			console.log(game, "lose", a, score[a]);
			break;
	}
	// win
	//if (wins[game[1]] === game[0]) {
		////console.log(game, "win", sum);
		//sum += 6;
		//continue;
	//}
	//// draw
	//if (game[0] === game[1]) {
		////console.log(game, "draw", sum);
		//sum += 3;
		//continue;
	//}

	// loss
	//console.log(game, "loss", sum);
	//sum += 0;
}
console.log(sum);
