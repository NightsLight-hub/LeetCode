//
// Created by sunxy on 2024/4/1.
//

#ifndef CLION_DEMO_DEMOLIB_H
#define CLION_DEMO_DEMOLIB_H

#ifdef __cplusplus
extern "C" {
#endif

extern void *getInstance();
extern void producer(BlockingQueue<int> *q);
extern int pop(BlockingQueue<int> *q);
extern void hello();

#ifdef __cplusplus
}
#endif

#endif // CLION_DEMO_DEMOLIB_H
