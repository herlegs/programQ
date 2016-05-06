package com.solutions.xiaoImpl;

import com.questions.Permutation;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

import static com.util.Util.*;

/**
 * Created by xiao on 6/5/16.
 */
public class PermutationImpl implements Permutation{
    public static void main(String args[]){
        char[] input;
        Map<String, Integer> result = new HashMap<>();

        input = new char[]{'y', '么', '哒'};
        result.clear();
        permutation(input, 0, input.length - 1, result);
        p("---unique permutation---\n");
        pResult(result);

    }

    public static void permutation(char[] array, int start, int end, Map<String, Integer> result){
        if(start == end){
            pArray(array);
            result.put(Arrays.toString(array), 0);
            return;
        }
        for(int i = start; i <= end; i++){
            swap(array, start, i);
            permutation(array, start + 1, end, result);
            swap(array, start, i);
        }
    }

    private static void swap(char[] array, int i, int j){
        char tmp = array[i];
        array[i] = array[j];
        array[j] = tmp;
    }

    public static void pArray(char[] array){
        p(Arrays.toString(array) + "\n");
    }

    public static void pResult(Map<String, Integer> result){
        StringBuilder sb = new StringBuilder();
        for(Map.Entry entry: result.entrySet()){
            sb.append(entry.getKey() + "\n");
        }
        p(sb.toString());
    }
}
