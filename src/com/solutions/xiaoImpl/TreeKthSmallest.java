package com.solutions.xiaoImpl;


import com.util.Tree;
import com.util.TreeUtil;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by xiao on 9/11/16.
 *
 * get kth smallest element in a BST
 */
public class TreeKthSmallest {
    public static void main(String[] args) {
        //BST
        Tree tree = TreeUtil.deserialize("5,3,6,2,4,#,9,1,#,#,#,7,10");
        TreeUtil.printTree(tree);
    }

    public List<Integer> getKth(Tree root, int k){
        List<Integer> result = new ArrayList<>();
        return result;
    }
}
