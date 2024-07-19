#include "iostream"
#include "thread"
#include <gst/gst.h>
#include <malloc.h>
#include <string.h>

using namespace std;

const int sleepInterval = 3;

GMainLoop *main_loop;
GstElement *pipeline;

string compPipeline =
    "videotestsrc pattern=1 !"
    " video/x-raw,format=AYUV,framerate= (fraction "
    ")10/1,width=100,height=100 !  "
    "videobox border-alpha=0 top=-70 bottom=-70 right=-220 !  "
    "compositor name=comp sink_0::alpha=0.7 sink_1::alpha=0.5 !  "
    "videoconvert ! autovideosink  "
    "videotestsrc ! video/x-raw,format=AYUV,framerate= (fraction "
    ")5/1,width=320,height=240 ! comp.";

string testPipeline = "videotestsrc pattern=1 !"
                      " video/x-raw,format=AYUV,framerate= (fraction "
                      ")10/1,width=100,height=100 !  "
                      "videoconvert ! autovideosink  ";

/* This function is called when an error message is posted on the bus */
static void error_cb(GstBus *bus, GstMessage *msg, GstElement *data) {
  GError *err;
  gchar *debug_info;

  /* Print error details on the screen */
  gst_message_parse_error(msg, &err, &debug_info);
  g_printerr("Error received from element %s: %s\n", GST_OBJECT_NAME(msg->src),
             err->message);
  g_printerr("Debugging information: %s\n", debug_info ? debug_info : "none");
  g_clear_error(&err);
  g_free(debug_info);

  g_main_loop_quit(main_loop);
}

void entrypoint(int argc, char *argv[]) {
  cout << "run pipeline" << endl;
  pipeline = gst_parse_launch(compPipeline.c_str(), NULL);
  //  pipeline = gst_parse_launch(testPipeline.c_str(), NULL);

  //  GstBus *bus = gst_element_get_bus(pipeline);
  //  gst_bus_add_signal_watch(bus);
  //  g_signal_connect(G_OBJECT(bus), "message::error", (GCallback)error_cb,
  //                   pipeline);
  //  gst_object_unref(bus);

  /* Start playing the pipeline */
  gst_element_set_state(pipeline, GST_STATE_PLAYING);

  /* Create a GLib Main Loop and set it to run */
  main_loop = g_main_loop_new(NULL, FALSE);
  g_main_loop_run(main_loop);
}

int main(int argc, char *argv[]) {
  gst_init(&argc, &argv);
  cout << "start" << endl;
  int i = 0;
  while (true) {
    this_thread::sleep_for(chrono::seconds(sleepInterval));
    cout << "start " << ++i << "  test" << endl;
    thread([&] { entrypoint(argc, argv); }).detach();

    this_thread::sleep_for(chrono::seconds(sleepInterval*3));
    cout << "exit pipeline" << endl;

    /* Free resources */
    g_main_loop_quit(main_loop);
    g_main_loop_unref(main_loop);
    main_loop = nullptr;
    gst_element_set_state(pipeline, GST_STATE_NULL);
    gst_object_unref(pipeline);
    malloc_trim(0);
  }

  return 0;
}