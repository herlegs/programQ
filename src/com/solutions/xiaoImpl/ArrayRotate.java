package com.solutions.xiaoImpl;

import com.util.Util;

import java.util.Arrays;

/**
 * Created by xiao on 14/11/16.
 */
public class ArrayRotate {
    public static void main(String[] args) {
        int[] array = new int[]{1,2,3,4,5,6,7,8,9,10};
        System.out.println(Arrays.toString(array));

        rotate(array, 7);
        System.out.println(Arrays.toString(array));
    }

    public static void rotate(int[] nums, int k){
        int n = nums.length;
        reverse(nums, 0, k-1);
        reverse(nums, k, n - 1);
        reverse(nums, 0, n - 1);
    }

    private static void reverse(int[] array, int start, int end){
        int mid = start + (end - start)/2;
        for (int i = start; i < mid; i++) {
            swap(array, i, start + end - i);
        }
    }

    private static void swap(int[] array, int i, int j){
        int tmp = array[i];
        array[i] = array[j];
        array[j] = tmp;
    }
}
