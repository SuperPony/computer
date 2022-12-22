# 创建型模式

创建型模式关注点是如何创建对象，其核心思想是隐藏创建对象的细节，并将创建和使用对象相分离，使两者能相对独立的变换。

# 工厂模式（Factory Pattern）

工厂模式将对象的创建和实际使用进行分离，非常适合用于创建类型不同但又具有相同特性的对象，且对象之间彼此相互独立；工厂模式本质上：
1. 通过接口约束产品之间的一致性；
2. 定义工厂对象，专门用于构造不同对象(简单工厂模式)；
    1. 如果只有一个工厂对象，会导致工厂对象和各个产品对象以及未来加入的产品对象之间的强耦合（每次加入新的对象，需要修改工厂的代码）；
    2. 将工厂定义为抽象类，并将创建对象的方法作为抽象方法，在抽象类工厂的基础上，派生出多个专门的工厂类构造专门的对象即可。

工厂模式的优点：

- 避免创建和使用的紧密耦合；
- 遵循了单一指责原则，可以将对象的创建放在单一位置，从而使代码更加容易维护；
- 遵循了开闭原则，无需修改原有工厂的代码，通过继承就可以拓展新的工厂。

## 与其他模式的区别

- 抽象工厂通常基于一组工厂方法，是一组工厂聚合的产物；
- 原型模式不基于继承，而工厂基于继承，但它不需要初始化步骤；
- 工厂方法是模版模式的一种特殊形式，且可以作为一个大型模版模式中的一个步骤。

## Demo

```java

import com.sun.istack.internal.NotNull;// 产品

// 对象的相同特性
interface Car {
    void run();

    void stop();
}

// 不同类型的对象
class Benz implements Car {
    @Override
    public void run() {
    }

    @Override
    public void stop() {
    }
}

class Audi implements Car {
    @Override
    public void run() {
    }

    @Override
    public void stop() {
    }
}

// 简单工厂
class SimpleFactory {
    public Car createCar(@NotNull String type) {
        if (type == null) {
            return null;
        }
        if (type.equalsIgnoreCase("benz")) {
            return new Benz();
        }
        if (type.equalsIgnoreCase("audi")) {
            return new Audi();
        }
        // Other product
        // ...

        return null;
    }
}

abstract class Factory {
    abstract Car createCar();
}

class BenzFactory extends Factory {
    @Override
    Car createCar() {
        return new Benz();
    }
}

class AudiFactory extends Factory {
    @Override
    Car createCar() {
        return new Audi();
    }
}

class Demo {
    void demo() {
        Factory f = new BenzFactory();
        Car c = f.createCar();
    }
}
```

# 抽象工厂

抽象工厂在工厂模式的基础上，将工厂进行抽象以及行为的描述（定义为接口），从而派生出多个相互独立的子工厂对象，很经典的场景为，根据环境的不同，使用不同组的配置对象。

优点：

- 可以确保同一工厂生成的产品相互匹配；
- 具有工厂模式所有的优点。

## Demo

```java
public class Example {
    public final String env = "dev";
    public void example() {
        Factory f = null;
        Logger l = null;
        Store s = null;
        if (env.equalsIgnoreCase("dev")) {
             f = new DevelopmentFactory();
        } else if (env.equalsIgnoreCase("product")) {
             f = new ProductFactory();
        }
        l = f.createLogger();
        s = f.createStore();
    }
}
// 将工厂进行抽象
interface Factory {
    Logger createLogger();
    Store createStore();
}

class DevelopmentFactory implements Factory {
    @Override
    public Store createStore() {
        return new DevelopmentStore();
    }

    @Override
    public Logger createLogger() {
        return new DevelopmentLog();
    }
}

class ProductFactory implements Factory{
    @Override
    public Logger createLogger() {
        return new ProductLog();
    }
    @Override
    public Store createStore() {
        return new ProductStore();
    }
}

interface Logger {
    void write();
    void close();
}

interface Store {
    boolean connect();
    boolean close();
}

class DevelopmentLog implements Logger {
    @Override
    public void write() {
    }
    @Override
    public void close() {
    }
}

class ProductLog implements Logger {
    @Override
    public void write() {
    }
    @Override
    public void close() {
    }
}

class DevelopmentStore implements Store {
    @Override
    public boolean connect() {
        return false;
    }
    @Override
    public boolean close() {
        return false;
    }
}

class ProductStore implements Store {
    @Override
    public boolean connect() {
        return false;
    }
    @Override
    public boolean close() {
        return false;
    }
}
```

## 与其他模式的区别

- 生成器模式的侧重点在于如何分步生成复杂的对象，而抽象工厂的侧重点在于生产一系列相关对象，抽象工厂会立即返回产品，生成器则允许在获取产品前进行一些额外的构造选项。

# 生成器模式

生成器模式用于生成、构造参数配置特别多的对象，并且拥有分步构造的能力；生成器模式的本质是通过创建一个专门的构造对象, 使用构造对象来对指定的对象进行构建。

```java

class Car {
    private final int seats;
    private final String color;
    private final String brand;
    // Other params....
    
    public Car(String brand, int seats, String color) {
        this.brand = brand;
        this.seats = seats;
        this.color = color;
    }
}

interface Builder {
    void setSeats(int seats);
    void setColor(String color);
    void setBrand(String brand);
    // Other setter...
}

class CarBuilder implements Builder {
    private int seats;
    private String color;
    private String brand;

    @Override
    public void setSeats(int seats) {
        this.seats = seats;
    }

    @Override
    public void setColor(String color) {
        this.color = color;
    }
    
    @Override 
    public void setBrand(String brand) {
        this.brand = brand;
    }
    
    public Car generateCar(){
        return  new Car(brand, seats, color);
    }
}

public class Example {
   public static void main(String[] args) {
      CarBuilder b = new CarBuilder();
      Director d = new Director();
      d.constructCityCar(b);
      Car cityCar = b.generateResult();
      d.constructSuv(b);
      Car suv = b.generateResult();
      System.out.println(cityCar.getColor()); // red
      System.out.println(suv.getColor()); // green
   }
}
```

在上述的基础上，还可以引入"管理员"的角色对象，通过"管理员"进行构造对象的调用，从而达到控制构建步骤的效果。

```java
class Director {
    public void constructBenz(Builder builder) {
        builder.setSeats(4);
        builder.setColor("red");
        builder.setBrand("Benz");
    }

    public void constructAudi() {
        builder.setBrand("Audi");
        builder.setColor("red");
        builder.setSeats(4);
    }
}
```

# 原型模式

原型模式用于克隆对象，其本质是把克隆从外部克隆转为对象内部实现克隆，因此所有可克隆的对象，必须继承并实现一个通用的克隆接口；相对于外部克隆而言，内部克隆更加内聚，而且可以复制内部属性。

```java
public abstract class Store {
    public String user;
    public String pwd;

    public Store(String user, String pwd) {
        this.user = user;
        this.pwd = pwd;
    }

    public Store(Store store) {
        if (store != null) {
            this.user = store.user;
            this.pwd = store.pwd;
        }
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Store store = (Store) o;
        return Objects.equals(user, store.user) && Objects.equals(pwd, store.pwd);
    }

    @Override
    public int hashCode() {
        return Objects.hash(user, pwd);
    }

    @Override
    public abstract Store clone();
}

public class Mysql extends Store{
   public int port;

   public Mysql(Mysql mysql) {
      super(mysql);
      this.port = mysql.port;
   }

   public Mysql(String user, String pwd, int port) {
      super(user, pwd);
      this.port = port;
   }

   public Mysql(Store store, int port) {
      super(store);
      this.port = port;
   }

   @Override
   public boolean equals(Object o) {
      if (this == o) return true;
      if (o == null || getClass() != o.getClass()) return false;
      if (!super.equals(o)) return false;
      Mysql mysql = (Mysql) o;
      return port == mysql.port;
   }

   @Override
   public int hashCode() {
      return Objects.hash(super.hashCode(), port);
   }

   @Override
   public Store clone() {
      return new Mysql(this);
   }
}

public class Example {

   public static void main(String[] args) {
      Mysql s1 = new Mysql("root", "123456", 3306);
      Mysql s2 = (Mysql) s1.clone();
      System.out.println(s2.pwd); // 123456
      System.out.println(s1.equals(s2)); // true
      s1.pwd = "654321";
      System.out.println(s1.equals(s2)); // false
   }
}

```

# 单例模式

单例模式用于确保对象只会被实例化一次，常用于构建全局唯一的对象资源。单例模式的本质是：

- 构造函数私有化，不允许外部访问；
- 通过一个静态方法作为构造函数，由该方法调用私有的构造函数创建对象实例，并将实例保存至私有静态变量中，后续对于该静态方法的调用，都将返回这一缓存对象。

使用单例模式有几个点需要注意：
- 多线程环境下，需要特殊处理，防止多个线程多次创建单例对象；
- 单例模式做单元测试可能比较困难，因为许多测试框架基于继承的方式模拟创建测试对象，由于单例模式的构造函数是私有的，因此绝大部分语言无法重写静态方法。

## Demo

```java

public class Store {

    private static Store instance;
    private String user;

    private Store(String user) {
        this.user = user;
    }

    public static Store getInstance(String user) {
        if (instance == null) {
            instance = new Store(user);
        }

        return instance;
    }
}

```

## 与其他模式的区别

- 如果对象的内所有的状态能简化为一个享元对象，则享元模式和单例模式就类似，但其有着本质上的区别：
  - 单例模式只会有1个实例，但享元可以有多个实例，且实例间内在状态可以不同；
  - 单例对象是可以变的，享元对象不可变。