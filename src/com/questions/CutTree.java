package com.questions;

import com.util.Tree;

import java.util.List;

/**
 * Created by xiao on 31/5/16.
 */
public interface CutTree {
    /***
     * given function to check each of node in a Tree can be removed
     * output a list of Tree which remains after cut;
     */

    //List<Tree> getRemainingTree(Tree root);
    /***
     * whether the node will be cut
     * @param node
     * @return
    public static boolean canCut(Tree node){
        if(node == null){
            return false;
        }
        if(node.val % 2 == 1){
            return true;
        }
        return false;
    }
    */
}
