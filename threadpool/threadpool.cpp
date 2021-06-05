#include "threadpool.h"
#include <pthread.h>
//task struct
typedef struct Task{
    void (*func)(void* arg);
    void* arg;
}Task;
//threadpool struct
struct ThreadPool{
    Task* taskQ;
    //task nums
    int queueSize;
    int queueCap;
    //front
    int queueFront;
    //rear
    int queueRear;

    //managerID
    pthread_t managerID;
    //workerID
    pthread_t *threadIDs;
    //min thread nums
    int minNum;
    //max thread nums
    int maxNum;
    //busy thread nums
    int busyNum;
    //live thread nums
    int liveNum;
    //thread that need to be destory
    int exitNum;
    //threadpool mutex
    pthread_mutex_t mutexPool;
    //busyNum mytex
    pthread_mutex_t mutexBusy;
    //queue is full
    pthread_cond_t notFull;
    //queue is empty
    pthread_cond_t notEmpty;

    //1=destory threadpool,0=keep
    int shutdown;
};