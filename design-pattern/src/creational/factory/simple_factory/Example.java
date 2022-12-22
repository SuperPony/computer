package creational.factory.simple_factory;

import creational.factory.Store;

public class Example {
    public static void main(String[] args) {
        Factory f = new Factory();
        Store s = f.generateStore("mysql");
    }
}
