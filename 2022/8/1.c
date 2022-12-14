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
	int visibleTrees = 0;
	for (int y = 1; y < MAX_DIMENSION - 1; y++) {
		for (int x = 1; x < MAX_DIMENSION - 1; x++) {
			// currently checking the tree at trees[y][x]
			// check horizontally (x), from left
			bool visibleLeft = true;
			for (int x1 = 0; x1 < x; x1++) {
				if (trees[y][x] <= trees[y][x1]) {
					visibleLeft = false; // we already know it will not be visible
					break;
				}
			}
			// we already know current tree will be visible from at least one angle, no need to
			// check its visibility from other sides
			if (visibleLeft) {
				visibleTrees++;
				continue;
			}
			// check horizontally (x), from right
			bool visibleRight = true;
			for (int x1 = x + 1; x1 < MAX_DIMENSION; x1++) {
				if (trees[y][x] <= trees[y][x1]) {
					visibleRight = false;
					break;
				}
			}
			if (visibleRight) {
				visibleTrees++;
				continue;
			}
			// check vertically (y), from above
			bool visibleTop = true;
			for (int y1 = 0; y1 < y; y1++) {
				if (trees[y][x] <= trees[y1][x]) {
					visibleTop = false;
					break;
				}
			}
			if (visibleTop) {
				visibleTrees++;
				continue;
			}
			// check vertically (y), from below
			bool visibleBottom = true;
			for (int y1 = y + 1; y1 < MAX_DIMENSION; y1++) {
				if (trees[y][x] <= trees[y1][x]) {
					visibleBottom = false;
					continue;
				}
			}
			if (visibleBottom) {
				visibleTrees++;
				continue;
			}
			// end of current
		}
	}
	printf("%d\n", MAX_DIMENSION * 4 - 4 + visibleTrees);

	// finish work
	fclose(input);
	return 0;
}
