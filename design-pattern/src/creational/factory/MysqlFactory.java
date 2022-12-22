package creational.factory;

public class MysqlFactory  extends Factory{
    @Override
    Store generateStore() {
        return new Mysql();
    }
}
