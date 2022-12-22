package structural.proxy;

public class MysqlProxy extends Mysql {

    public MysqlProxy(String user, String pwd) {
        super(user, pwd);
    }

    @Override
    public boolean connect() {
        // Do something....
        System.out.println("MysqlProxy connect");
        super.connect();

        return false;
    }

    @Override
    public boolean close() {
        // Do something....
        System.out.println("MysqlProxy close");
        super.close();

        return false;
    }
}


class MysqlProxy2 implements Store {
    private final Mysql mysql;

    public MysqlProxy2(String user, String pwd) {
        this.mysql = new Mysql(user, pwd);
    }

    @Override
    public boolean connect() {
        // Do something....
        System.out.println("MysqlProxy connect");
        this.mysql.connect();

        return false;
    }

    @Override
    public boolean close() {
        // Do something....
        System.out.println("MysqlProxy close");
        this.mysql.close();

        return false;
    }
}