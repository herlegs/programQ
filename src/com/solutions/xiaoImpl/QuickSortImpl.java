package com.solutions.xiaoImpl;

import com.questions.QuickSort;
import com.util.Util;

import java.util.Arrays;

/**
 * Created by xiao on 21/6/16.
 */
public class QuickSortImpl implements QuickSort{
    public static void main(String[] args){
        int[] array = null;
//        array = new int[]{4,4,0,1,0,10,2,2};
//        qsort(array, 0, array.length - 1);
//        System.out.println(Arrays.toString(array));

//        array = new int[]{4,5,0,10,1,2,10,4,3};
//        qsort(array, 0, array.length - 1);
//        System.out.println(Arrays.toString(array));

        array = new int[]{1,2,30,4,5};
        qsort(array, 0, array.length - 1);
        System.out.println(Arrays.toString(array));

//        array = new int[]{9,8,7,6,5};
//        qsort(array, 0, array.length - 1);
//        System.out.println(Arrays.toString(array));
    }

    private static void insertion(int[] array, int start, int end){
        if(start >= end || start < 0 || end >= array.length){
            return;
        }
        for(int i = start + 1; i <= end; i++){
            int cur = array[i];
            for(int j = i - 1; j >= 0; j--){
                if(array[j] <= cur){
                    array[j] = array[j-1];
                }
                else{
                    array[i] = array[j+1];
                    break;
                }
            }
        }
    }

    public static void qsort(int[] array, int start, int end){
        if(start >= end || start < 0 || end >= array.length){
            return;
        }
        if(end - start < 10){
            insertion(array,start,end);
            return;
        }
        //int mid = (start+end)/2;
        int mid = (start+end)/2;
        int midVal = array[mid];
        int i = start;
        int j = end;
        //why i <=j but not i < j (Rule 4)
        while(i <= j){
            //use < but not <= (Rule 1)
            //use <= need to check index out of bound:
            //maybe all the element in array is less than or equal to the pivot
            //use < will not because at least the pivot itself is equal to midVal

            //if pivot is the greatest element in array, i will just go to the end
            //then for the right side 

            //pointer will stop at element equals or greater than the midVal
            while(array[i] < midVal){
                i++;
            }
            while(array[j] > midVal){
                j--;
            }
            //why if condition? (Rule 2)
            //
            if(i <= j){
                if(array[i] != array[j]){
                    swap(array, i, j);
                }
                //why increment after swap (Rule 3)
                //in case there two element with same value of pivot
                i++;
                j--;
            }
            else{
                //System.out.println(i + "..." + j);
            }
        }
        
        qsort(array, start, i-1);
        qsort(array, i, end);
    }

    private static void swap(int[] array, int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
}
