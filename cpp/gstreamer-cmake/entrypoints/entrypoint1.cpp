#include "algorithm"
#include "chrono"
#include "cq/eventbus.h"
#include "cstring"
#include "dmi.h"
#include "filesystem"
#include "thread/trythread.h"
#include "tool.h"
#include "vector"
#include <iostream>
using namespace std;

int main() {
  string cdir = std::filesystem::current_path();
  makeDir(cdir);
  initLog("log.txt", "info");
  TryThread tt;
  tt.start();
  this_thread::sleep_for(chrono::milliseconds(1000));
  tt.stop();

}
