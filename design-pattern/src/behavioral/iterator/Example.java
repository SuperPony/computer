package behavioral.iterator;

public class Example {
    public static void main(String[] args) {
        MyCollection<Integer> myCollection = new MyCollection<>(1, 3, 5);
        for (Iterator<Integer> it = myCollection.iterator(); it.hasNext();) {
            Integer i = it.getNext();
            System.out.println(i);
        }
    }
}
