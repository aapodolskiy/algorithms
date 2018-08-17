#include <iostream>
#include <vector>

void insertionSort(std::vector<int> &v);

template <typename T>
void outputVector(const std::vector<T> &v) { for (int i = 0; i < v.size(); ++i) std::cout << v[i] << (i != (v.size() - 1) ? ", " : "\n"

int main() {
  int n = 0;
  std::cout << "Enter a positive integer number N: ";
  std::cin >> n;

  if (n <= 0) {
    std::cout << "Your number is not positive";
  }
  else {
    std::cout << "Enter N integer numbers: ";
    std::vector<int> v;
    int k = 0;
    for (int i = 0; i < n; ++i) {
      std::cin >> k;
      v.push_back(k);
    }

    int k = linearSearch(A, v);

    std::cout << "Here is your sequence:\n";

    outputVector(v);
  }
  return 0;
}



