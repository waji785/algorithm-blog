package cn.wj.implementation;
import java.util.ArrayList;
import java.util.Arrays;

import static javafx.scene.input.KeyCode.T;

public class stackArray {
    int[] stack;
    int maxSize;
    int top;
    public stackArray{(int maxSize)
        top = -1;
        this.maxSize= maxSize;
        stack = new int[maxSize];
    }
    //判断空栈
    public boolean isEmpty(){
        return top == -1;
    }
    //判断栈满
    public  boolean isFull(){
        return top+1 == maxSize;
    }
    //入栈
    public void push(int data){
        if (isFull()){
            throw new RuntimeException("栈满");
        }
        top++;
        stack[top] = data;

    }
    //出栈
    public int pop() {
        if (isEmpty()) {
            throw new RuntimeException("栈空");
        }else return stack[top--];
    }
    //扩容
    public void expandCapacity(int maxSize){
        int len = stack.length;
        if(maxSize>len){
            maxSize = maxSize*2;
            stack = Arrays.copyOf(stack,maxSize);
        }
    }
    //遍历
    public  void showStack(){
        if(isEmpty()){
            System.out.println("empty");
        }else
            for(int i = top; i>=0;i--){
                System.out.println(stack[i]);
            }
    }
}
