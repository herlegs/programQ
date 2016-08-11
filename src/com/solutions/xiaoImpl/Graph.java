package com.solutions.xiaoImpl;

import java.util.*;

/**
 * Created by xiao on 3/8/16.
 */
public class Graph {
    public static void main(String[] args){
        int[][] test = new int[][]{
                {1,3},{1,2},
                {3,4},{2,5},
                {4,1}
                //{3,7},{7,8},{7,4},{8,4}
        };
        Map<Integer, List<Integer>> graph = tog(test);
        dfs(graph, 1, new HashSet<Integer>());
    }
    
    public static Map<Integer, List<Integer>> tog(int[][] m){
        Map<Integer, List<Integer>> graph = new HashMap<>();
        for(int[] link : m){
            int f = link[0];
            int t = link[1];
            if(!graph.containsKey(f)){
                graph.put(f, new ArrayList<Integer>());
            }
            graph.get(f).add(t);
        }
        return graph;
    }
    
    public static void dfs(Map<Integer, List<Integer>> graph, int root, Set<Integer> processing){
        List<Integer> adjs = graph.get(root);
        Stack<Integer> stack = new Stack<>();
        processing.add(root);
        stack.push(root);
        if(adjs != null){
            for(int adj: adjs){
                if(processing.contains(adj)) continue;
                dfs(graph, adj, processing);
            }
        }
        processing.remove(root);
        System.out.print(root + " , ");
        //iterative
//        while(!stack.isEmpty()){
//            int node = stack.pop();
//            processing.add(node);
//            System.out.print(node);
//            List<Integer> adjs = graph.get(node);
//            if(adjs != null)
//                adjs.forEach(adj->{
//                    if(!processing.contains(adj))
//                        stack.push(adj);
//                });
//        }
    }

}
