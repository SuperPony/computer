package creational.prototype;

import java.util.Objects;

public class Mysql extends Store{
    public int port;

    public Mysql(Mysql mysql) {
        super(mysql);
        this.port = mysql.port;
    }

    public Mysql(String user, String pwd, int port) {
        super(user, pwd);
        this.port = port;
    }

    public Mysql(Store store, int port) {
        super(store);
        this.port = port;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        if (!super.equals(o)) return false;
        Mysql mysql = (Mysql) o;
        return port == mysql.port;
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), port);
    }

    @Override
    public Store clone() {
        return new Mysql(this);
    }
}
