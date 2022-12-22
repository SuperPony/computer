package creational.factory;

public class Etcd implements Store {

    @Override
    public boolean connect() {
        return false;
    }

    @Override
    public boolean close() {
        return false;
    }

}
