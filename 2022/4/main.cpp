#include <cstdlib>
#include <fstream>
#include <string>
// #include <vector>

template <typename T>

std::pair<T, T> splitTwo(T str, char c) {
	auto split = str.find(c);
	return {str.substr(0, split), str.substr(split + 1)};
}

int main() {
	// input
	std::ifstream file("input");
	if (!file.is_open()) {
		printf("failed to open input file\n");
		return 1;
	}

	// hybrid: input + logic
	int sum1 = 0, sum2 = 0;
	std::string buffer;

	while (file.good() && std::getline(file, buffer)) {
		auto elves = splitTwo(buffer, ',');

		auto rangeA = splitTwo(elves.first, '-');
		auto rangeB = splitTwo(elves.second, '-');
		std::pair<std::pair<int, int>, std::pair<int, int>> fdm{
			{atoi(rangeA.first.c_str()), atoi(rangeA.second.c_str())},
			{atoi(rangeB.first.c_str()), atoi(rangeB.second.c_str())},
		};

		// std::vector<std::pair<int, int>> fdm; // a's and b's range
		// for (const auto &e : std::vector<std::string>{elves.first, elves.second}) {
		// auto range = splitTwo(e, '-');
		// fdm.push_back({atoi(range.first.c_str()), atoi(range.second.c_str())});
		//}

		// if ((fdm[0].first >= fdm[1].first && fdm[0].second <= fdm[1].second) ||
		//(fdm[1].first >= fdm[0].first && fdm[1].second <= fdm[0].second)) {
		// sum1++;
		//}
		// if ((fdm[0].first <= fdm[1].second && fdm[1].first <= fdm[0].second) ||
		//(fdm[1].second <= fdm[0].first && fdm[0].first <= fdm[1].second)) {
		// sum2++;
		//}

		if ((fdm.first.first >= fdm.second.first && fdm.first.second <= fdm.second.second) ||
			(fdm.second.first >= fdm.first.first && fdm.second.second <= fdm.first.second)) {
			sum1++;
		}
		if ((fdm.first.first <= fdm.second.second && fdm.second.first <= fdm.first.second) ||
			(fdm.second.second <= fdm.first.first && fdm.first.first <= fdm.second.second)) {
			sum2++;
		}
	}

	printf("%d\n%d\n", sum1, sum2);
	return 0;
}
