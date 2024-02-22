//
// Created by sunxy on 2024/2/21.
//

#ifndef CLION_DEMO_TOOL_H
#define CLION_DEMO_TOOL_H
#include "random"

#include "iostream"
#include <condition_variable>
using namespace std;

int randomInt(int min, int max) {
  std::random_device
      rd; // ������õĻ�����һ��������������ϻ��һ�������������
  std::mt19937 gen(
      rd()); // gen��һ��ʹ��rd()�����ӳ�ʼ���ı�׼÷ɭ��ת�㷨�������������
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
