#define _THREADPOOL_H
#include<stdio.h>
#include<threadpool.h>
typedef struct ThreadPool ThreadPool;

//create and initialize
ThreadPool* threadPoolCreate(int min,int max,int queueSize){
    ThreadPool* pool=(ThreadPool*)malloc(sizeof(ThreadPool));
    do{
            if(pool==NULL){
            printf("malloc thread pool faill...\n")
            break;
        }
        pool->threadIDs=(pthread_t*)malloc(sizeof(pthread_t)*max);
        if(pool->threadIDs==NULL){
            printf("malloc threadIDs faill...\n");
            break;
        }
        memset(pool->threadIDs,0,sizeof(pthread_t)*max);
        pool->minNum=min;
        pool->maxNum=max;
        pool->busyNum=0;
        pool->liveNum=min;
        pool->exitNum=0;
        if( pthread_mutex_init(&pool->mutexPool,NULL)!=0||
            pthread_mutex_init(&pool->mutexBusy,NULL)!=0||
            pthread_mutex_init(&pool->notEmpty,NULL)!=0||
            pthread_mutex_init(&pool->notFull,NULL)!=0
            ){
                printf("mutex or condition init fail...\n");
                break;
            }
        pool->taskQ=(Task)malloc(sizeof(Task)*queueSize);
        pool->queueCap=queueSize;
        pool->queueSize=0;
        pool->queueFront=0;
        pool->queueRear=0;
        pool->shutdown=0;

        pthread_create(&pool->managerID,manager,NULL);
        for(int i=0;i<min;i++){
            pthread_create(&pool->threadIDs,NULL,worker,NULL);
        }
        return pool
    }while(0);
    if(pool->threadIDs)free(pool->threadIDs);
    if(pool->taskQ)free(pool->taskQ);
    if(pool)free(pool);
    return NULL;
};
void* worker(void* arg) {
    ThreadPool*pool=(ThreadPool*)arg;
    while(1){
        pthread_mutex_lock(&pool->mutexPool);
        while(pool->queueSize==0&&!pool->shutdown){
            pthread cond wait(&pool->notEmpty,&pool->mutexPool);
        }
        if(pool->shutdown){
            pthread_mutex_unlock(&pool->mutexPool);
            pthread_exit(NULL);
        }
        Task task;
        task.funct8ion=pool->taskQ[pool->queueFront].function;
        task.arg=pool->taskQ[pool->queueFront].arg;

        pthread_mutex_unlock(&pool->mutexPool);
    }
    return NULL;
}