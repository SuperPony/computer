package creational.factory.abstract_factory;

import creational.factory.Store;

public class MysqlStore implements Store {
    @Override
    public boolean connect() {
        return false;
    }

    @Override
    public boolean close() {
        return false;
    }
}
