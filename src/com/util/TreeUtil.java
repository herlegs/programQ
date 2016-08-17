package com.util;

import java.util.LinkedList;
import java.util.Queue;
import static com.util.Util.*;

/**
 * Created by Xiao on 30/5/2016.
 */
public class TreeUtil {
    public static final String NULL_NODE = "#";
    public static final String SEPARATOR = ",";

    public static String serialize(Tree root){
        if(root == null){
            return "";
        }
        Queue<Tree> queue = new LinkedList<Tree>();
        StringBuilder sb = new StringBuilder();
        queue.offer(root);
        while(!queue.isEmpty()){
            Tree node = queue.poll();
            if(node == null){
                sb.append(NULL_NODE + SEPARATOR);
            }
            else{
                sb.append(String.valueOf(node.val) + SEPARATOR);
                queue.offer(node.left);
                queue.offer(node.right);
            }
        }
        String result = sb.toString();
        int i = result.length() - 1;
        for(; i >= 0; i--){
            String c = String.valueOf(result.charAt(i));
            if(!c.equals(SEPARATOR) && !c.equals(NULL_NODE)){
                break;
            }
        }
        return result.substring(0, i + 1);
    }

    public static Tree deserialize(String str){
        Queue<Tree> queue = new LinkedList<>();
        String[] elements = str.split(SEPARATOR);
        if(elements.length == 0){
            return null;
        }
        Tree root = new Tree(Integer.parseInt(elements[0]));
        queue.offer(root);
        for(int i = 1; i < elements.length;){
            Tree parent = queue.poll();
            if(!elements[i].equals(NULL_NODE)){
                //left child
                parent.left = new Tree(Integer.parseInt(elements[i]));
                queue.offer(parent.left);
            }
            i++;
            if(i < elements.length && !elements[i].equals(NULL_NODE)){
                //right child
                parent.right = new Tree(Integer.parseInt(elements[i]));
                queue.offer(parent.right);
            }
            i++;
        }
        return root;
    }

    public static void printTree(Tree root){
        if(root == null){
            return;
        }
        Queue<Tree> queue = new LinkedList<>();
        Queue<Integer> depth = new LinkedList<>();
        queue.offer(root);
        depth.offer(getTreeRootLeftDepth(root));
        while(!queue.isEmpty()){
            int count = queue.size();
            int prevDepth = 0;
            for(int i = 0; i < count; i++){
                Tree node = queue.poll();
                int leftDepth = depth.poll();
                printSpace(leftDepth - prevDepth);
                if(leftDepth - prevDepth == 0){
                    System.out.print(" ");
                }
                if(node != null){
                    System.out.print(node.val);
                }
                else{
                    System.out.print("#");
                }
                prevDepth = leftDepth;
                if(node != null){
                    queue.offer(node.left);
                    queue.offer(node.right);
                    depth.offer(leftDepth - 1);
                    depth.offer(leftDepth + 1);
                }
            }
            System.out.println("\n");
        }
    }
    
    private static void printSpace(int depth){
        for(int i = 0; i < depth; i++){
            System.out.print("\t");
        }
    }

    private static int getTreeRootLeftDepth(Tree root){
        if(root == null) return 0;
        return Math.max(getTreeRootLeftDepth(root.left) + 1, getTreeRootLeftDepth(root.right) - 1);
    }

//    public static int getTreeHeight(Tree node){
//        if(node == null){
//            return 0;
//        }
//        return Math.max(getTreeHeight(node.left), getTreeHeight(node.right)) + 1;
//    }
//
//    public static String getWhiteSpace(int count){
//        StringBuilder sb = new StringBuilder();
//        for(int i = 0; i < count; i++){
//            sb.append(" ");
//        }
//        return sb.toString();
//    }
}
