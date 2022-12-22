package creational.factory.abstract_factory;

import creational.factory.Store;

public class DevelopmentFactory implements Factory{
    @Override
    public Store createStore() {
        return new EtcdStore();
    }

    @Override
    public Logger createLogger() {
        return new FileLogger();
    }
}
