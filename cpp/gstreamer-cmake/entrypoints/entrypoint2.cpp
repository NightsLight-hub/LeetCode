#include "cq/blockingconcurrentqueue.h"
#include "cq/eventbus.h"
#include "dmi_base.h"
#include "tool.h"

using namespace std;

class Data {
public:
  string name;
  string value;
};

class producer {
public:
  producer(EventBus<Data> &bus) : bus(bus) {}

  EventBus<Data> &bus;
  void produce() {
    for (int i = 0; i != 100; ++i) {
      std::this_thread::sleep_for(std::chrono::milliseconds(100));
      Data data;
      data.name = "name" + to_string(i);
      data.value = "name" + to_string(i);
      bus.publish(data);
    }
  }
};

class consumer {
public:
  consumer() = default;

  void consume(Data &data) {
    printf("consume data: name: %s value: %s\n", data.name.c_str(),
           data.value.c_str());
  }
};

int main() {
  EventBus<Data> eBus;
  producer p(eBus);
  consumer consumer;
  eBus.addListener("c1", [&consumer](Data data) { consumer.consume(data); });
  eBus.start();

  thread p1(&producer::produce, &p);
  p1.join();

  eBus.removeListener("c1");
  eBus.stop();
  this_thread::sleep_for(chrono::milliseconds(1000));
}

void testBQ() {
  moodycamel::BlockingConcurrentQueue<int> q;
  std::thread producer([&]() {
    for (int i = 0; i != 100; ++i) {
      std::this_thread::sleep_for(std::chrono::milliseconds(i % 10));
      q.enqueue(i);
    }
  });
  std::thread consumer([&]() {
    for (int i = 0; i != 100; ++i) {
      int item;
      q.wait_dequeue(item);
      assert(item == i);

      if (q.wait_dequeue_timed(item, std::chrono::milliseconds(5))) {
        ++i;
        assert(item == i);
      }
    }
  });
  producer.join();
  consumer.join();

  assert(q.size_approx() == 0);
}