package com.solutions.xiaoImpl;

import java.util.Arrays;

/**
 * Created by xiao on 15/8/16.
 */
public class BinaryIndexTree {
    public static void main(String[] args){
        buildBIT1D(new int[]{3,5,-6,4,4,9});
        
        System.out.println(Arrays.toString(bit));
        
        System.out.println(getRangeSum(0,2));
        System.out.println(getRangeSum(1,4));
        update(1,10);
        System.out.println(getRangeSum(0,2));
        System.out.println(getRangeSum(1,4));
    }

    private static int[] origin;
    private static int[] bit;
    //size of binary index tree array size
    private static int size;


    public static void buildBIT1D(int[] array){
        int n = array.length;
        size = n + 1;
        bit = new int[size];
        origin = new int[n];
        for(int i = 0; i < n; i++){
            update(i, array[i]);
        }
    }

    public static void update(int index, int val){
        int i = index + 1;
        int diff = val - origin[index];
        origin[index] = val;
        while(i < size){
            bit[i] += diff;
            i = i + lsb(i);
        }
    }

    //from and to: index in original array
    public static int getRangeSum(int from, int to){
        return getSum(to + 1) - getSum(from);
    }

    //i: num of elements
    public static int getSum(int i){
        int sum = 0;
        while(i > 0){
            sum += bit[i];
            i = i - lsb(i);
        }
        return sum;
    }

    //get least significant bit
    private static int lsb(int n){
        return n & (-n);
    }


}
