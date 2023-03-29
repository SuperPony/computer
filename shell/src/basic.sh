# 第一行是一个约定标记，告诉脚本使用什么解释器来执行。
#!/bin/bash


# 只读变量在初始化后，无法再次赋值
readonly userName="jack";
echo ${userName}

cars=("Benz" "Audi" "BMW");
el=${cars[0]};
echo ${el};
# 获取所有元素
echo ${cars[@]};
# 获取元素个数
echo ${#cars[@]};

# 在执行 shell 脚本时，可以携带参数，在脚本中通过 $n 进行获取，n 表示传入的第几个参数。
echo $1;

echo $?

date > ./demo.txt