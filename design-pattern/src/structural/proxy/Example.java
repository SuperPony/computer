package structural.proxy;

public class Example {
    public static void main(String[] args) {
        Mysql mysql = new MysqlProxy("root", "123456");
        mysql.connect();
    }
}
