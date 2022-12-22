package creational.prototype;

public class Example {

    public static void main(String[] args) {
        Mysql s1 = new Mysql("root", "123456", 3306);
        Mysql s2 = (Mysql) s1.clone();
        System.out.println(s2.pwd); // 123456
        System.out.println(s1.equals(s2)); // true
        s1.pwd = "654321";
        System.out.println(s1.equals(s2)); // false
    }
}