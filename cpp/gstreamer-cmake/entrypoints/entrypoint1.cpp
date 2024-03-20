#include "bq/blockingqueue.h"
#include "chrono"
#include "filesystem"
#include "jitterbuffer/jitterbuffer.h"
#include "model.h"
#include "thread/trythread.h"
#include "tool.h"
#include <iostream>

using namespace std;
void tryBlockingQueue();

void tryJitterBuffer() {
  mutex lock;
  JitterBuffer<DataBlock> jitterBuffer;
  std::thread producer([&]() {
    std::srand(std::time(nullptr));
    unsigned short int i = 65000;
    while (true) {
      this_thread::sleep_for(chrono::milliseconds(3));
      std::unique_lock<std::mutex> locker(lock);
      if (rand() % 2 == 0) {
        DataBlock dataBlock(5, "hello", i + 1);
        shared_ptr<RtpData<DataBlock>> sptr =
            make_shared<RtpData<DataBlock>>(RtpData(dataBlock, i + 1));
        jitterBuffer.push(sptr);
        DataBlock dataBlock2(5, "hello", i);
        sptr = make_shared<RtpData<DataBlock>>(RtpData(dataBlock2, i));
        jitterBuffer.push(sptr);
        i += 2;
      } else {
        DataBlock dataBlock(5, "hello", i);
        auto sptr = make_shared<RtpData<DataBlock>>(RtpData(dataBlock, i));
        i++;
        jitterBuffer.push(sptr);
      }
    }
  });
  std::thread consumer([&]() {
    this_thread::sleep_for(chrono::milliseconds(1000));
    printf("length %d \n", jitterBuffer.length());
    int lastSentSeq = 0;
    while (true) {
      std::unique_lock<std::mutex> locker(lock);
      if (jitterBuffer.length() < 20) {
        continue;
      }
      shared_ptr<RtpData<DataBlock>> d = jitterBuffer.pop();
      if (d != nullptr) {
        if (lastSentSeq != 0 && d->content.seq != lastSentSeq + 1) {
          printf("wrong order packet, expect %d, but got %d \n",
                 lastSentSeq + 1, d->content.seq);
        }
        lastSentSeq = d->content.seq;
      } else {
        continue;
      }
    }
  });
  producer.detach();
  consumer.join();
};

void tryCopy() {
  cout << "make DataBlock" << endl;
  DataBlock d(5, "hello", 1);
  cout << "make RtpData" << endl;
  RtpData<DataBlock> rtp = RtpData(d, 1);
  cout << rtp.content.data << endl;
}

int main() {
  std::setbuf(stdin, 0);
  tryJitterBuffer();
}

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
          new DataBlock((int)d.length(), oss.str().data(), 1));
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
