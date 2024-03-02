//
// Created by sunxy on 2024/3/2.
//

#pragma once
#include "thread"
#include "atomic"
class TryThread{
public:
    TryThread():shouldStop(false), isRunning(false) {};
    ~TryThread() = default;
    void start();
    void stop();

  private:
    std::thread t;
    std::atomic_bool shouldStop;
    std::atomic_bool isRunning;
};
