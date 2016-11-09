package com.solutions.xiaoImpl;

import com.util.Util;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * Created by xiao on 16/9/16.
 */
public class SubSets {
    //print all possible sub set for a given list
    public static void main(String[] args){
            testSubSetK();
     }

    private static void testSubSet(){
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        int[] array = new int[]{1,2,3};
        Arrays.sort(array);
        subSet(array, 0, array.length - 1, new ArrayList<Integer>(), result);
        printSet(result);
    }

    private static void testSubSetK(){
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        int[] array = new int[]{1,2,3};
        Arrays.sort(array);
        subSetK(array, 0, array.length - 1, 2, new ArrayList<Integer>(), result);
        printSet(result);
    }

    private static void subSet(int array[], int start, int end, List<Integer> curSet, List<List<Integer>> result){
        result.add(new ArrayList<Integer>(curSet));
        for(int i = start; i <= end; i++){
            if(i != start && array[i] == array[i-1]){
                continue;
            }
            curSet.add(array[i]);
            subSet(array, i+1, end, curSet, result);
            curSet.remove(curSet.size()-1);
        }
    }

    private static void subSetK(int array[], int start, int end, int k, List<Integer> curSet, List<List<Integer>> result){
        if(curSet.size() == k){
            result.add(new ArrayList<Integer>(curSet));
        }
        if(curSet.size() > k){
            return;
        }
        for(int i = start; i <= end; i++){
            curSet.add(array[i]);
            subSetK(array, i+1, end, k, curSet, result);
            curSet.remove(curSet.size()-1);
        }
    }

    private static void printSet(List<List<Integer>> sets){
        sets.forEach(list -> {
            Util.pList(list);
            System.out.println("");
        });
    }
}
