package com.solutions.xiaoImpl;

import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;

/**
 * Created by xiao on 14/11/16.
 */
public class LRUCache {
    class Node{
        int key;
        int val;
        Node prev;
        Node next;
        public Node(int key, int val){
            this.key = key;
            this.val = val;
        }
    }

    Node head, tail;
    Map<Integer, Node> cache;

    private final int capacity;
    private int size;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.size = 0;
        this.cache = new HashMap<>(size);
        head = new Node(0, 0);
        tail = new Node(0, 0);
        head.next = tail;
        tail.prev = head;
    }

    public int get(int key) {
        if(cache.containsKey(key)){
            Node node = cache.get(key);
            moveToFront(node);
            return node.val;
        }
        return -1;
    }

    public void set(int key, int value) {
        if(cache.containsKey(key)){
            cache.get(key).val = value;
            moveToFront(cache.get(key));
        }
        else{
            Node node = new Node(key, value);
            cache.put(key, node);
            addToFront(node);
            size++;
            if(size > capacity){
                Node last = tail.prev;
                removeLast();
                cache.remove(last.key);
                size--;
            }
        }
    }

    private void removeLast(){
        deleteNode(tail.prev);
    }

    private void moveToFront(Node node){
        deleteNode(node);
        addToFront(node);
    }

    private void addToFront(Node node){
        Node next = head.next;
        head.next = node;
        next.prev = node;
        node.prev = head;
        node.next = next;
    }

    private void deleteNode(Node node){
        Node prev = node.prev;
        Node next = node.next;
        prev.next = next;
        next.prev = prev;
    }
}