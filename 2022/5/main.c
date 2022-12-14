#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_BUFFER 40
#define STACK_COUNT 9
#define MAX_INNER_STACK 50

void push_front(char stacks[STACK_COUNT][MAX_INNER_STACK], int stack, char value) {
	// find the last item of the stack
	size_t index = strlen(stacks[stack]);

	// shift all the items to the right to make room for new value
	for (int i = index; i > 0; i--) {
		stacks[stack][i] = stacks[stack][i - 1];
	}
	// finally, set given value as the first value in the front of the stack
	stacks[stack][0] = value;
}

// pop amount values from senderStack and push all of them in reverse order to destinationStack
// bulk_move means we're grabbing and moving multiple crates at once with our crane
void move_many(char stacks[STACK_COUNT][MAX_INNER_STACK], int sender_stack, int destination_stack,
			   int amount, bool bulk_move) {
	size_t sender_i = strlen(stacks[sender_stack]);
	size_t destination_i = strlen(stacks[destination_stack]);

	// move items to destinationStack here to save further iterations
	/*printf("debug: sender_i = %lu, destination_i = %lu\n", sender_i, destination_i);*/
	for (int i = 0; i < amount; i++) {
		int pop_at = sender_i - 1 - i;
		if (bulk_move) {
			pop_at = sender_i - amount + i;
		}

		/*printf("debug2: sender[%d] -> destination[%lu]\n", pop_at, destination_i + i);*/
		char popped = stacks[sender_stack][pop_at];
		stacks[sender_stack][pop_at] = '\0';
		stacks[destination_stack][destination_i + i] = popped;
	}
}

// helper function for debugging, courtesy of ChatGPT <3
void print_stacks(char stacks[STACK_COUNT][MAX_INNER_STACK]) {
	// iterate over each stack
	for (int i = 0; i < STACK_COUNT; i++) {
		// print the stack header
		printf("stack #%d:\n", i + 1);

		// iterate over the elements in the stack
		for (int j = 0; j < MAX_INNER_STACK; j++) {
			// print the element
			printf("%4c ", stacks[i][j]);

			// print a newline after every 10 elements
			if ((j + 1) % 10 == 0) {
				printf("\n");
			}
		}

		// print a newline after each stack
		printf("\n");
	}
}

// declare stacks; it's important to initialize all values or else might get garbage data
// if it would've been done inside main, it could be filled with garbage
// in that case, either set its value to {{0}} or use memset(stacks, 0, x*y);
char stacks[STACK_COUNT][MAX_INNER_STACK];

int main() {
	FILE *p_input = fopen("input", "r");

	if (!p_input) {
		perror("failed to read input file");
		exit(1);
	}

	// handle logic while reading input
	char buffer[MAX_BUFFER];
	bool crate_part_done = false;
	int line = 1; // used for crate part
	while (!feof(p_input) && fgets(buffer, MAX_BUFFER, p_input) != NULL) {
		// handle crate stacks... this is pain
		if (!crate_part_done) {
			if (*buffer == '\n') {
				// printf("crate part is done!\n");
				/*print_stacks(stacks);*/
				crate_part_done = true;
				continue;
			}

			size_t line_len = strlen(buffer); // paj mentioned it's better to do it outside
			for (int i = 0; i < line_len; i++) {
				if (buffer[i] == ' ') {
					continue;
				}

				if (buffer[i] == '[') {
					int stackIndex = i / 4;
					printf("found start of crate '%c' from stack %d at line %d:%d\n", buffer[i + 1],
						   stackIndex, line, i);
					push_front(stacks, stackIndex, buffer[i + 1]);
				}
			}

			line++;
			continue; // get another line, don't go to 'handle instructions' part yet
		}

		// handle crate move instructions
		int amount, sender_stack, destination_stack;
		sscanf(buffer, "move %d from %d to %d", &amount, &sender_stack, &destination_stack);

		// part 1: move crates one by one - similar to how regular stack in programming works
		move_many(stacks, sender_stack - 1, destination_stack - 1, amount, false);
		// part 2: account for ability to move multiple crates at once
		/*move_many(stacks, sender_stack - 1, destination_stack - 1, amount, true);*/
	}

	for (int i = 0; i < STACK_COUNT; i++) {
		printf("%c", stacks[i][strlen(stacks[i]) - 1]);
	}
	printf("\n");

	fclose(p_input);
	return 0;
}
