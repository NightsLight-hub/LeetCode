#include "bq/blockingqueue.h"
#include "chrono"
#include "filesystem"
#include "thread/trythread.h"
#include "tool.h"
#include <iostream>

using namespace std;

class DataBlock {
public:
  DataBlock() {
    data = nullptr;
    length = 0;
  }
  ~DataBlock() { delete[] data; }
  DataBlock(int length, char *_data) : length(length) {
    data = new char[1500];
    memcpy(data, _data, length);
  }
  DataBlock(const DataBlock &d) {
    data = new char[1500];
    memcpy(data, d.data, d.length);
    length = d.length;
    printf("copy constructor\n");
  }
  int length;
  char *data;
};

void tryBlockingQueue() {
  BlockingQueue<shared_ptr<DataBlock>> queue;
  std::thread producer([&]() {
    int i = 0;
    ostringstream oss;
    while (true) {
      this_thread::sleep_for(chrono::milliseconds(1));
      oss.clear();
      oss.str("");
      oss << "hello " << i++;
      string d = oss.str();
      shared_ptr<DataBlock> dataBlock(
          new DataBlock((int)d.length(), oss.str().data()));
      queue.push(dataBlock);
    }
  });
  std::thread consumer([&]() {
    this_thread::sleep_for(chrono::milliseconds(1000));
    while (true) {
      shared_ptr<DataBlock> d;
      if (queue.tryWaitAndPop(d, 1)) {
        printf("%s \n", d->data);
      }
    }
  });
  producer.join();
}

int main() {
  tryBlockingQueue();
}
