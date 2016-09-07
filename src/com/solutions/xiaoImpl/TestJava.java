package com.solutions.xiaoImpl;

import java.util.*;

import static com.util.Util.p;

/**
 * Created by Xiao on 30/5/2016.
 */
public class TestJava {
    public static void main(String[] args){
        Map<Integer, List<Integer>> map = new HashMap<>();
        List<Integer> graph[] = new ArrayList[3];
        p(graph[0]);
        map.put(1, null);
        List<Integer> list = new ArrayList<>();

        for(Map.Entry<Integer, List<Integer>> entry : map.entrySet()){
            map.remove(1);
            p(map.remove(1, null));
        }
        Set<Integer> set = new HashSet<>();

    }

    private static int[] aaa(){
        return new int[]{};
    }

}
