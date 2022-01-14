# Pflag

Pflag 是用于设置和解析命令行参数的一个成熟的第三方库；

# Index
- 术语
- Flag
  - 设置 flag；
  - 隐藏、废弃 flag；
  - flag struct。
- FlagSet
  - 设置 FlagSet。
- 其他


# 术语

`git clone URL --bare`:
  - clone 称之为命令；
  - URL 称之为非选项参数（非标志参数）；
  - --bare 称之为选项参数（标志参数）。

# Flag

flag 用于描述注册在命令行中的标志参数。

## 设置 flag

- `TypeP(name string, shorthand string, value string, usage string)*Type`: 用于设置指定名称的 flag， value 参数表示在不指定该标志时的默认值。
- `Type(name string, value string, usage string)*Type`: 与 TypeP 相似，区别在于没有声明；
- `TypeVar(p *type, name string, shorthand string, value uint8, usage string)`: 使用对应类型变量作为接收 flag 值的声明方式；

- `Lookup("flag").NoOptDefVal = "value"`: 当该 flag 作为了指定的选项，但是又没有输入值时，使用的默认值（不设置该字段值时，当 flag 作为指定选项，没有对应输入值时，报错）。


## 隐藏、废弃 flag

- `CommandLine.MarkHidden(name string)`: 隐藏指定标志，可以正常使用，但不会出现在帮助文档中；
- `CommandLine.MarkDeprecated(name string, usageMessage string) error`:用于废弃指定标志、短声明，废弃的标志不会出现在帮助文档中，但是依然可以正常的使用，当使用废弃的标志时，会出输出携带 usageMessage 的警告提醒。

## flag struct

flag struct 保存着对应的标志的全部信息;
获取指定结构体 `pflag.Lookup(name string) *pflag.Flag`;
结构体说明如下：

```
// type Flag struct {
    // 	Name string // flag长选项的名称
    // 	Shorthand string // flag短选项的名称，一个缩写的字符
    // 	Usage string // flag的使用文本
    // 	Value Value // flag的值
    // 	DefValue string // flag的默认值
    // 	Changed bool // 记录flag的值是否有被设置过
    // 	NoOptDefVal string // 当flag出现在命令行，但是没有指定选项值时的默认值
    // 	Deprecated string // 记录该flag是否被放弃
    // 	Hidden bool // 如果值为true，则从help/usage输出信息中隐藏该flag
    // 	ShorthandDeprecated string // 如果flag的短选项被废弃，当使用flag的短选项时打印该信息
    // 	Annotations map[string][]string // 给flag设置注解
// }
```

其中 Value 是一个接口，表示着 flag 的值；接口说明如下：

```
// type Value interface {
    // 	String() string   // 将flag类型的值转换为string类型的值，并返回string的内容
    // 	Set(string) error // 将string类型的值转换为flag类型的值，转换失败报错
    // 	Type() string     // 返回flag的类型，例如：string、int、ip等
// }
```


# FlagSet

FlagSet 是一组标志的集合；即便是直接使用 Flag 进行注册，其本质上也是在全局的标志集合中添加了一个标志。

```
pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit.")

func BoolVarP(p *bool, name, shorthand string, value bool, usage string) { 
    flag := CommandLine.VarPF(newBoolValue(value, p), name, shorthand, usage) 
    flag.NoOptDefVal = "true"
}

// CommandLine 是一个包级别变量
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

## FlagSet 常用方法

- `NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet`: 设置指定集合名称的标志集合；
- `func (f *FlagSet) LookUp(name string) *Flag`: 返回集合下的指定 flag；
- `func (f *FlagSet) AddFlagSet(newSet *FlagSet)`: 注入指定的 FlagSet；
- `func (f *FlagSet) AddGoFlagSet(newSet *goflag.FlagSet)`: 将 flag 包内 FlagSet 注入；
- `func (f *FlagSet) AddGoFlag(goflag *goflag.Flag)`: 注入一个 flag 包内 Flag；
- `func (f *FlagSet) SetNormalizeFunc(n func(f *FlagSet, name string) NormalizedName)`: 设置一个转换选项参数的函数，常用于将非标准格式的选项参数转为标准格式。
```
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

flagSetDemo.SetNormalizeFunc(WordSepNormalizeFunc)
```

# 其他


- `Parse()`: 注册选项参数；
- `Args()`: 返回非选项参数集合切片；
- `Lookup(name string) *Flag`: 返回指定名称的 Flag。