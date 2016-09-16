package com.solutions.xiaoImpl;

import sun.text.normalizer.Trie;

import java.util.Arrays;

/**
 * Created by xiao on 17/8/16.
 */
public class SegmentTree {
    public static void main(String[] args){
        int[] array = new int[]{1,-2,3,4};
        buildSeg(array);
        System.out.println(Arrays.toString(seg));
        updateRange(1,2,0);
        System.out.println(Arrays.toString(seg));
        System.out.println(Arrays.toString(lazy));
        System.out.println(query(1,2));
    }

    private static int[] seg;
    private static int[] lazy;
    private static int arraySize;
    private static int segSize;

    public static void buildSeg(int[] array){
        arraySize = array.length;
        segSize = getSegSize(arraySize);
        seg = new int[segSize];
        lazy = new int[segSize];
        for(int i = 0; i < arraySize; i++){
            updatePoint(i, array[i]);
        }
    }

    private static int getSegSize(int arraySize){
        int size = 1;
        while(size < arraySize){
            size *= 2;
        }
        return size * 2 - 1;
    }

    public static void updatePoint(int i, int val){
        updatePointRecur(i, val, 0, arraySize - 1, 0);
    }

    private static void updatePointRecur(int i, int val, int start, int end, int root){
        if(start == end){
            seg[root] = val;
            return;
        }
        int mid = start + (end - start)/2;
        int leftIndex = root*2 + 1;
        int rightIndex = root*2 + 2;
        if(i <= mid){
            updatePointRecur(i, val, start, mid, leftIndex);
        }
        else{
            updatePointRecur(i, val, mid + 1, end, rightIndex);
        }
        seg[root] = Math.max(seg[leftIndex], seg[rightIndex]);
    }

    public static void updateRange(int i, int j, int val){
        updateRangeRecur(i, j, val, 0, arraySize - 1, 0);
    }

    public static void updateRangeRecur(int i, int j, int val, int low, int high, int root){
        if(lazy[root] != 0){
            seg[root] = lazy[root];
            if(low != high){
                lazy[root*2+1] = lazy[root];
                lazy[root*2+2] = lazy[root];
            }
            lazy[root] = 0;
        }
        //totally within
        if(low >= i && high <= j){
            if(val < seg[root]){
                return;
            }
            else{
                seg[root] = val;
                if(low != high){
                    lazy[root*2+1] = val;
                    lazy[root*2+2] = val;
                }
            }
            return;
        }
        //outside
        if(high < i || low > j){
            return;
        }
        //partially
        int mid = low + (high - low)/2;
        updateRangeRecur(i, j, val, low, mid, root*2+1);
        updateRangeRecur(i, j, val, mid + 1, high, root*2+2);
        seg[root] = Math.max(seg[root*2+1], seg[root*2+2]);
    }

    public static int query(int i, int j){
        return query(i, j, 0, arraySize - 1, 0);
    }

    private static int query(int i, int j, int low, int high, int root){
        if(lazy[root] != 0){
            seg[root] = lazy[root];
            if(high != low){
                lazy[root*2+1] = lazy[root];
                lazy[root*2+2] = lazy[root];
            }
            lazy[root] = 0;
        }
        //totally within
        if(low >= i && high <= j){
            return seg[root];
        }
        //outside
        if(high < i || low > j){
            return Integer.MIN_VALUE;
        }
        //partially
        int mid = low + (high - low)/2;
        int left = query(i, j, low, mid, root*2+1);
        int right = query(i, j, mid + 1, high, root*2+2);
        return Math.max(left, right);
    }
}
