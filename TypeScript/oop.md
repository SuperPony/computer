# Class

- TS 是单继承；
- 如果不写权限修饰符，则默认 public；
- 子类无法重写继承而来的属性的权限修饰符；
- `readonly`关键字可以用于修饰属性，用于表示只读，仅允许默认赋值以及在构造函数中赋值。

```typescript

enum Sex {
    _,
    Man,
    Woman,
}

// TS 中抽象类与普通类区别不大，主要区别在于：
// 1. 抽象类不允许被实例化
// 2. 抽象类中可以添加抽象方法，继承抽象类的子类必须实现抽象方法
abstract class Animal {
    age: number
    sex: Sex

    // 抽象类构造函数建议将修饰符设置为受保护的
    protected constructor(age: number = 0, sex: Sex = Sex.Man) {
        this.sex = sex
        this.age = age
    }

    say() {
        console.log(this.age)
    }

    abstract eating(): string
}

// TS 只允许单继承
// TS 不能重写属性的修饰符
class People extends Animal {

    // protected ,private ,public 三种访问修饰符，默认 public，作用和传统 oop 语言一致。
    // 通常情况下，默认值会声明在构造函数的参数上
    protected name: string = "tom"

    // 不能重写属性的修饰符
    // protected age:number

    // // 只读属性, 只允许设置默认值或在构造方法中赋值
    // readonly age: number
    // // 静态属性
    // static sex: number = 1
    // // 只读静态属性, 只允许声明默认值
    // static readonly sex2: number = 2

    constructor(name: string, age: number = 0, sex: Sex = Sex.Man) {
        // super 表示父类，在构造函数中调用表示调用父类构造函数
        // super 通常只会在构造函数中使用，其他地方使用 this
        super(age, sex);
        this.name = name
    }

    dance() {
        console.log(`${this.name} age is ${this.age}, he is dancing.`)
    }

    eating(): string {
        console.log(`${this.name} is eating`)
        return this.name
    }

    // public say() {
    //     super.say()
    //     console.log(this.name)
    // }
}
```


# InterFace

- 接口用于让类继承，从而确保类必须实现接口；
- 接口也可以达到自定义 Object 类型的作用；
- 接口与自定义类型的区别在于，接口可多次声明，表示合并。

```typescript
interface Store {
    host: string,
    port: number
}

// 接口对于变量而言与自定义类型类似，区别在于接口可以重复声明，表示合并
// 接口也可以用类的继承，与传统 OOP 语言意义一样。
interface Store {
    addData(): void
}

let s: Store = {
    host: "127.0.0.1",
    port: 3306,
    addData() {
    }
}

class Mysql implements Store {
    host: string;
    port: number;

    constructor(host: string, port: number) {
        this.host = host
        this.port = port
    }

    addData(): void {
        // do something...
    }
}

const addData = (s: Store) => {
    s.addData()
}

let m = new Mysql("127.0.0.1", 3306)

addData(m)
```
