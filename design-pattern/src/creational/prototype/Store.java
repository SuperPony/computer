package creational.prototype;

import java.util.Objects;

public abstract class Store {
    public String user;
    public String pwd;

    public Store(String user, String pwd) {
        this.user = user;
        this.pwd = pwd;
    }

    public Store(Store store) {
        if (store != null) {
            this.user = store.user;
            this.pwd = store.pwd;
        }
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Store store = (Store) o;
        return Objects.equals(user, store.user) && Objects.equals(pwd, store.pwd);
    }

    @Override
    public int hashCode() {
        return Objects.hash(user, pwd);
    }

    @Override
    public abstract Store clone();
}
