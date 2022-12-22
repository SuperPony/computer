package creational.single;

import java.util.Objects;

public class Store {

    private static Store instance;
    private String user;

    private Store(String user) {
        this.user = user;
    }

    public static Store getInstance(String user) {
        if (instance == null) {
            instance = new Store(user);
        }

        return instance;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Store store = (Store) o;
        return Objects.equals(user, store.user);
    }

    @Override
    public int hashCode() {
        return Objects.hash(user);
    }


    public String getUser() {
        return user;
    }

    public void setUser(String user) {
        this.user = user;
    }

}
