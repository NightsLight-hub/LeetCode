//
// Created by sunxy on 2024/3/2.
//

#include "thread/trythread.h"
#include "iostream"
#include <tool.h>

using namespace std;
void TryThread::start() {
  t = thread([this] {
    PLOG_INFO << "thread started, id is " << this_thread::get_id() << endl;
    while (shouldStop) {
      this_thread::sleep_for(chrono::microseconds(100));
    }
  });
  isRunning = true;
  PLOG_INFO << "outside thread started,id is " << t.get_id() << endl;
}

void TryThread::stop() {
  shouldStop = true;
  t.join();
  isRunning = false;
}
