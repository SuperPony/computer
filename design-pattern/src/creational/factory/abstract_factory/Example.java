package creational.factory.abstract_factory;

import creational.factory.Store;

public class Example {
    public static final String env = "development";

    public static void main(String[] args) {
        Factory f = null;
        Store s = null;
        Logger l = null;
        if (env.equalsIgnoreCase("development")) {
            f = new DevelopmentFactory();
        }
        if (env.equalsIgnoreCase("production")) {
            f = new ProductFactory();
        }
        s = f.createStore();
        l = f.createLogger();
    }
}
