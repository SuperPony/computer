let age: number = 18;
let user: string = "tom";
let isLoading: boolean = true;
let un: undefined = undefined;
let n: null = null;
// @ts-ignore
let s: symbol = Symbol();

// array

let arr1: number[] = [1, 2, 3]; // 推荐
let arr2: Array<string> = ["tom", "jack"];
let arr3: (number | string)[] = [26, "jack"]; // 联合类型，带有小括号表示值可以是任意类型


// function
function func1(name: string, age: number): string {
    console.log(name, age);
    return name;
}

const func2 = (name: string): string => {
    console.log(name);
    return name;
};

let func3: (name: string) => string = (name) => name;

// 可选参数
function func4(name: string, age: number = 35, sex?: number) {
    console.log(name, age, sex);
}

func4("jack");


// object
let obj: {
    name: string,
    age?: number,
    print?(): void,
};

obj = {
    name: "tom",
    print() {
        console.log("obj age:", this.age);
        return this.age;
    }
};

obj.print();


// interface
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

obj3.print();

// 接口继承
interface Animal {
    sex: number,

    run(): void,
}

interface People {
    name: string;
}

interface Woman extends Animal, People {
    dance(): void;
}

let wm: Woman = {
    name: "lucy",
    sex: 2,
    run(): void {
        console.log("i am running");
    },
    dance() {
        console.log("lucy is dancing");
    }
};

wm.dance();


// 联合类型
let demo: string | number = 33;
demo = "jack";
let arr4: number | string[] = ["tom", "jack"];
arr4 = 26;


// 类型别名
type MyType = (number | string)[]
let arr5: MyType = ["jack", 44];

type Man = {
    name: string,
    age?: number | string,
    print(): void,
}

let p: Man = {
    name: "pony",
    age: 22,
    print() {
        console.log(this.name, this.age);
    }
};
p.print();

// 元组

let position: [number, number | string] = [34.5554, "116.3216"];

// void
function print(): void {
    console.log("void type");
}

// 等价 print
function print2() {
    console.log("void type");
}


// 类型推论
let userAge = 33; // 声明并初始化赋值时，出发类型推论
//  age ="jack" // 报错，类型推论为 number

// 不声明返回值类型时，触发类型推论
function printAge() {
    return userAge;
}


// 类型断言
function typeAssertion() {
    const link1 = document.getElementById("my-link");
    // link1 的类型为 HTMLElement，这种类型太过宽泛，只有 DOM 的一些通用属性，此处无法获取 a 标签的特有属性，可以通过类型断言来使得类型更加具体。
    console.log(link1);

    let link2 = document.getElementById("my-link") as HTMLAnchorElement;
    console.log(link2.href);
}


// 字面量
const literal = () => {
    let name = "jack"; // name 是一个变量，因此被推导为 string 类型
    const name2 = "jack"; // name2 是一个常量，值不能变化，故而被推导成字面量类型 "jack"

    const printDirection = (direction: "up" | "down" | "left" | "right") => {
        console.log(direction);
    };
    printDirection("up");
};

