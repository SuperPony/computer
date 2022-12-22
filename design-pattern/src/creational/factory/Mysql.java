package creational.factory;

public class Mysql  implements Store {
    @Override
    public boolean connect() {
        return false;
    }

    @Override
    public boolean close() {
        return false;
    }
}
