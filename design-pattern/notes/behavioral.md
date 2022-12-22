# 行为型模式

行为型模式用于负责对象之间的高效沟通以及指责委派。

# 责任链模式（Chain of Responsibility）

责任链模式用于将请求沿着链条传递，从而让每一个处理者都有机会处理该请求，并控制是否继续传递请求；

当一个责任链，每个处理者都有机会接触请求，并且它的目的不是为了找到某个处理者并处理掉请求，而是让每个处理者都做一些事情，这种责任链也被称为拦截器（Interceptor）或者过滤器（Filter）。

责任链的实现：

1. 定义处理者接口并声明处理方法的签名，以保证处理者的一致性；
2. 需要有一个"链条"来描述整条指责链，具体实现可以是在每个处理者中保存下一个处理者，或是直接使用组合模式；
3. 实现每一个处理者的处理方法，在处理方法中，需要注意两点：
    - 处理方法是否自行处理这个请求；
    - 是否沿着请求链，继续传递。

责任链的优点：

- 可以控制一连串的处理顺序；
- 遵循了单一职责，每个处理者之间相互独立；
- 遵循了开闭原则，在不需要修改现有代码的前提下，可以新增处理者。

缺点：

- 部分请求可能未被处理。

应用场景：

- 当程序需要根据请求的不同调用不同的处理方式时，可以使用责任链模式；
- 当必须按照顺序执行多个处理者时，可以使用该模式；
- 如果需要处理顺序在运行时动态变化时，可以使用该模式（通过在处理方法中，动态改变处理队列来实现）。

## Demo

```java
interface Handler {
    void setNext(Handler handle);

    void handle(Request req);
}

interface Handler2 {
    void handle(Request req);

    void addHandle(Handler2... handles);
}

// 通过组合模式，来实现责任链
public class HandleChain implements Handler2 {
    private List<Handler2> handlers = new ArrayList<>();

    public HandleChain() {
    }

    public HandleChain(Handler2... handlers) {
        this.addHandle(handlers);
    }

    public void addHandle(Handler2... handlers) {
        this.handlers.addAll(Arrays.asList(handlers));
    }

    public void handle(Request request) {
        if (this.handlers.isEmpty()) {
            return;
        }
        for (Handler2 h :
                this.handlers) {
            if (this.check(request)) {
                h.handle(request);
            } else {
                System.out.println("stop");
            }
        }
    }

    private boolean check(Request req) {
        return req.code == 200;
    }
}

// 通过设置下一个处理者的方式实现责任链
public class StartHandle implements Handler {
    private Handler next;

    @Override
    public void setNext(Handler next) {
        this.next = next;
    }

    @Override
    public void handle(Request req) {
        System.out.println("StartHandle");
        // Do something...
        if (this.next != null) {
            this.next.handle(req);
        }
    }
}

public class Example {
    public static void main(String[] args) {
        Handler startHandle = new StartHandle();
        Handler endHandle = new EndHandle();
        startHandle.setNext(endHandle);
        startHandle.handle(new Request(200, "hello world \n"));
    }
}

```

## 与其他模式的关系

- 责任链、命令、中介者、观察者4个模式用于处理请求发送者和接收者之间的不同连接方式：
    - 责任链按照顺序将请求动态的传递给一系列潜在接收者，直到请求被接收者进行处理（可以是一名也可以是多名接收者，具体看处理方法的实现逻辑）；
    - 命令模式在发送者和请求者之间建立单向连接；
    - 中介者模式清除了发送者和请求者之间的直接联系，强制它们通过一个中介对象进行间接沟通；
    - 观察者模式允许接收者动态订阅或需求接受请求。
- 责任链模式往往和组合模式配合使用；
- 责任链模式的处理者可以通过命令模式实现，从而达到对由请求代表的同一个上下文对象执行许多不同的操作；
- 责任链和装饰模式的结构非常相似，两者也都依赖递归机制将需要执行的操作传递给一系列工作对象，但两者有根本的区别：
    - 责任链的处理者都是独立的,而装饰器依赖于被装饰的对象；
    - 责任链可以随时中止请求的传递，装饰器做不到。

# 命令模式（Command）

命令模式用于将请求封装成一个对象（命令对象），然后执行该命令，同时可以达到请求排队，记录请求日志甚至是撤销、恢复的操作。其核心思想是将客户端的请求与请求实现解耦，客户端只需要了解各个命令的含义并进行调用而无需了解具体实现。

具体实现：

1. 声明一个仅有执行命令的接口；
2. 根据需要创建各种命令类，命令类中保存对接收者类（执行具体任务的类）的引用，并在执行方法中调用接收者对应的处理接口；
3. 如果需要记录历史执行或撤销、恢复操作，可以创建发送者类，由发送者来统一执行命令的调度；
    - 发送者类中通过栈类型存储历史命令；
    - 命令接口中声明撤销操作，发送者内将历史记录中推出最新的命令执行其撤销操作；
    - 发送者内实现恢复接口（将历史记录中命令遍历执行撤销操作）。
4. 客户端创建并调用命令对象来执行具体的业务，或由发送者进行命令的调度执行。

命令模式的优点：

- 遵循了单一职责原则，可以将请求的触发与具体实现的类分开；
- 遵循了开闭原则，可以在不修改已有客户端代码的情况下，在程序中创建新的命令；
- 可以实现撤销（undo）和恢复（redo）的功能；
- 可以实现具体操作的延迟执行；
- 可以将一组简单命令组合成一个复杂的命令。

命令模式的缺点：

- 导致代码变得更复杂，因为在发送者和接收者之间增加的一个全新的层次。

应用场景：

- 需要将操作作为参数时可以使用命令模式，因为命令模式将操作转为了独立的命令对象，因此操作就可以作为参数进行传递或保存在其他对象中；
- 需要将操作放入队列时，可以使用命令模式；
- 想要实现操作的回滚功能，可以使用命令模式。

## 与其他模式的关系

- 命令模式与策略模式看起来很想，因为两者都可以通过一些行为来参数化对象，但其各自的意图却不同；
    - 命令模式可以将任何操作转为对象，操作的参数将成为命令对象的成员属性，可以借此来实现接收者对象的延迟执行，操作放入队列或历史命令等功能；
    - 策略模式通常用于实现完成某件事的不同方式，让客户端可以在同一个上下文类中切换算法。
- 命令模式可以结合备忘录模式使用，来实现撤销操作，命令模式用于执行执行对象，备忘录用来保存一条命令执行前的内部状态；

# 迭代器（Iterator）

迭代器模式让聚合对象在不暴露内部的前提下，按顺序遍历内部各个元素；

具体实现：

1. 创建 Iterator 接口，并声明访问下一个元素以及判断是否遍历完毕的方法；
2. 创建一个专门创建 Iterator 的接口；
3. 聚合对象实现专门创建 Iterator 的接口；
4. 创建具体的 Iterator 对象：
    - Iterator 对象内部存储当前索引的具体位置；
    - 每次调用获取下一个元素的方法时，返回索引对应的聚合对象内的元素，同时索引进行移动；
    - 通过是否遍历完毕的方法判断聚合对象内部的元素已经遍历完毕。

迭代器模式的优点：

- 遵循了单一职责原则，通过将繁杂的遍历算法抽取为独立的迭代器类；
- 遵循了开闭原则，可以按需新增新的迭代器，而无需修改现有的迭代器；
- 由于迭代器内部保存了当前的索引位置，所以可以做到暂停迭代或是回滚等操作。

迭代器模式的缺点：

- 增加程序的复杂度。

## Demo

```java
interface Iterator<T> {
    int getIndex();

    T getNext();

    boolean hasNext();
}

interface ConcreteCollection<T> {
    Iterator<T> iterator();
}

class MyCollection<T> implements ConcreteCollection<T> {
    private T[] array;

    public MyCollection(T... array) {
        this.array = Arrays.copyOfRange(array, 0, array.length);

    }

    @Override
    public Iterator<T> iterator() {
        return new MyIterator();
    }

    class MyIterator implements Iterator<T> {
        int index;

        @Override
        public int getIndex() {
            return this.index;
        }

        @Override
        public T getNext() {
            T res = array[index];
            this.index++;
            return res;
        }

        @Override
        public boolean hasNext() {
            return this.index < MyCollection.this.array.length;
        }
    }
}

public class Example {
    public static void main(String[] args) {
        MyCollection<Integer> myCollection = new MyCollection<>(1, 3, 5);
        for (Iterator<Integer> it = myCollection.iterator(); it.hasNext(); ) {
            Integer i = it.getNext();
            System.out.println(i);
        }
    }
}
```

# 中介模式（Mediator）

中介模式专门用于解决各个模块之间的耦合关系，通过中介对象来作为各个子模块之间的沟通桥梁，从原本子模块之间相互沟通的多对多关系，变成了子模块与中介对象交互的一对一沟通关系;外观模式常用语有着众多交互的
UI 组件上。

举例说明，假定有一个表单，里面有着多选框以及全选、取消全选、反选
4个组件，这4个组件之间相互作用影响，因此导致了高度的耦合以及复杂性；此时就可以通过中介对象来负责与4个组件进行沟通，而不是组件之间相互沟通。

中介模式的优势：

- 遵循了单一职责原则，多个组件之间的交流抽取到同一位置，使其更容易维护和理解；
- 遵循了开闭原则，无需修改组件就能添加新的中介者；
- 可以更加方便的复用各个组件同时减轻组件之间的耦合。

中介模式的缺点：

- 中介者很容易变成所有组件的上帝对象。

# 备忘录模式（Memento）

备忘录模式用于在不暴露对象的细节（状态）的前提下，为对象进行快照备份；其本质上是通过创建一个专门的状态对象来保存和返回对象的状态，并通过历史对象来存储所有的状态对象，当对象需要进行回滚操作时，从历史对象中拿到需要的状态对象并进行状态回滚即可；
在备忘录模式中主要有三个角色：

- 原发器（Originator）：即需要进行状态存储的对象，原发器对象起码要有保存当前状态以及重置当前状态的接口；
- 备忘录（Memento）：即保存原发器状态的对象，备忘录对象通常使用构造函数来创建并保存状态，同时要有返回状态的方法以便后续原发器进行回滚；
- 责任人（CareTaker）：即保存备忘录的历史队列对象，通常该对象只需有保证有添加、返回备忘录的接口即可。

备忘录模式的优点：

- 可以在不破坏对象封装的前提下，进行对象的快照创建；
- 可以通过责任人来维护原发器的历史状态，从而简化原发器的代码。

备忘录模式的缺点：

- 如果客户端频繁地创建备忘录，会导致大量内存被消耗；
- 负责人必须完整跟踪原发器的生命状态，才能销毁弃用的备忘录；
- 大部分动态语言无法确保备忘录中的状态不被修改。

## Demo

```java
public class Example {
    public static void main(String[] args) {
        Car car = new Car("Benz", "red");
        System.out.println(car); // Car{band='Benz', color='red'}
        car.setBand("Audi");
        System.out.println(car); // Car{band='Audi', color='red'}
        car.undo();
        System.out.println(car); // Car{band='Benz', color='red'}
    }
}

interface Originator {
    void saveState();

    void restore(Memento memento);
}

interface Memento {
    Memento getState();
}

interface MementoHistory {
    Memento pop();

    void push(Memento memento);

    boolean isEmpty();
}

class Car implements Originator {
    private String band;
    private String color;

    private MementoHistory mementoHistory = new History();

    public Car(String band, String color) {
        this.band = band;
        this.color = color;
    }

    public void setBand(String band) {
        this.saveState();
        this.band = band;
    }

    public void setColor(String color) {
        this.saveState();
        this.color = color;
    }

    @Override
    public void saveState() {
        Memento memento = new CarMemento(this.band, this.color);
        this.mementoHistory.push(memento);

    }

    public void undo() {
        Memento carMemento = this.mementoHistory.pop().getState();
        this.restore(carMemento);
    }

    @Override
    public void restore(Memento memento) {
        CarMemento carMemento = (CarMemento) memento;
        this.band = carMemento.band;
        this.color = carMemento.color;
    }

    @Override
    public String toString() {
        return "Car{" +
                "band='" + band + '\'' +
                ", color='" + color + '\'' +
                '}';
    }

    class CarMemento implements Memento {
        private final String band;
        private final String color;

        public CarMemento(String band, String color) {
            this.band = band;
            this.color = color;
        }

        @Override
        public Memento getState() {
            return this;
        }
    }
}

class History implements MementoHistory {
    private Stack<Memento> history = new Stack<>();

    public Memento get(int index) {
        return this.history.get(index);
    }

    @Override
    public Memento pop() {
        return this.history.pop();
    }

    @Override
    public void push(Memento memento) {
        if (this.history.size() >= 10) {
            this.pop();
        }
        this.history.push(memento);
    }

    @Override
    public boolean isEmpty() {
        return this.history.isEmpty();
    }
}
```

# 观察者模式（Observer）

观察者模式又称为发布订阅模式（Publish-Subscribe：Pub/Sub），该模式是一种通知机制，让发送通知对的一方（被观察方 Observable/发布方
Publisher）和接收通知的一方（观察者 Observer/订阅方 Subscriber ）能够彼此分离，互不影响。

实现步骤

- 声明被观察方接口，确保里面有添加、删除、通知订阅者的方法，实际的被观察者；
- 声明观察者接口，里面具有执行方法，以便后续被观察方进行统一的调度通知；
    - 通常而言，被观察者在通知订阅方时，往往需要传递一些参数，可根据实际开发时情况声明在执行方法中，或将这些参数抽离为一个事件对象进行传递。

观察者模式的优点：

- 遵循了开闭原则，无需修改被观察者，即可拓展新的观察者类；
- 可以在运行时，建立双方的关系。

## Demo

```java
// Publisher
interface Observable {
    void addSubscriber(Observer subscriber);

    void removeSubscriber(Observer subscriber);
}

// Subscriber
interface Observer {
    void execute(Event event);
}

class Event {
    private final String eventType;
    private final String data;

    public Event(String eventType, String data) {
        this.eventType = eventType;
        this.data = data;
    }

    public String getEventType() {
        return eventType;
    }

    public String getData() {
        return data;
    }
}

class EmailPublisher implements Observable {
    private List<Observer> subscribers = new ArrayList<>();

    private Event event;

    public EmailPublisher(Event event) {
        this.event = event;
    }

    @Override
    public void addSubscriber(Observer observer) {
        this.subscribers.add(observer);
    }

    @Override
    public void removeSubscriber(Observer observer) {
        this.subscribers.remove(observer);
    }

    public void notifySubscribers() {
        for (Observer s :
                this.subscribers) {
            s.execute(this.event);
        }
    }
}

class User implements Observer {
    @Override
    public void execute(Event event) {
        System.out.println(event.getEventType());
    }
}
```

# 状态模式（State）

状态模式允许一个对象在其内部状态改变时，改变它的行为；对象看起来就好像是修改了它的类。其实就是，如果对象因为内部某个状态的不同，而产生不同的行为时，普通的写法会出现大量的 `if else`
来判断和处理各种不同的状态值；状态模式则是将各种 `if else` 的逻辑分别拆分到不同的状态类中进行处理。

状态模式的优势

- 遵循了单一职责，将与特定状态相关的逻辑放到单独的类中；
- 遵循了开闭原则，无需修改已有状态类；
- 消除了臃肿的状态机条件语句 `if else`，简化上下文代码。

## Demo

```java
interface Replier {
    void reply(String name);
}

class LeisureReplier implements Replier {
    @Override
    public void reply(String name) {
    }
}

class BusyReplier implements Replier {
    @Override
    public void reply(String name) {
    }
}

enum ReplyStates {
    LEISURE, BUSY
}

class Service {
    private ReplyStates state;

    public void reply(String userName) {
        Replier replier;
        switch (this.state) {
            case BUSY:
                replier = new BusyReplier();
                break;
            case LEISURE:
            default:
                replier = new LeisureReplier();
                break;
        }
        replier.reply(userName);
    }

    public ReplyStates getState() {
        return state;
    }

    public void setState(ReplyStates replyState) {
        this.state = replyState;
    }
}
```

# 策略模式（Strategy）

策略模式指定义一组算法，然后在运行时，根据需要客户端可以自行设置要使用的算法；最典型的例子，在下单时，根据活动的不同，计算各自的活动价；其核心思想是，当功能逻辑相同，而核心算法不同时，将这部分核心算法进行抽离，在运行时动态决定使用那种算法。

策略模式的的实现：

- 定义策略的接口，并声明算法的接口，以保证算法的一致性；
- 创建具体的策略对象，并实现各自的算法；
- 定义使用策略的上下文对象，其内部引用策略接口，并创建策略对象的设置方法；

策略模式的优点：

- 可以在运行时决定使用的算法；
- 可以将算法的实现和使用相互独立；
- 遵循了开闭原则，无需修改上下文对象，即可托增新的策略。

策略模式的缺点：

- 客户端必须了解每个策略的不同；

## Demo

```java
interface DiscountStrategy {
    BigDecimal computedDiscount(BigDecimal total);
}

class UserDiscountStrategy implements DiscountStrategy {
    /**
     * 计算满减后的价格
     * @param total 价格
     * @return 计算后的价格
     */
    @Override
    public BigDecimal computedDiscount(BigDecimal total) {
        BigDecimal magnification = total.divide(BigDecimal.valueOf(100)).setScale(0, RoundingMode.DOWN);
        return total.subtract(magnification.multiply(BigDecimal.valueOf(30)));
    }
}

class VipUserDiscountStrategy implements DiscountStrategy {
    /**
     * VIP 用户在满减基础上，再7折.
     * @param total 价格.
     * @return 计算后的 VIP 用户价格.
     */
    @Override
    public BigDecimal computedDiscount(BigDecimal total) {
        return total.multiply(BigDecimal.valueOf(0.7));
    }
}

// 上下文对象
class Order {
    private BigDecimal price;
    private DiscountStrategy strategy;

    public Order(BigDecimal price) {
        this.price = price;
    }

    public void setStrategy(DiscountStrategy strategy) {
        this.strategy = strategy;
    }

    public BigDecimal computedTotal() {
        if (this.strategy == null) {
            return this.price;
        }
        this.price = this.strategy.computedDiscount(this.price);

        return this.price;
    }
}

public class Example {
    public static void main(String[] args) {
        Order order = new Order(BigDecimal.valueOf(150));
        System.out.println(order.computedTotal()); // 150
        order.setStrategy(new UserDiscountStrategy());
        System.out.println(order.computedTotal()); // 120
        order.setStrategy(new VipUserDiscountStrategy());
        System.out.println(order.computedTotal()); // 84
    }
}
```

## 与其他模式的区别

- 装饰模式用于增强对象的某些功能，策略模式用于改变某些功能的本质；
- 策略模式可以看成策略模式的拓展，两者都是基于组合机制，通过将部分工作委派给"帮手"对象，从而改变不同情境下的行为；区别在于，策略模式下，"帮手"之间完全独立，相互之间不知道其他对象的存在。


# 模版方法模式（Template Method）

模版方法模式的主要思想是，定义一个操作的一系列步骤，对于某些暂时确定不下来的步骤，作为抽象方法，留给子类去实现，这样不同的子类就可以定义出不同的步骤。
需要注意的是，为了防止子类重写父类的骨架方法，可以将骨架方法骨定义为 `final` ；同时对于需要子类实现的方法，应定义为外部不可见。

模版方法模式的优点：

- 可仅允许修改一个操作中的特定步骤；
- 可以讲重复代码提升到一个基类中。

模版方法模式的缺点：

- 模版方法中的步骤越多，其维护工作就越麻烦；
- 通过子类抑制默认步骤实现可能导致违反里氏替换原则。

# 访问者模式（Visitor）

访问者模式是一种用于操作一组对象的操作，其目的在于在不改变对象定义的前提下，通过定义和新增不同的访问者，来处理一组对象中不同类型的对象。

实现方法：

1. 在访问者接口中，声明一组方法，分别对应要处理的组内对象的不同类型；
2. 声明并实现各个访问者对象，每个访问者对象仅需实现自己对应的对象类型处理方法即可；
3. 实际操作对象中声明接收方法，用于接收访问者，并调用实际要用的方法，并将接收者传入。

访问者模式的优点：

- 遵循了开闭原则，可以引入在不同类对象上执行的新行为， 且无需对这些类做出修改；
- 遵循了单一职责，可以将同一行为的不同版本移到同一个类中。

访问模式的缺点：

- 每当复合对象中新增了一个类型的对象时，都需要在访问者接口中添加新的处理方法，同时导致所有具体访问者都需要修改代码；
- 在访问者与其对应的对象进行交互时，访问者可能没有访问对象内部属性或方法的权限。

## Demo

```java

// 访问者接口
interface Visitor {
    void visitDir(File dir);

    void visitFile(File file);
}

// 实际的访问者，该访问者专门处理 class 类型文件
class ClassFileVisitor implements Visitor {
    @Override
    public void visitDir(File dir) {
        // ...
    }

    @Override
    public void visitFile(File file) {
        if (file.getName().endsWith(".class")) {
            System.out.println("Will clean class file: " + file);
        }
    }
}

// 实际的访问者，该访问者专门处理 java 类型文件
class JavaFileVisitor implements Visitor {
    @Override
    public void visitDir(File dir) {
        // ....
    }

    @Override
    public void visitFile(File file) {
        if (file.getName().endsWith(".java")) {
            System.out.println("Found java file: " + file);
        }
    }
}

// 实际的访问者，该访问者专门处理目录
class DirVisitor implements Visitor{
    @Override
    public void visitDir(File dir) {
        System.out.println("Found dir: " + dir);
    }

    @Override
    public void visitFile(File file) {
        // ...
    }
}

class FileStructure {
    private File path;

    public FileStructure(File path) {
        this.path = path;
    }

    public void handle(Visitor visitor) {
        this.scan(this.path, visitor);
    }

    private void scan(File file, Visitor visitor) {
        if (file.isDirectory()) {
            visitor.visitDir(file);
            for (File sub :
                    Objects.requireNonNull(file.listFiles())) {
                this.scan(sub, visitor);
            }
        } else if (file.isFile()) {
            visitor.visitFile(file);
        }
    }
}

public class Example {
    public static void main(String[] args) {
        FileStructure fileStructure = new FileStructure(new File("."));
        fileStructure.handle(new JavaFileVisitor());
    }
}
```





