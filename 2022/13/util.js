// if true, 'left' is smaller than 'right'
function compareLists(left, right) {
	for (let k = 0; k < left.length + 1; k++) {
		if (left[k] === undefined) {
			if (right[k] !== undefined) {
				// on last iteration left[k] should be undefined - just check if right isn't (that's what we want)
				return true;
			} else {
				// both elements ended together, return with no result
				return;
			}
		} else if (right[k] === undefined) {
			// right is shorter than left, that's bad
			return false;
		}

		if (typeof(left[k]) === typeof(right[k])) {
			// same types
			if (typeof(left[k]) === "number") {
				// int vs int
				if (left[k] > right[k]) {
					return false;
				}
				if (left[k] < right[k]) {
					return true;
				}
			} else if (typeof(left[k]) === "object") {
				// array vs array, here's funny
				let ret = compareLists(left[k], right[k]);
				if (typeof(ret) === "boolean") {
					return ret;
				}
			}
		} else {
			// mismatched types
			if (typeof(left[k]) === "number" && typeof(right[k]) === "object") {
				let ret = compareLists([left[k]], right[k]);
				if (typeof(ret) === "boolean") {
					return ret;
				}
			} else if (typeof(left[k]) === "object" && typeof(right[k]) === "number") {
				let ret = compareLists(left[k], [right[k]]);
				if (typeof(ret) === "boolean") {
					return ret;
				}
			}
		}
	}
}

module.exports = { compareLists };
