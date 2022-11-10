
import java.util.*;
public class Graph {
    private Map<Integer,List<Integer>> map = new  HashMap<>();
    
    private void addVertex(Integer vertex) {
        map.put(vertex, new ArrayList<>());
    }

    public void insert(Integer vertex, Integer edge, boolean bidirectional){
        if(!map.containsKey(vertex)){
            addVertex(vertex);

        }
        if(!map.containsKey(edge)){
            addVertex(edge);
        }
        map.get(vertex).add(edge);
        if(bidirectional){
            map.get(edge).add(vertex);
        }
    }

    public void display() {
    
        for(Integer x : map.keySet()){
            System.out.print(x+": ");
            for(Integer y: map.get(x)){
                System.out.print(y+" ");
            }
            System.out.println();
        }
    }
    public static void main(String[] args) {
        Graph graph = new Graph();
        graph.insert(3, 5, true);
        graph.insert(3, 4, true);
        graph.insert(5, 6, false);
        graph.display();
    }
    
}

import java.util.*;
public class Graph {
    private Map<Integer,List<Integer>> map = new  HashMap<>();
    
    private void addVertex(Integer vertex) {
        map.put(vertex, new ArrayList<>());
    }

    public void insert(Integer vertex, Integer edge, boolean bidirectional){
        if(!map.containsKey(vertex)){
            addVertex(vertex);

        }
        if(!map.containsKey(edge)){
            addVertex(edge);
        }
        map.get(vertex).add(edge);
        if(bidirectional){
            map.get(edge).add(vertex);
        }
    }

    public void display() {
    
        for(Integer x : map.keySet()){
            System.out.print(x+": ");
            for(Integer y: map.get(x)){
                System.out.print(y+" ");
            }
            System.out.println();
        }
    }
    public static void main(String[] args) {
        Graph graph = new Graph();
        graph.insert(3, 5, true);
        graph.insert(3, 4, true);
        graph.insert(5, 6, false);
        graph.display();
    }
    
}