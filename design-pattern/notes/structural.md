# 结构型模式

结构型模式用于将对象或类组装成较大的结构，同时保持结构的灵活和高效。

所有的结构型模式都是基于组合原则，即一个对象因该将部分工作委派给另一个对象。

# 适配器模式

适配器模式用于使不兼容的对象之间相互合作、兼容，例如，当 A 对象期望调用 a 接口，而 B 对象的仅有 b 接口时，即可通过装饰器包装
B 对象，由装饰器的 a 方法调用 B 接口的 b 方法完成适配；又或是 A 对象期望 json 数据，而 B 对象仅提供 xml 数据时，也可通过适配器模式解决。
本质上是通过包装类将不兼容的对象进行包装处理，从而达到适配兼容的目的。

适配器模式优点：

* 遵循了单一职责原则，将处理兼容的代码，从对象内部转移到适配器中；
* 遵循了开闭原则，当两个对象之间不兼容时，不需要修改内部的代码来进行适配；

## Demo

```java

// 客户端期望的接口、抽象类、类....
interface Expect {
    void req();
}

// Target.method 符合客户端期望的 req 方法，但需要适配
class Target {
    public void method() {
        // ....
    }
}

// 适配器
class TargetAdapter implements Expect {
    private Target target;

    public TargetAdapter(Target target) {
        this.target = target;
    }

    @Override
    public void req() {
        this.target.method();
    }
}


public class Client {
    public static void main(String[] args) {
        Target target = new Target();
        Expect ex = new TargetAdapter(target);
        ex.request();
    }
}
```

## 与其他模式的关系

- 桥接模式常用开发前期进行设计，使其能够将程序的各部分独立分开以便开发；而适配器通常用于已有的程序，或作用于一些不可控的对象上（例如第三方库），使其相互之间能够兼容合作；
- 适配器可以对已有对象的接口进行修改，而装饰模式能够在不改变对象接口的前提下强化对象的功能，此外，装饰模式还支持递归组合，适配器无法实现；
- 适配器为被封装的对象提供不同的接口（例如更改返回数据类型，或延伸出客户端期望的接口），代理模式则为对象提供相同的接口，装饰模式则能为对象提供加强的接口；
- 外观模式的目的在于内聚复杂的子模块，从而简化客户端的使用步骤，因此外观模式不介意去提供一些新的接口，同时可能会作用于多个对象（例如多个子系统的对象），而适配器模式则往往运用已有的接口，通常也只封装一个对象；

# 桥接模式

桥接模式用于将一个大类或一些列紧密相关的类拆分成抽象和实现两个部分，从而达到彼此之间可以独立变化的目的；说白了，桥接模式的核心思想是，在建模时，将模型的特性作为抽象类型存储在模型内部，在实例化时，将各个特性的具体实现注入进去；而并非通过派生各种子类来
实现；举例说明，一个汽车模型有品牌、颜色两个属性，假定现在有品牌 Benz，Audi， 颜色 Red，Blue，如果用继承的方式实现，就会出现 BenzRedCar，BenzBlueCar，AudiRedCar，AudiBlueCar 4个子类，当后期再增加其他的属性，或新增一个品牌、颜色都会导致子类的爆炸增长；
而通过桥接模式则是在汽车对象内声明品牌、颜色两个抽象属性，再创建与之对应的属性对象，当实例化汽车对象时，仅需把需要的品牌、颜色对象注入进去即可;

## Demo

```java
class Band {
  private final String name;

  public Band(String name) {
    this.name = name;
  }

  public String getName() {
    return name;
  }
}
class Color {
  private final String rgb;

  public Color(String rgb) {
    this.rgb = rgb;
  }

  public String getRgb() {
    return rgb;
  }
}
class Car {
    private Band band;
    private Color color;

    public Car(Band band, Color color) {
        this.band = band;
        this.color = color;
    }

    public String getBand() {
        return band.getName();
    }

    public String  getColor() {
        return color.getRgb();
    }
}
public class Example {
  public static void main(String[] args) {
    Car car = new Car(
            new Band("Benz"),
            new Color("#fff")
    );
    System.out.println(car.getBand()); // Benz
  }
}
```

# 组合模式

组合模式将对象组合成树状结构，并可以像独立对象一样使用，其本质是将对象视为一个容器以及叶节点，容器可以用来存储其他叶节点，而叶节点则作为实际工作的基本单元；为了保证一致性，叶节点需要使用接口来制定统一的标准。

组合模式优点：

* 可以利用多态或递归机制来更方便的使用复杂树结构；
* 遵循了开闭原则，无需更改代码，就可以在应用中添加新元素，使其成为对象树的一部分。

应用场景：

可以将业务的核心模型作为树状结构表示的，组合模式才有意义；很典型的例子为虚拟 DOM 树或目录树。

实现方法：

1. 确定应用的核心模型能够以树状结构表示，将其分解为简单元素（叶子结点）和容器，容器必须能够同时包含简单元素和其他容器；
2. 声明组件接口及其一系列方法，用于确保元素的一致型；
3. 组件接口因该具有添加和删除子元素的方法；
4. 将需要的容器和叶子节点继承组件接口并实现（容器节点内需要具有一个复合型数据类型的属性来存储其他子元素或容器）。

## Demo

```java

interface Dom {
    String getName();

    List<Dom> getChildren();
    // More method...
}

public class ElementDom implements Dom {
    public List<Dom> children = new ArrayList<>();
    public final String name;

    public ElementDom(String name, Dom... domes) {
        this.name = name;
        this.addAll(domes);
    }

    public ElementDom(String name) {
        this.name = name;
    }

    public boolean add(Dom dom) {
        return this.children.add(dom);
    }

    public boolean addAll(Dom... domes) {
        return this.children.addAll(Arrays.asList(domes));
    }

    public boolean remove(Dom dom) {
        return this.children.remove(dom);
    }

    @Override
    public String getName() {
        return name;
    }

    public List<Dom> getChildren() {
        return children;
    }
}
```

# 装饰模式（Decorator）

装饰模式用于为指定对象增添一些额外功能，就增加功能而言，相比于创造子类，更加灵活；其本质上是把一个个附加功能，用 Decorator
的方式给一层一层地累加到原始数据源上，最终，通过组合获得所需的功能。

装饰模式的优势：

* 无需派生子类，即可拓展对象的行为；
* 可以在运行时，按需动态的配置额外行为；
* 单一职责原则，可以将实现了许多不同行为的一个大类拆分为多个较小的类。

实现方法：

1. 找出基本组件和可选层次（装饰器）的通用方法，创建一个组件接口并在其中声明这些方法；
2. 创建基本组件及其行为；
3. 创建装饰基类，使用一个成员变量存储指向被封装对象的引用。该成员变量必须被声明为组件接口类型，从而能在运行时连接具体组件和装饰。装饰基类必须将所有工作委派给被封装的对象；
4. 确保所有类实现组件接口；
5. 将装饰基类派生出各种具体装饰，具体装饰必须在调用父类方法 （总是委派给被封装对象） 之前或之后执行自身的行为；
6. 客户端按需组合各种装饰。

## Demo

```java
interface Notifier {
    boolean send(String message);
}

abstract class BaseDecorator implements Notifier {
    protected Notifier notifier;

    public BaseDecorator(Notifier notifier) {
        this.notifier = notifier;
    }
}

class Email implements Notifier {
    @Override
    public boolean send(String message) {
        return false;
    }
}

class LogDecorator extends BaseDecorator {
    public LogDecorator(Notifier notifier) {
        super(notifier);
    }

    @Override
    public boolean send(String message) {
        boolean res = notifier.send(message);
        // write log...
        return res;
    }
}

class ValidationDecorator extends BaseDecorator {
    public ValidationDecorator(Notifier notifier) {
        super(notifier);
    }

    @Override
    public boolean send(String message) {
        // validation params...

        return notifier.send(message);
    }
}

public class Example {
    public static void main(String[] args) {
        Notifier n = new LogDecorator(new ValidationDecorator(new Email()));
        n.send("hello world");
    }
}
```

## 与其他模式的关系

- 装饰模式用于增加核心对象的功能，策略模式则是为了改变其核心对象的功能；
-
装饰模式和代理模式的区别主要在于意图，装饰器的生成由客户端按需进行组合控制，因此，客户端明确知道自身使用的增强过后的装饰器，而代理模式则不同，代理模式通常自行管理其服务对象（核心对象）的生命周期，意图在于通过代理替换其核心对象，让客户端无感知。

# 外观模式

外观模式将复杂的子系统各个接口，集中整合为一个统一的高层接口，让使用者只需要和高层接口打交道，而无需了解各个子系统该怎么用；
拿下单业务举例，一个正常的下单业务想要完成，需要使用者至少了解商品、用户、订单三个模块的使用，此时就可以通过外观模式，将三个模块的交互提炼到更高层次的对象中，由高层次的模块进行商品、用户、订单的处理，而使用者仅需要关注高层次模块即可。

外观模式的优点：

* 可以让高层对象独立于复杂的子系统。

外观模式的缺点：

* 可能会导致高层次对象与各个子系统的耦合。

## 与其他模式的区别

- 外观模式与中介模式的职责类似，都是用于协调、组织多个子模块的作业，但其有着本质上的区别：
  - 外观模式的核心思想是减轻客户端与对各子模块的依赖，通过内聚各个子模块的操作并提炼出一个新的接口，由该接口调度各个子模块之间的作业；因此子模块之间还是直接沟通，并存在耦合，子模块也不知道外观的存在；且本身并没有延伸出新的功能；
  - 中介模式其核心思想是解除各子系统间的耦合，通过一个中介对象作为沟通的中心，各个子模块间不再直接沟通，由中介对象负责模块之间的沟通传递，因此从多模块之间的多向沟通变成了模块与中介对象的双向沟通。

# 享元模式

享元模式的核心思想是，如果一个对象中，有大量不可变的状态或对象本身一经创建不可变；那么就可以考虑抽离这一部分不可变的状态，作为一个享元对象，同时也没有反复创建这个对象必要，直接调用方法返回一个共享的实例即可。

在 Java 中，有大量享元模式的使用案例，例如包装类型 `Byte`, `Integer` 这些不变类。

享元模式的优点：

- 当遇到有大量重复且不可变的状态或对象时，可以使用享元模式来节省内存。

享元模式的缺点：

- 代码会变得更加复杂。

# 代理模式

代理模式用于控制对某个对象的访问，从而使其在访问时，能够在原有的行为上，进行需要的拓展等；其意图是通过代理对象接管被代理对象，从而是客户端无感知。

代理模式的优点：

- 可以在客户端无感知的前提下，接管被代理的服务对象；
- 即使服务对象还未准备好，或不存在，代理也可以正常工作；
- 遵循了开闭原则，可以在不修改被代理对象的前提下，对功能进行拓展或修改。

实现方式：

1. 考虑是否可以将被代理类的方法统一成接口，以此来保证代理对象和被代理对象之间的可交换性，如果不可行，可以考虑将代理对象作为被代理对象的子类；
2. 创建实际的代理类，内部存储引用的被代理对象，通常而言，由代理类负责被代理类的整个生命周期管理（构建、销毁...）,特殊情况下，也可以由外部注入；

## Demo

```java

interface Store {
    boolean connect();
    // Other method...
    boolean close();
}

class Mysql implements Store {
    protected String user;
    protected String pwd;

    public Mysql(String user, String pwd) {
        this.user = user;
        this.pwd = pwd;
    }

    public boolean connect(){
        System.out.println("Mysql connect");
        return  true;
    }
    // Other method...
    public boolean close(){
        System.out.println("Mysql close");
        return  false;
    }
}

class MysqlProxy extends Mysql {

    public MysqlProxy(String user, String pwd) {
        super(user, pwd);
    }

    @Override
    public boolean connect() {
        // Do something....
        System.out.println("MysqlProxy connect");
        super.connect();

        return false;
    }

    @Override
    public boolean close() {
        // Do something....
        System.out.println("MysqlProxy close");
        super.close();

        return false;
    }
}

class MysqlProxy2 implements Store {
    private final Mysql mysql;

    public MysqlProxy2(String user, String pwd) {
        this.mysql = new Mysql(user, pwd);
    }

    @Override
    public boolean connect() {
        // Do something....
        System.out.println("MysqlProxy connect");
        this.mysql.connect();

        return false;
    }

    @Override
    public boolean close() {
        // Do something....
        System.out.println("MysqlProxy close");
        this.mysql.close();

        return false;
    }
}

public class Example {
    public static void main(String[] args) {
        Mysql mysql = new MysqlProxy("root", "123456");
        mysql.connect();
    }
}
```



