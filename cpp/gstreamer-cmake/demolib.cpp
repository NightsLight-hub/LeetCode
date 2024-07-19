//
// Created by sunxy on 2024/4/1.
//

#include "bq/blockingqueue.h"
#include <cstring>
using namespace std;

extern "C" {
#include "demolib.h"
}

BlockingQueue<int> *bqInst = nullptr;

void *getQueueInstance() {
  if (bqInst == nullptr) {
    bqInst = new BlockingQueue<int>();
  }
  return bqInst;
}

unsigned int queueSize(void *q) {
  auto *queue1 = (BlockingQueue<int> *)q;
  return queue1->size();
}

void producer(void *q) {
  auto *queue1 = (BlockingQueue<int> *)q;
  for (int i = 0; i < 100; i++) {
    queue1->push(i);
  }
}

int pop(void *q) {
  int res;
  auto *queue1 = (BlockingQueue<int> *)q;
  if (queue1->tryPop(res)) {
    return res;
  }
  return -1;
}

void hello() { printf("I'm C, hello!\n"); }

void outputCharArray(void *str, int length) {
  char p[length + 1];
  memcpy(p, str, length);
  p[length] = 0;
//  printf("%s\n", p);
}

char* mkMsg() {
  char* msg = new char[1000000];
  memset(msg, 'a', 1000000);
  return msg;
}

void freeMsg(char* msg) {
  free(msg);
}
