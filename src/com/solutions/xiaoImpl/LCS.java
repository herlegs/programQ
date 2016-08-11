package com.solutions.xiaoImpl;

/**
 * Created by xiao on 2/8/16.
 */
public class LCS {
    public static void main(String[] args){
        String test = lcs("abcdexcdef", "abcdef");
        System.out.println(test);
    }

    public static String lcs(String a, String b){
        int n = a.length();
        int m = b.length();
        //max common substring if the substring end at
        int[][] lcs = new int[n+1][m+1];
        int max = 0;
        int index = 0;
        for(int i = 1; i <= n; i++){
            for(int j = 1; j <=m; j++){
                if(a.charAt(i-1) == b.charAt(j-1)){
                    if(i == 1 || j == 1){
                        lcs[i][j] = 1;
                    }
                    else{
                        lcs[i][j] = lcs[i-1][j-1] + 1;
                    }
                    if(lcs[i][j] > max){
                        max = lcs[i][j];
                        index = i;
                    }
                }
                //else length will be 0
            }
        }
        return a.substring(index - max, index);
    }
}
