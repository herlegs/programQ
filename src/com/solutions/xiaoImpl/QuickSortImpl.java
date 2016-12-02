package com.solutions.xiaoImpl;

import com.questions.QuickSort;
import com.util.Util;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Created by xiao on 21/6/16.
 */
public class QuickSortImpl implements QuickSort {
    public static void main(String[] args) {
        int[] array = null;
//        array = new int[]{4,4,0,1,0,10,2,2};
//        qsort(array, 0, array.length - 1);
//        System.out.println(Arrays.toString(array));

//        array = new int[]{4,5,0,10,1,2,10,4,3};
//        qsort(array, 0, array.length - 1);
//        System.out.println(Arrays.toString(array));

        array = new int[]{3,0,3,3,0,4,3};
        //qsort(array, 0, array.length - 1);
        //System.out.println(Arrays.toString(array));

//        array = new int[]{9,8,7,6,5};
//        qsort(array, 0, array.length - 1);
//        System.out.println(Arrays.toString(array));
        qsort(array, 0, array.length - 1);
        System.out.println(Arrays.toString(array));
    }

    public static void qsort(int[] array, int start, int end) {
        if (start >= end || start < 0 || end >= array.length) {
            return;
        }
        if (end - start < 10) {
            insertionSort(array, start, end);
            return;
        }
        //int mid = (start+end)/2;
        int mid = (start + end) / 2;
        //int midVal = array[mid];
        int midVal = choosePivot(array[mid], array[start], array[end]);
        int i = start;
        int j = end;
        //why i <=j but not i < j (Rule 4)
        while (i <= j) {
            //use < but not <= (Rule 1)
            //use <= need to check index out of bound:
            //maybe all the element in array is less than or equal to the pivot
            //use < will not because at least the pivot itself is equal to midVal

            //if pivot is the greatest element in array, i will just go to the end
            //then for the right side 

            //pointer will stop at element equals or greater than the midVal
            while (array[i] < midVal) {
                i++;
            }
            while (array[j] > midVal) {
                j--;
            }
            //why if condition? (Rule 2)
            //
            if (i <= j) {
                if (array[i] != array[j]) {
                    swap(array, i, j);
                }
                //why increment after swap (Rule 3)
                //in case there two element with same value of pivot
                i++;
                j--;
            } else {
                //System.out.println(i + "..." + j);
            }
        }
        //when i j meet at pivot value point then i - j = 2
        //otherwise i - j = 1
        System.out.println("i  j: " + i + "   " + j);
        qsort(array, start, j);
        qsort(array, i, end);
    }

    private static int choosePivot(int a, int b, int c) {
        if (a > b) {
            if (c > a) {
                return a;
            } else if (c > b) {
                return c;
            } else {
                return b;
            }
        } else {
            if (c > b) {
                return b;
            } else if (c > a) {
                return c;
            } else {
                return a;
            }
        }
    }

    private static void swap(int[] array, int i, int j) {
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }

    private static void insertionSort(int[] array, int start, int end) {
        if (start >= end || start < 0 || end >= array.length) {
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
            while (j >= 0 && array[j] > cur) {
                array[j + 1] = array[j];
                j--;
            }
            array[j + 1] = cur;
        }
    }

    public static int[] mergesort(int[] array, int start, int end) {
        if (start > end) {
            return new int[0];
        }
        if (start == end) {
            return new int[]{array[start]};
        }
        int mid = start + (end - start) / 2;
        int[] left = mergesort(array, start, mid);
        int[] right = mergesort(array, mid + 1, end);
        return combine(left, right);
    }

    private static int[] combine(int[] a, int[] b) {
        int n = a.length;
        int m = b.length;
        int[] c = new int[n + m];
        int i = 0, j = 0, k = 0;
        while (i < n && j < m) {
            if (a[i] < b[j]) {
                c[k] = a[i];
                i++;
            } else {
                c[k] = b[j];
                j++;
            }
            k++;
        }
        int[] remained = (i == n) ? b : a;
        int index = (i == n) ? j : i;
        while (k < n + m) {
            c[k] = remained[index];
            index++;
            k++;
        }
        return c;
    }
}
