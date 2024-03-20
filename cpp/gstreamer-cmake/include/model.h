//
// Created by sunxy on 2024/3/20.
//

#ifndef CLION_DEMO_MODEL_H
#define CLION_DEMO_MODEL_H

#include "cstdlib"
#include "cstring"
#include "iostream"

class DataBlock {
public:
  DataBlock() : length(0), seq(0), data(nullptr) {}
  ~DataBlock() {
//    printf("DataBlock destructor \n");
    delete[] data;
  }
  DataBlock(int length, char *_data, int seq) : length(length), seq(seq) {
    data = new char[1500];
    memcpy(data, _data, length);
  }
  DataBlock(const DataBlock &d) {
//    printf("DataBlock copy constructor \n");
    data = new char[1500];
    memcpy(data, d.data, d.length);
    length = d.length;
    seq = d.seq;
  }
  int length;
  int seq;
  char *data;
};

#endif // CLION_DEMO_MODEL_H
