package creational.factory.simple_factory;

import creational.factory.Etcd;
import creational.factory.Mysql;
import creational.factory.Store;

public class Factory {
    public Store generateStore(String type) {
        if (type == null) {
            return null;
        }
        if (type.equalsIgnoreCase("mysql")) {
            return new Mysql();
        }
        if (type.equalsIgnoreCase("etcd")) {
            return new Etcd();
        }

        return null;
    }

}



