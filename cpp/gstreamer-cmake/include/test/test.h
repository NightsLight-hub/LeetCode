//
// Created by sunxy on 2024/3/2.
//
#pragma once
#include "algorithm"
#include "chrono"
#include "cstring"
#include "dmi.h"
#include "tool.h"
#include "vector"
#include <iostream>
#include "hiredis/hiredis.h"
#include <sw/redis++/redis++.h>
#include "cq/eventbus.h"

using namespace std;
class DemoListener : public ActiveSpeakerChangedListener {
public:
  void activeSpeakerChanged(string id, int energyScore) override {
    cout << "============= activeSpeakerChanged  id: " << id
         << " energyScore: " << energyScore << "=============" << endl;
  }
};

void levelChangeWorker(DominantSpeakerIdentification *dmi, string id,
                       int periodCount, int packetInterval,
                       const vector<Byte> &levels) {
  for (Byte level : levels) {
    int i = 0;
    printf("%s level: %d\n", id.c_str(), (int)level);
    while (i++ < periodCount) {
      dmi->levelChanged(id, level);
      this_thread::sleep_for(chrono::milliseconds(packetInterval));
    }
  }
}

void testLevelChanged1() {
  auto dmi = new DominantSpeakerIdentification(-1);
  auto listener1 = new DemoListener();
  dmi->addActiveSpeakerChangedListener(listener1);
  CountDownLatch latch(3);
  //  changeDmiAudioLevel(dmi, sxy, 100);
  vector<Byte> v1 = vector<Byte>{
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
  };
  vector<Byte> v2 = vector<Byte>{
      0,
      0,
      0,
      0,
      (Byte)randomInt(60, 127),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      (Byte)randomInt(60, 100),
      0,
      0,
      0,
      0,
      0,
      0,
      0,
  };
  vector<Byte> v3 = vector<Byte>{
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      (Byte)randomInt(100, 127),
      (Byte)randomInt(100, 127),
      (Byte)randomInt(80, 100),
      (Byte)randomInt(80, 100),
      (Byte)randomInt(40, 60),
      (Byte)randomInt(80, 100),
      (Byte)randomInt(80, 100),
  };
  thread *t1 = new thread([&] {
    levelChangeWorker(dmi, "sxy", 50, 20, v1);
    latch.countDown();
  });
  t1->detach();
  thread *t2 = new thread([&] {
    levelChangeWorker(dmi, "hzl", 50, 20, v2);
    latch.countDown();
  });
  t2->detach();
  thread *t3 = new thread([&] {
    levelChangeWorker(dmi, "hbz", 50, 20, v3);
    cout << "t3 finished" << endl;
    latch.countDown();
  });
  t3->detach();
  latch.await();
}

void testDmiAddAndRemoveListener() {
  auto dmi = new DominantSpeakerIdentification(-1);
  auto listener1 = new DemoListener();
  cout << "listener1 address " << listener1 << endl;
  auto listener2 = new DemoListener();
  cout << "listener2 address " << listener2 << endl;
  dmi->addActiveSpeakerChangedListener(listener1);
  dmi->addActiveSpeakerChangedListener(listener2);
  dmi->outputListeners();
  dmi->removeActiveSpeakerChangedListener(listener1);
  dmi->outputListeners();
  dmi->removeActiveSpeakerChangedListener(listener2);
  dmi->outputListeners();
  delete listener1;
  delete listener2;
  delete dmi;
}

void testMemMove() {
  Byte a1[5] = {1, 2, 3, 4, 5};
  memmove(a1 + 1, a1, 4 * sizeof(Byte));
  for (Byte i : a1) {
    std::cout << (int)i;
  }
}

bool testAnyOf() {
  vector<int> a1;
  a1.push_back(1);
  a1.push_back(2);
  a1.push_back(3);
  return any_of(a1.begin(), a1.end(),
                [](int element) -> bool { return element == 1; });
}

void testRedis(){
  try {
    setbuf(stdout, NULL);
    sw::redis::ConnectionOptions opts1;
    opts1.host = "172.18.23.184";
    opts1.port = 6379;
    auto redis1 = sw::redis::Redis(opts1);

    // sub1's socket_timeout is 100ms.
    auto subscriber = redis1.subscriber();

    // Subscribe to channels and patterns.
    subscriber.on_message([](std::string channel, std::string msg) {
      printf("get message from %s\n", channel.c_str());
      printf("%s\n", msg.c_str());
      printf("------------------");
    });

    vector<string > channels = {"StartForwardConferenceAVTopic", "StopForwardConferenceAVTopic", "EndpointStatusMessageTopic", "ConferenceStatusTopic"};
    std::for_each(channels.begin(), channels.end(), [&](const auto &item) {
      subscriber.subscribe(item);
    });

    while (true) {
      try {
        subscriber.consume();
      } catch (const sw::redis::Error &err) {
        // Handle exceptions.
      }
    }

  } catch (const sw::redis::Error &e) {
    // Error handling.
  }
}