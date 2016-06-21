package com.solutions.xiaoImpl;

import com.questions.CutTree;
import com.util.Tree;
import com.util.TreeUtil;
import com.util.Util;

import java.util.*;

import static com.util.TreeUtil.printTree;
import static com.util.Util.*;

/**
 * Created by xiao on 31/5/16.
 */
public class CutTreeImpl implements CutTree {
    public static void main(String args[]){
        Tree test = TreeUtil.deserialize("1,#,3,4,5,6,7,#,8");
        List<Tree> remain = getRemainingTree(test);
        for(Tree tree: remain){
            printTree(tree);
            System.out.println("---------------");
        }
        
    }

    public static List<Tree> getRemainingTree(Tree root){
        Set<Tree> result = new HashSet<>();
        Queue<Tree> queue = new LinkedList<>();
        queue.offer(root);
        result.add(root);
        while(!queue.isEmpty()){
            Tree node = queue.poll();
            if(node != null){
                if(canCut(node) && result.contains(node)){
                    result.remove(node);
                    if(node.left != null){
                        result.add(node.left);
                        queue.offer(node.left);
                    }
                    if(node.right != null){
                        result.add(node.right);
                        queue.offer(node.right);
                    }
                }
            }
        }
        return new ArrayList<>(result);
    }

    /***
     * whether the node will be cut
     * @param node
     * @return
     */
    public static boolean canCut(Tree node){
        if(node == null){
            return false;
        }
        if(node.val % 2 == 1){
            return true;
        }
        return false;
    }

}
