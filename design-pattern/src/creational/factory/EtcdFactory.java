package creational.factory;

import creational.factory.abstract_factory.EtcdStore;

public class EtcdFactory extends Factory {
    @Override
    Store generateStore() {
        return new EtcdStore();
    }
}
