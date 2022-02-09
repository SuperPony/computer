var age = 18;
var user = "tom";
var isLoading = true;
var un = undefined;
var n = null;
// @ts-ignore
var s = Symbol();
// array
var arr1 = [1, 2, 3]; // 推荐
var arr2 = ["tom", "jack"];
var arr3 = [26, "jack"]; // 联合类型，带有小括号表示值可以是任意类型
// function
function func1(name, age) {
    console.log(name, age);
    return name;
}
var func2 = function (name) {
    console.log(name);
    return name;
};
var func3 = function (name) { return name; };
// 可选参数
function func4(name, age, sex) {
    if (age === void 0) { age = 35; }
    console.log(name, age, sex);
}
func4("jack");
// object
var obj;
obj = {
    name: "tom",
    print: function () {
        console.log("obj age:", this.age);
        return this.age;
    }
};
obj.print();
var obj3 = {
    name: "tom",
    print: function () {
        console.log(this.name, this.age); // tom undefined
    }
};
obj3.print();
var wm = {
    name: "lucy",
    sex: 2,
    run: function () {
        console.log("i am running");
    },
    dance: function () {
        console.log("lucy is dancing");
    }
};
wm.dance();
// 联合类型
var demo = 33;
demo = "jack";
var arr4 = ["tom", "jack"];
arr4 = 26;
var arr5 = ["jack", 44];
var p = {
    name: "pony",
    age: 22,
    print: function () {
        console.log(this.name, this.age);
    }
};
p.print();
// 元组
var position = [34.5554, "116.3216"];
// void
function print() {
    console.log("void type");
}
// 等价 print
function print2() {
    console.log("void type");
}
// 类型推论
var userAge = 33; // 声明并初始化赋值时，出发类型推论
//  age ="jack" // 报错，类型推论为 number
// 不声明返回值类型时，触发类型推论
function printAge() {
    return userAge;
}
// 类型断言
function typeAssertion() {
    var link1 = document.getElementById("my-link");
    // link1 的类型为 HTMLElement，这种类型太过宽泛，只有 DOM 的一些通用属性，此处无法获取 a 标签的特有属性，可以通过类型断言来使得类型更加具体。
    console.log(link1);
    var link2 = document.getElementById("my-link");
    console.log(link2.href);
}
// 字面量
var literal = function () {
    var name = "jack"; // name 是一个变量，因此被推导为 string 类型
    var name2 = "jack"; // name2 是一个常量，值不能变化，故而被推导成字面量类型 "jack"
    var printDirection = function (direction) {
        console.log(direction);
    };
    printDirection("up");
};
