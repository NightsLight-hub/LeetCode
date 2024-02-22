#include "algorithm"
#include "chrono"
#include "cstring"
#include "vector"
 #include "dmi.h"
#include <iostream>
using namespace std;

typedef signed char Byte;
void testTime();
bool testAnyOf();

int main() {
  testTime();
  cout << endl;
  cout << (testAnyOf() ? "true" : "false");
}

void testTime() {
  chrono::time_point<chrono::system_clock> lastLevelChangedTime =
      chrono::system_clock::now();
  auto ms = chrono::time_point_cast<chrono::milliseconds>(lastLevelChangedTime);
  long msV = ms.time_since_epoch().count();
  cout << msV;
}

void testMemMove() {
  Byte a1[5] = {1, 2, 3, 4, 5};
  memmove(a1 + 1, a1, 4 * sizeof(Byte));
  for (Byte i : a1) {
    std::cout << (int)i;
  }
}

bool testAnyOf() {
  vector<int> a1;
  a1.push_back(1);
  a1.push_back(2);
  a1.push_back(3);
  return any_of(a1.begin(), a1.end(), [](int element) -> bool { return element == 1; });
}