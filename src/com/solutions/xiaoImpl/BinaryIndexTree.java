package com.solutions.xiaoImpl;

import java.util.Arrays;

/**
 * Created by xiao on 15/8/16.
 */
public class BinaryIndexTree {
    public static void main(String[] args){
        test2D();
    }

    private static void test1D(){
        buildBIT1D(new int[]{3,5,-6,4,4,9});

        System.out.println(Arrays.toString(bit));

        System.out.println(getRangeSum(0,2));
        System.out.println(getRangeSum(1,4));
        update(1,10);
        System.out.println(getRangeSum(0,2));
        System.out.println(getRangeSum(1,4));
    }

    private static void test2D(){
        int[][] mat = new int[][]{
                {1,2,4,-1},
                {2,-1,6,7},
                {9,0,1,0}
        };
        buildBIT2D(mat);
        for (int[] i : mat) {
            System.out.println(Arrays.toString(i));
        }
        System.out.println("----------------");
        for (int[] i : bit2d) {
            System.out.println(Arrays.toString(i));
        }

        
        System.out.println(getSum(1,1));
    }

    private static int[] origin;
    private static int[] bit;
    //size of binary index tree array size
    private static int size;


    public static void buildBIT1D(int[] array){
        int n = array.length;
        size = n + 1;
        bit = new int[size];
        origin = new int[n];
        for(int i = 0; i < n; i++){
            update(i, array[i]);
        }
    }

    public static void update(int index, int val){
        int i = index + 1;
        int diff = val - origin[index];
        origin[index] = val;
        while(i < size){
            bit[i] += diff;
            i = i + lsb(i);
        }
    }

    //from and to: index in original array
    public static int getRangeSum(int from, int to){
        return getSum(to + 1) - getSum(from);
    }

    //i: num of elements
    public static int getSum(int i){
        int sum = 0;
        while(i > 0){
            sum += bit[i];
            i = i - lsb(i);
        }
        return sum;
    }

    //get least significant bit
    private static int lsb(int n){
        return n & (-n);
    }

    //2D BIT
    private static int[][] matrix;
    private static int[][] bit2d;
    private static int rowSize;
    private static int colSize;

    public static void buildBIT2D(int[][] mat){
        rowSize = mat.length;
        if(rowSize == 0){
            return;
        }
        colSize = mat[0].length;
        matrix = new int[rowSize][colSize];
        bit2d = new int[rowSize+1][colSize+1];
        for(int i = 0; i < rowSize; i++){
            for(int j = 0; j < colSize; j++){
                update2d(i, j, mat[i][j]);
            }
        }
    }

    public static void update2d(int row, int col, int val){
        int diff = val - matrix[row][col];
        matrix[row][col] = val;
        for(int i = row + 1;i < rowSize + 1; i += lsb(i)){
            for(int j = col + 1;j < colSize + 1; j += lsb(j)){
                bit2d[i][j] += diff;
            }
        }
    }

    public static int getSum(int row, int col){
        int sum = 0;
        int i = row + 1;
        int j = col + 1;
        for(; i > 0; i -= lsb(i)){
            for(; j > 0; j -= lsb(j)){
                sum += bit2d[i][j];
            }
        }
        return sum;
    }
}
