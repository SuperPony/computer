package creational.single;

public class Example {
    public static void main(String[] args) {
        Store s = Store.getInstance("root");
        Store s2 = Store.getInstance("jack");
        System.out.println(s2.getUser()); // root
    }
}
