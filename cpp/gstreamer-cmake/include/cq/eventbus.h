
//
// Created by sunxy on 2024/2/29.  cq  不保证交付顺序，音视频场景下不适用
//

#ifndef CLION_DEMO_EVENTBUS_H
#define CLION_DEMO_EVENTBUS_H
#include "blockingconcurrentqueue.h"
#include "map"
#include "vector"
#include <cstdlib>
#include <cstring>
#include <functional>
#include <iostream>
using namespace std;

template <typename Event> class EventBus {
public:
  EventBus() : shouldStop(false){};
  ~EventBus() = default;
  void publish(const Event &event) { q.enqueue(event); }
  void addListener(string listenerId,
                   function<void(const Event &)> const &listener) {
    listeners[listenerId] = listener;
  }
  void removeListener(string listenerId) {
    listeners.erase(listenerId);
  }

  void start() { thread(&EventBus::run, this).detach(); }
  void stop() { shouldStop.store(true); }

private:
  moodycamel::BlockingConcurrentQueue<Event> q;
  atomic_bool shouldStop;
  map<string, function<void(const Event &)>> listeners;

  void run() {
    printf("eventbus start...\n");
    while (!shouldStop.load()) {
      Event e;
      if (q.wait_dequeue_timed(e, std::chrono::milliseconds(5))) {
        for (auto l : listeners) {
          l.second(e);
        }
      }
    }
    printf("eventbus stop...\n");
  }
};

#endif // CLION_DEMO_EVENTBUS_H
