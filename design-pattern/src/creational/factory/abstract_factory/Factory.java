package creational.factory.abstract_factory;

import creational.factory.Store;

public interface Factory {
    Logger createLogger();

    Store createStore();
}