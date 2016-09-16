package com.solutions.xiaoImpl;

/**
 * Created by xiao on 9/9/16.
 */
public class MinSubArraySum {
    //given an array, an target sum, give an sub array which has sum bigger than target
    //and has minimum length
    public static void main(String[] args){
        int[] array = new int[]{};
        int target = 22;
        int min = minSub(array, target);
        System.out.println(min);
    }

    private static int minSub(int[] array, int target){
        int min = Integer.MAX_VALUE;
        int i = 0, j = 0;
        int n = array.length;
        int sum = 0;
        while(j < n){
            while(j < n){
                sum += array[j];
                if(sum >= target){
                    break;
                }
                else{
                    j++;
                }
            }
            //System.out.println("after expand i j :" + i +" "+ j + " sum:"+sum);
            while(i<=j){
                if(sum - array[i] >= target){
                    sum -= array[i];
                    i++;
                }
                else{
                    break;
                }
            }
            int length = j - i + 1;
            System.out.println("i j :" + i +" "+ j + " sum:"+sum);
            if(sum >= target){
                min = Math.min(min, length);
                j++;
            }
        }
        return min == Integer.MAX_VALUE ? 0 : min;
    }
}
