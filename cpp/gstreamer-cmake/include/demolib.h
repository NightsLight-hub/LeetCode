//
// Created by sunxy on 2024/4/1.
//

#ifndef CLION_DEMO_DEMOLIB_H
#define CLION_DEMO_DEMOLIB_H

#ifdef __cplusplus
extern "C" {
#endif


extern void *getQueueInstance();
extern unsigned int queueSize(void *q);
extern void producer(void* q);
extern int pop(void* q);
extern void hello();
extern void outputCharArray(void* str, int length);

#ifdef __cplusplus
}
#endif

#endif // CLION_DEMO_DEMOLIB_H
