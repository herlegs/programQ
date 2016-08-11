package com.solutions.xiaoImpl;

import java.util.Arrays;

/**
 * Created by xiao on 4/7/16.
 */
public class MaxSubArraySum {
    public static void main(String[] args){
        int result = max1D(new int[]{1,2,-4,4,6,-2,-1,-1,5});
        //System.out.println(result);

        result = max2D(new int[][]{
                {1,-1,3},
                {1,-4,3}
        });
        System.out.println(result);
        long a = 2323;
        int s = 12;
        System.out.println(a>s);
    }

    public static int max2D(int[][] matrix){
        int row = matrix.length;
        if(row == 0){
            return 0;
        }
        int col = matrix[0].length;
        int max = matrix[0][0];
        for(int i = 0; i < col; i++){
            int[] array = new int[row];
            for(int j = i; j < col; j++){
                for(int k = 0; k < row; k++){
                    array[k] += matrix[k][j];
                }
                int colMax = max1D(array);
                if(colMax > max){
                    max = colMax;
                }
            }
        }
        return max;
    }


    public static int max1D(int[] array){
        int length = array.length;
        if(length == 0){
            return 0;
        }
        int maxEndHere = array[0];
        int[] maxEndRecord = new int[]{0,1};
        int result = array[0];
        int[] resultRecord  = new int[]{0,1};
        
        for(int i = 1; i < length; i++){
            if(maxEndHere < 0){
                maxEndHere = array[i];
                maxEndRecord[0] = i;
                maxEndRecord[1] = i + 1;
            }
            else{
                maxEndHere = maxEndHere + array[i];
                maxEndRecord[1] = i + 1;
            }
            if(maxEndHere > result){
                result = maxEndHere;
                resultRecord[0] = maxEndRecord[0];
                resultRecord[1] = maxEndRecord[1];
            }
        }
        //System.out.println(Arrays.toString(resultRecord));
        return result;
    }

}
