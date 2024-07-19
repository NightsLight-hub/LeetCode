//
// Created by sunxy on 2024/3/2.
//

#include "tool.h"
plog::ColorConsoleAppender<plog::TxtFormatter> consoleAppender;

void initLog(const string& logFilePath, const string& level) {
  if (level == "info") {
    plog::init(plog::info, logFilePath.c_str(), 50 * 1000 * 1000, 50);
  } else if (level == "debug") {
    plog::init(plog::debug, logFilePath.c_str(), 50 * 1000 * 1000, 50);
  } else if (level == "warning") {
    plog::init(plog::warning, logFilePath.c_str(), 50 * 1000 * 1000, 50);
  }
  plog::get()->addAppender(&consoleAppender);
  PLOG_INFO << "Logs path: " << getcwd(NULL, 0) << "/" << logFilePath;
}

bool makeDir(string folderPath) {
  if (0 != access(folderPath.c_str(), 0)) {
    int ret = mkdir(folderPath.c_str(), MODE);
    if (ret != 0) {
      cout << "Create directory failed: " << folderPath << endl;
      return false;
    }
  }
  return true;
}

int randomInt(int min, int max) {
  std::random_device
      rd; // 如果可用的话，从一个随机数发生器上获得一个真正的随机数
  std::mt19937 gen(
      rd()); // gen是一个使用rd()作种子初始化的标准梅森旋转算法的随机数发生器
  std::uniform_int_distribution<> distrib(min, max);

  return distrib(gen);
}

/**
 * get current timestamp milliseconds
 * @return
 */
long now() {
  std::chrono::time_point<std::chrono::system_clock> lastLevelChangedTime =
      chrono::system_clock::now();
  auto ms = chrono::time_point_cast<chrono::milliseconds>(lastLevelChangedTime);
  return ms.time_since_epoch().count();
}

