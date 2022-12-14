#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

int x = 1, cycle = 0, strengths = 0;
int pixel = 0; // XXX: when you get home, try to remove this variable

void processCycle() {
	++cycle;
	++pixel;
	// printf("debug: #%d: i = %d; x = %d-%d\n", cycle, pixel, x, x + 2);
	// part 1
	if ((cycle + 20) % 40 == 0) {
		strengths += x * cycle;
	}

	// part 2
	if (x <= pixel && pixel <= x + 2) {
		printf("%c", '#');
	} else {
		printf("%c", ' ');
	}
	if (pixel % 40 == 0) {
		printf("\n"); // print a newline after every 40 characters have been printed
		pixel = 0;
	}
}

char buffer[10];
int main() {
	FILE *input = fopen("input", "r");
	if (!input) {
		printf("failed to open input file!\n");
		return 1;
	}

	clock_t start = clock();
	while (!feof(input) && fgets(buffer, 10, input) != NULL) {
		if (strncmp(buffer, "noop", 4) == 0) {
			processCycle();
		} else if (strncmp(buffer, "addx", 4) == 0) {
			char *space = strchr(buffer, ' ');
			int amount = atoi(space + 1);
			for (int i = 0; i < 2; i++) {
				processCycle();
			}
			x += amount;
		} else {
			printf("fatal: undefined instruction \"%s\"!\n", buffer);
			return 1;
		}
	}
	printf("\n%d\n", strengths); // part 1
	clock_t end = clock();

	printf("elapsed time: %ld microseconds\n", (end - start));
	return 0;
}
