package com.util;


import java.util.List;

/**
 * Created by xiao on 6/5/16.
 */
public class Util {
    public static void p(Object o){
        System.out.print(o);
    }

    public static <T> void  pList(List<T> list){
        StringBuilder sb = new StringBuilder();
        sb.append("{");
        if(list != null){
            for (int i = 0; i < list.size(); i++) {
                sb.append(list.get(i).toString());
                if(i != list.size() - 1){
                    sb.append(",");
                }
            }
        }
        sb.append("}");
        p(sb.toString());
    }
}
