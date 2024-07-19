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
  bool operator<(const RtpData<T> &b) const // ע�⣬�˴���û��const��ᱨ��
  {
    return priority < b.priority; // ��value��ֵ�ɴ�С���У��γ�Node�Ĵ����
  }
  bool operator>(const RtpData<T> &b) const // ע�⣬�˴���û��const��ᱨ��
  {
    return priority > b.priority; // ��value��ֵ�ɴ�С���У��γ�Node�Ĵ����
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
    return *a > *b; // ����value��С��������
  }
};

// priorityԽС��Խ�緢��
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
      // ��ǰ����rtp����seq�� ����ͷ��rtp��seq С10000�� ˵��rtp seq
      // ��ת�ˣ�0 -- 1000��
      //      printf("insert to queue2, priority %lld \n", _data->priority);
      queue2.push(_data);
      return;
    }
    if (isRollOver && _data->priority > 60000) {
      // ˵����ʱ�շ���rtp seq��ת�� ������rtp����seq����60000
      // �շ��������л��� queue2 �϶��ǿյģ�����queue�
      // �����ȼ��ŵ����
      _data->priority -= threshold;
    }
    queue.push(_data);
  }

  bool empty() const { return queue.empty(); }

  int length() { return queue.size() + queue2.size(); }

  // pop ��֤������� queueû�а���queue2�а������
  shared_ptr<RtpData<T>> pop() {
    if (queue.empty() && queue2.empty()) {
      return nullptr;
    }
    if (queue.empty() && !queue2.empty()) {
      isRollOver = true;
      // queue2 �� seq��С�� 10000
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

  // ��queue2�Ķ����ƶ��� queue��
  void moveQueue() {
    while (!queue2.empty()) {
      queue.push(queue2.top());
      queue2.pop();
    }
  }
};

// priorityԽС��Խ�緢��
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

  // pop ��֤������� queue û�а���queue2�а������
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