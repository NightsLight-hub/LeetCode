//
// Created by sunxy on 2024/4/1.
//

#include "bq/blockingqueue.h"

extern  "C" {
  #include "demolib.h"
}

BlockingQueue<int> *queue = nullptr;

void* getInstance() {
  if (queue == nullptr){
    queue = new BlockingQueue<int>();
  }
  return queue;
}

void producer(BlockingQueue<int>* q) {
  for (int i = 0; i < 100; i++) {
    q->push(i);
  }
}

int pop(BlockingQueue<int>* q) {
  int res;
  if (q->tryPop(res)) {
    return res;
  }
  return -1;
}

void hello() {
  printf("I'm C, hello!\n");
}