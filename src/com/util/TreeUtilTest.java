package com.util;

/**
 * Created by xiao on 15/8/16.
 */
public class TreeUtilTest {
    public static void main(String[] args){
        Tree tree = TreeUtil.deserialize("1,2,#,4,5,7,8,9,10");
        TreeUtil.printTree(tree);
    }
}
