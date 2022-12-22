package creational.factory;

public class Example {
    public static void main(String[] args) {
        Factory f = new MysqlFactory();
        Store s = f.generateStore();
    }
}
