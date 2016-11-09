package com.solutions.xiaoImpl;

import com.questions.Permutation;

import java.util.*;

import static com.util.Util.*;

/**
 * Created by xiao on 6/5/16.
 */
public class PermutationImpl implements Permutation{
    public static void main(String args[]){
        char[] input;
        Map<String, Integer> result = new HashMap<>();

//        input = new char[]{'1', '2', '1'};
//        result.clear();
//        Arrays.sort(input);
//        permutation(input, 0, input.length - 1, result);
//        p("---unique permutation---\n");
//        pResult(result);

        List<Integer> list = new ArrayList<Integer>();
        list.add(2);list.add(2);
        list.add(4);list.add(5);
        Collections.sort(list);
        permutation(list, new boolean[list.size()], new ArrayList<>());
    }

    public static void permutation(char[] array, int start, int end, Map<String, Integer> result){
        if(start == end){
            pArray(array);
            result.put(Arrays.toString(array), 0);
            return;
        }
        for(int i = start; i <= end; i++){
            if(i != start && array[i] == array[i-1]){
                continue;
            }
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

    private static void permutation(List<Integer> list, boolean[] visited, List<Integer> cur){
        if(cur.size() == list.size()){
            pList(cur);
            System.out.print("\n");
            return;
        }
        for(int i = 0; i < list.size(); i++){
            //if this element is already added in current result OR
            //if this element is same as previous element and previous element is not picked yet
            //so we can ensure same element is always follow the sorted order (and treated as a whole)
            if(visited[i] || (i != 0 && list.get(i) == list.get(i-1) && !visited[i-1])){
                continue;
            }
            visited[i] = true;
            cur.add(list.get(i));
            permutation(list, visited, cur);
            cur.remove(cur.size()-1);
            visited[i] = false;
        }
    }

//    private static void swap(List<Integer> list, int i, int j){
//        int a = list.get(i);
//        int b = list.get(j);
//        list.set(i, b);
//        list.set(j, a);
//    }

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
