#include "dmi_base.h"
#include "hiredis/hiredis.h"
#include "tool.h"
#include <sw/redis++/redis++.h>

using namespace std;
// using namespace sw::redis;

int main() {
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
