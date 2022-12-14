#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_DIMENSION 99
char trees[MAX_DIMENSION][MAX_DIMENSION];

int main() {
	FILE *input = fopen("input", "r");
	if (!input) {
		printf("failed to open input file!\n");
		exit(1);
	}

	// read input
	char buf;
	int x = 0, y = 0; // vertical and horizontal line index
	while ((buf = fgetc(input)) != EOF) {
		if (buf == '\n') {
			y++;
			x = 0;
			continue;
		}
		trees[y][x] = buf;
		x++;
	}

	// logic
	int highestScore = 0; // vis trees at top * bottom * left * right
	for (int y = 1; y < MAX_DIMENSION - 1; y++) {
		for (int x = 1; x < MAX_DIMENSION - 1; x++) {
			// currently checking the tree at trees[y][x]
			int visibleLeft = 0, visibleRight = 0, visibleTop = 0, visibleBottom = 0;
			// check horizontally (x), to left
			for (int x1 = x - 1; x1 >= 0; x1--) {
				visibleLeft++;
				if (trees[y][x] <= trees[y][x1]) {
					break;
				}
			}
			// check horizontally (x), to right
			for (int x1 = x + 1; x1 < MAX_DIMENSION; x1++) {
				visibleRight++;
				if (trees[y][x] <= trees[y][x1]) {
					break;
				}
			}
			// check vertically (y), to top
			for (int y1 = y - 1; y1 >= 0; y1--) {
				visibleTop++;
				if (trees[y][x] <= trees[y1][x]) {
					break;
				}
			}
			// check vertically (y), to bottom
			for (int y1 = y + 1; y1 < MAX_DIMENSION; y1++) {
				visibleBottom++;
				if (trees[y][x] <= trees[y1][x]) {
					break;
				}
			}
			int currentScore = visibleLeft * visibleRight * visibleTop * visibleBottom;
			if (highestScore < currentScore) {
				highestScore = currentScore;
			}
			// end of current
		}
	}
	printf("%d\n", highestScore);

	// finish work
	fclose(input);
	return 0;
}
