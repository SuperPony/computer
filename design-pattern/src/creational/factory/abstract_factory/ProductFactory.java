package creational.factory.abstract_factory;

import creational.factory.Store;

public class ProductFactory implements Factory{
    @Override
    public Store createStore() {
        return new MysqlStore();
    }

    @Override
    public Logger createLogger() {
        return new MongoLogger();
    }
}
