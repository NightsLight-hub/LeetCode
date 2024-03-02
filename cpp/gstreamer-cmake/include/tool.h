//
// Created by sunxy on 2024/2/21.
//

#ifndef CLION_DEMO_TOOL_H
#define CLION_DEMO_TOOL_H
#include "random"
#include "iostream"
#include <condition_variable>
#include "plog/log.h"
#include <plog/Initializers/RollingFileInitializer.h>
#include <plog/Formatters/TxtFormatter.h>
#include <plog/Appenders/ColorConsoleAppender.h>

using namespace std;

#define MODE (S_IRWXU | S_IRWXG | S_IRWXO)
extern plog::ColorConsoleAppender<plog::TxtFormatter> consoleAppender;
void initLog(const string& logFilePath, const string& level);

bool makeDir(string folderPath);

int randomInt(int min, int max);

/**
 * get current timestamp milliseconds
 * @return
 */
long now();

class CountDownLatch {
private:
  mutex cvMtx;
  condition_variable cv;
  int count;

public:
  explicit CountDownLatch(int c) : count(c){};
  void await() {
    std::unique_lock<std::mutex> locker(cvMtx);
    cv.wait(locker, [this] { return this->count <= 0; });
  }

  void countDown() {
    std::unique_lock<std::mutex> locker(cvMtx);
    count--;
    cv.notify_all();
  }

  int getCount() {
    std::unique_lock<std::mutex> locker(cvMtx);
    return count;
  }
};

#endif // CLION_DEMO_TOOL_H
