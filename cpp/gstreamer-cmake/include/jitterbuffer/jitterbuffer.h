#pragma once

#include <atomic>
#include <condition_variable>
#include <mutex>
#include <queue>
#include <vector>
using namespace std;

template <typename T> class RtpData {
public:
  RtpData(T &data, long long priority) : content(data), priority(priority){};
  ~RtpData() = default;
  T content;
  long long priority;
  bool operator<(const RtpData<T> &b) const // 注意，此处若没有const则会报错
  {
    return priority < b.priority; // 将value的值由大到小排列，形成Node的大根堆
  }
  bool operator>(const RtpData<T> &b) const // 注意，此处若没有const则会报错
  {
    return priority > b.priority; // 将value的值由大到小排列，形成Node的大根堆
  }
  //  static bool compare(const RtpData<T> &a, const RtpData<T> &b) {
  //    return a.priority > b.priority;
  //  }
};

// return true, b in front of a
// return false, a in front of b
template <typename T> struct cmp {
  bool operator()(const shared_ptr<RtpData<T>> &a,
                  const shared_ptr<RtpData<T>> &b) {
    return *a > *b; // 按照value从小到大排列
  }
};

// priority越小，越早发送
template <typename T> class JitterBuffer {
public:
  JitterBuffer() : isRollOver(false){};

  void push(shared_ptr<RtpData<T>> const &_data) {
    if (queue.empty()) {
      queue.push(_data);
      return;
    }
    if (isRollOver && _data->priority >= 30000 && _data->priority <= 35000) {
      isRollOver = false;
    }
    //    printf("priority is %lld, top priority is %lld  \n", _data->priority,
    //    queue.top()->priority);
    if (_data->priority + 10000 < queue.top()->priority) {
      // 当前来的rtp包的seq比 队列头的rtp包seq 小10000， 说明rtp seq
      // 翻转了（0 -- 1000）
      //      printf("insert to queue2, priority %lld \n", _data->priority);
      queue2.push(_data);
      return;
    }
    if (isRollOver && _data->priority > 60000) {
      // 说明此时刚发生rtp seq反转， 新来的rtp包的seq大于60000
      // 刚发生队列切换， queue2 肯定是空的，都在queue里。
      // 把优先级放到最大。
      _data->priority -= threshold;
    }
    queue.push(_data);
  }

  bool empty() const { return queue.empty(); }

  int length() { return queue.size() + queue2.size(); }

  // pop 保证不会出现 queue没有包，queue2有包的情况
  shared_ptr<RtpData<T>> pop() {
    if (queue.empty() && queue2.empty()) {
      return nullptr;
    }
    if (queue.empty() && !queue2.empty()) {
      isRollOver = true;
      // queue2 的 seq都小于 10000
      moveQueue();
    }

    shared_ptr<RtpData<T>> value = queue.top();
    queue.pop();

    if (queue.empty() && !queue2.empty()) {
      isRollOver = true;
      moveQueue();
    }
    return value;
  }

  shared_ptr<RtpData<T>> top() {
    if (queue.empty() && queue2.empty()) {
      return nullptr;
    }
    if (queue.empty() && !queue2.empty()) {
      return queue2.top();
    }
    return queue.top();
  }

private:
  std::priority_queue<shared_ptr<RtpData<T>>, vector<shared_ptr<RtpData<T>>>,
                      cmp<T>>
      queue;
  std::priority_queue<shared_ptr<RtpData<T>>, vector<shared_ptr<RtpData<T>>>,
                      cmp<T>>
      queue2;
  int threshold = 100000;
  atomic_bool isRollOver;

  // 将queue2的东西移动到 queue中
  void moveQueue() {
    while (!queue2.empty()) {
      queue.push(queue2.top());
      queue2.pop();
    }
  }
};

// priority越小，越早发送
template <typename T> class JitterBuffer2 {
public:
  JitterBuffer2(){};

  void push(shared_ptr<RtpData<T>> const &_data) {
    if (_data->priority < 1000) {
      queue2.push(_data);
    } else {
      queue.push(_data);
    }
  }

  bool empty() const { return queue.empty() && queue2.empty(); }

  int length() { return queue.size() + queue2.size(); }

  // pop 保证不会出现 queue 没有包，queue2有包的情况
  shared_ptr<RtpData<T>> pop() {
    if (queue.empty() && queue2.empty()) {
      return nullptr;
    }
    if (queue.empty() && !queue2.empty()) {
      shared_ptr<RtpData<T>> value = queue2.top();
      queue2.pop();
      return value;
    }
    if (!queue.empty() && queue2.empty()) {
      shared_ptr<RtpData<T>> value = queue.top();
      queue.pop();
      return value;
    }
    if (queue.top()->priority < threshold) {
      shared_ptr<RtpData<T>> value = queue2.top();
      queue2.pop();
      return value;
    } else {
      shared_ptr<RtpData<T>> value = queue.top();
      queue.pop();
      return value;
    }
  }

  shared_ptr<RtpData<T>> top() {
    if (queue.empty() && queue2.empty()) {
      return nullptr;
    }
    if (queue.empty() && !queue2.empty()) {
      return queue2.top();
    }
    if (queue.top()->priority < threshold) {
      return queue2.top();
    } else {
      return queue.top();
    }
  }

private:
  std::priority_queue<shared_ptr<RtpData<T>>, vector<shared_ptr<RtpData<T>>>,
                      cmp<T>>
      queue;
  std::priority_queue<shared_ptr<RtpData<T>>, vector<shared_ptr<RtpData<T>>>,
                      cmp<T>>
      queue2;
  int threshold = 30000;
};