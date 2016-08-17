package com.solutions.xiaoImpl;

/**
 * Created by xiao on 16/8/16.
 */
public class KMP {
    public static void main(String[] args){

    }

    private int[] lps(String s){
        int n = s.length();
        int[] lps = new int[n];
        lps[0] = 0;
        int len = 0;
        int i = 1;
        while(i < n){
            if(s.charAt(i) == s.charAt(len)){
                len++;
                lps[i] = len;
                i++;
            }
            else{
                if(len != 0){
                    len = lps[len - 1];
                }
                else{
                    lps[i] = 0;
                    i++;
                }
            }
        }
        return lps;
    }
}

