package com.solutions.xiaoImpl;

import com.questions.QuickSort;

import java.util.Arrays;

/**
 * Created by xiao on 21/6/16.
 */
public class QuickSortImpl implements QuickSort{
    public static void main(String[] args){
        int[] test = null;
        test = new int[]{3,4,1,7,0,10,5,8};
        insertionSort(test, 0, test.length - 1);
        System.out.println(Arrays.toString(test));
    }

    private static void insertionSort(int[] array, int start, int end){
        if(start >= end || start < 0 || end >= array.length){
            return;
        }
        for (int i = start + 1; i <= end; i++) {
            int cur = array[i];
            int j = i - 1;
//            for(; j >= 0; j--){
//                if(array[j] > cur){
//                    array[j + 1] = array[j];
//                }
//                else{
//                    break;
//                }
//            }
            while(j >= 0 && array[j] > cur){
                array[j+1] = array[j];
                j--;
            }
            array[j+1] = cur;
        }
    }
}
