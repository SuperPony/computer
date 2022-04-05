# types

常用数据类型，主要分为两大块：

- JS 现有类型：number/string/boolean, object...;
- TS 新增类型：联合类型，自定义类型（类型别名），接口，元组，字面量，枚举，void，any...

# Index

- 原始类型
- 对象类型
    - array
    - 函数
        - 可选参数
    - object
        - 可选属性
- 接口类型
- 联合类型
- 类型别名
- 元组
- 枚举
- void
- never
- any, unknown
- 类型推论
- 类型断言
- 字面量类型

# 原始类型

JS 中的原始类型，声明方式与 js 中类型名称一致； 格式为 `let param: type = value;`。

```typescript
let age: number = 18;
let user: string = "tom";
let isLoading: boolean = true;
let un: undefined = undefined;
let n: null = null;
let s: symbol = Symbol()
```

# 对象类型

TS 中的对象类型更加细化，每个具体对象都有自己的类型语法。

## array

声明方式有两种：

- 推荐 `let param: type[] = [val1...];`;
- `let param: Array<type>= [val1...];`;

注意：

- 数组使用联合类型时，要加小括号以表示数组的值可以是指定类型；

```typescript
// array
let arr1: number[] = [1, 2, 3]; // 推荐
let arr2: Array<string> = ["tom", "jack"];
let arr3: (number | string)[] = [26, "jack"] // 联合类型，带有小括号表示值可以是其中任意类型
```

## 函数

函数的类型实际上指参数、返回值的类型；

- `function name(param: type, ...): returnType {...};`;
- `const func = (param: type, ...): returnType => {...};`;
- 同时指定参数和返回类型： `let func: (param: type, ...l) => returnType;`。
    - `let func3: (name: string) => string = (name) => name;`。

### 可选参数

由于 TS 是强类型，故而可选参数必须声明出来；可选参数的声明方式有两种：

- 参数名后加 ? 表示值为 undefined；
- 声明参数默认值。

注意：

- 可选参数必须放到必须参数之后；
- 当声明了默认值后，不可再使用 ? 重复声明为可选参数。

```typescript
function func(name: string, age: number = 33, sex?: number) {
    console.log(name, age, sex);
}

func4("jack"); // jack 35 undefined
```

## Object

JS 中对象由属性和方构成，而 TS 中对象的类型就是在描述对象的结构。

写法：

```typescript
let obj: { name: string, age: number, printName(): string, printAge: () => number } = {
    age: 28,
    name: "tom",
    printAge(): number {
        return this.age;
    },
    printName(): string {
        return this.name;
    }
}

let obj2: {
    name: string,
    age: number,
    printName(): string,
} = {
    age: 22,
    name: "jack",
    printName(): string {
        return this.name;
    }
};
```

### 可选属性

由于 TS 是强类型，声明完属性或方法后，默认情况下实例化时必须为所有属性或方法进行赋值，通过声明为可选属性可以使其变为可选项，声明方法与函数可选参数一样使用 `?` 进行声明，默认值为 `undefined`;

```typescript
let obj: {
    name: string,
    age?: number,
    print?(): void,
};

obj = {
    name: "tom",
    print() {
        console.log("obj age:", this.age); // obj age: undefined
        return this.age;
    }
};
```

# 接口类型

当一个对象结构被多次使用时，一般使用接口（interface） 来描述对象结构，以达到复用和简化书写的目的; 声明接口使用关键字 `interface`。 接口类型与类型别名的区别在于，接口可以声明多次，表示合并； 接口可以用于被类继承。
具体细节看 oop.md
```typescript
interface MyObj {
    name: string,
    age?: number,

    print(): void
}

let obj3: MyObj = {
    name: "tom",
    print(): void {
        console.log(this.name, this.age); // tom undefined
    }
};
```

## 继承

接口之间继承达到继承属性的效果，格式 `interface MyInterFace extends MyInterFace2, MyInterFace3... {...}`

## 接口与类型别名的区别

- 接口只能用于对象类型；
- 类型别名可以用于任意类型。

# 联合类型

联合类型由多个类型组成，表示变量可以是其中任意一种，联合类型以 | 区分类型。

```typescript
let demo: string | number = 33;
demo = "jack";

let arr: (number | string)[] = [26, "jack"] // 联合类型
let arr4: number | string[] = ["tom", "jack"]; // 不带有小括号，表示变量可以是其中任意一种
arr4 = 26;
```

# 类型别名（自定义类型）

类型别名用于为任意类型起别名，当某些类型（书写很复杂）被多次使用时，可以通过类型别名来简化该类型的使用;声明类型别名使用 `type` 关键字来表示（type 是 TS 中的关键字）。

格式： `type MyType = type`;

```typescript
type MyType = (number | string)[]
let arr5: MyType = ["jack", 44];

type People = {
    name: string | number,
    age?: number | string,
    print(): void,
}

let p: People = {
    name: "pony",
    age: 22,
    print() {
        console.log(this.name, this.age);
    }
};

type myType = string | number
type myType2 = "A" | "B"
```

# 元组

元组是一种固定长度的数组类型，性能要优于数组。
`let param:[number, number | string] = [22, "jack"];`

# 枚举

```typescript
enum Sex {
    Man,
    Woman,
}

type People = {
    name: string,
    age: number,
    sex: Sex
}

let human: People = {
    name: "jack",
    age: 27,
    sex: Sex.Man
}
```

# void 类型

如果函数没有返回值，则返回 void 类型表示空值; 默认情况下，不声明返回类型且不`return`时为 void。

```typescript
function print(): void {
    console.log("void type");
}

// 等价 print
function print2() {
    console.log("void type");
}
```

# never 类型

never 类型用于函数返回值，表示永远不返回任何值，连 undefined 都不返回，在抛出错误的函数中会用到该类型。

# any, unknown 类型

- any 类型表示任意类型，使用该类型的变量， TS 会跳过类型检查，实际开发中原则上不去使用该类型。
- unknown 类型是安全的 any 类型，该类型变量可以赋予任意值，但是变量赋予其他变量时，则报错（与 any 的区别）。

```typescript
let k: unknown = "jack"
// 无法将 unknown 类型赋予其他类型变量，即便实际类型相同。
// let s:string = k

// 通过类型断言方式可以使 unknown 类型赋予其他类型变量
let s = k as string
k = 123

// 通过 typeof 进行类型判断也可以实现将 unknown 类型赋予其他类型变量
function f(param: unknown) {
    let s: string
    if (typeof param == "string") {
        s = param
    }
}


// any 类型可以赋予任意类型,规避掉 ts 的类型检查
let a: any = "jack"
let n: number = a
```

# 类型推论

在 TS 中，某些没有明确指出类型的地方， TS 的类型推论机制可以帮助提供类型；发生类型推论主要有两种场景：

- 声明变量并同时初始化赋值时；
- 函数不声明返回类型时。

备注：实际开发中能触发类型推论时优先触发类型推论，减少代码量。

```typescript
let userAge = 33; // 声明并初始化赋值时，出发类型推论
//  age ="jack" // 报错，类型推论为 number

// 不声明返回值类型时，触发类型推论
function printAge() {
    return userAge;
}
```

# 类型断言

当获取的类型过于宽泛而导致无法获取到需要的一些属性或方法时，可以通过类型断言来将获取的类型更加具体；

典型场景：

- 将父类断言为更加具体的子类；
- 将联合类型断言为其中一种类型；
- any 类型断言为具体类型；
- 类型断言为 any 类型。

断言有两种方式:

- 通过 `as`关键字进行断言：`type as type2`（常用的方式）;
- 通过 `<type2> type` 进行断言。

# 字面量类型

在 TS 中，JS 的基础类型中字符串、数字、布尔值的值也可以作为类型，被称之为字面量类型；通常字面量类型会配合联合类型一起使用，用于限制变量的值为其中之一；

注意：

- 类型推导时，变量由于值可变，所以被推论成相应类型，而常量由于值不可变，因此上述三种类型的常量被推导成字面量。

```typescript
const literal = () => {
    let name = "jack"; // name 是一个变量，因此被推导为 string 类型
    const name2 = "jack"; // name2 是一个常量，值不能变化，故而被推导成字面量类型 "jack"

    const printDirection = (direction: "up" | "down" | "left" | "right") => {
        console.log(direction);
    };
    printDirection("up");
};
```
