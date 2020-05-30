#  PHP5.6升级到PHP7.1带来的新特性



## 1. 概述

> PHP7将在2015年10月正式发布，PHP7是PHP编程语言全新的一个版本，主要在性能方面获得了极大的提升。官方的文档显示，PHP7可以达到PHP5.x版本两倍的性能。同时还对PHP的语法做了梳理，提供了很多其他语言流行的语法格式。难能可贵的是，做了如此大的改动，PHP7的兼容性还是非常好的。对于绝大多数的应用来讲，可以不用做修改即可迁移到PHP7版本。



## 2. 性能提升



### 2.1. Web站点服务器端编程语言百分比

> 注：一个web站点可以会使用多种语言作为它的开发语言 [数据来源于W3Techs.com(2018.11.15)](https://w3techs.com/technologies/overview/programming_language/all)

![Web站点服务器端编程语言百分比](https://i.loli.net/2018/11/15/5bed10c08c713.png?ynotemdtimestamp=1590716616551)

### 2.2. Benchmark对比

> PHP7的性能测试结果，性能压测结果，耗时从2.991下降到1.186，大幅度下降60%。

![image](http://box.kancloud.cn/2015-09-16_55f901728bce2.png?ynotemdtimestamp=1590716616551)

### 2.3. WordPress的QPS压测

> 在WordPress项目中，PHP7对比PHP5.6，QPS提升2.77倍。

![image](http://box.kancloud.cn/2015-09-16_55f90177cc877.png?ynotemdtimestamp=1590716616551)

> 附录：[PHP5.6-7.2性能比较](https://segmentfault.com/a/1190000013690281)



## 3. 新特性



### 3.1. 强制模式和严格模式

#### 3.1.1 强制模式

> 强制模式是默认模式，强制模式下，它会帮我们把数字类型的string类型，int整型，bool，强制类型。其他类型不能转换，就会报错。

```
function test() : int {
    return '1';
}
var_dump(test()); // 返回值被转成int类型，输出：int(1)

function test() : int {
    return [1];
}
var_dump(test()); // 类型不能转换，报错：Uncaught TypeError: Return value of test() must be of the type integer...
```

#### 3.1.2. 严格模式

> 严格模式下，参数必须符合规定，不然就会报错。

```
declare(strict_types=1); // 声明指令必须放在文件的顶部

function test() : int {
    return 1;
}
var_dump(test()); // 输出：int(1)

function test() : int {
    return '1';
}
var_dump(test()); // 报错：Uncaught TypeError: Return value of test() must be of the type integer...
```

### 3.2. 标量类型声明

#### 3.2.1. 指定标量类型

```
function test(int $a)
{
    var_dump($a);
}
test(1); // 输出：int(1)
test() // 报错：Uncaught TypeError: Argument 1 passed to test() must be of the type integer, string given...
```

#### 3.2.2. 可为空（Nullable）返回类型

```
function test(?string $name)
{
    var_dump($name);
}
test('Hello PHP'); // 输出：string(9) "Hello PHP"
test(null); // 输出：NULL
test(); // 报错：Uncaught ArgumentCountError: Too few arguments to function test()...
```

### 3.3. 返回值类型声明

#### 3.3.1. 指定返回类型

```
function test(): int 
{
    return 123;
}
var_dump(test()); // 输出：int(123)


function test(): int
{
    return '123';
}
var_dump(test()); // 输出：int(123)，P.s. 字符串会自动转为整型返回，反亦然，严格模式下则报返回类型错误

function test(): int
{
    return [123];
}
var_dump(test()); // 报错：Uncaught TypeError: Return value of test() must be of the type integer, array returned
```

#### 3.3.2. 可为空（Nullable）返回类型

```
function test(): ?int
{
    return 1;
}
var_dump(test()); // 输出：int(1)


function test(): ?int
{
    return null;
}
var_dump(test()); // 输出：NULL

function test(): ?int
{
    return;
}
var_dump(test()); // 报错：A function with return type must return a value
```

#### 3.3.3. Void 类型

```
function test(): void
{
    return 1;
}
test(); // 报错：A void function must not return a value


function test(): void
{
    return null;
}
test(); // 报错：A void function must not return a value

function test(): void
{
    return;
}
test(); // 正常运行，无报错
```

### 3.4. null合并运算符

###### 由于日常使用中存在大量同时使用三元表达式和 isset()的情况， 我们添加了null合并运算符 (??) 这个语法糖。如果变量存在且值不为NULL， 它就会返回自身的值，否则返回它的第二个操作数。

```
// 如果 $_GET['user'] 不为 NULL, 等于 $_GET['user']，否则等于'nobody'
$username = $_GET['user'] ?? 'nobody';

// 上面表达式等同于此表达式
$username = isset($_GET['user']) ? $_GET['user'] : 'nobody';

// 如果 $_GET['user'] 不为 NULL, 等于 $_GET['user']，
// 如果 $_GET['user'] 为NULL, $_POST['user'] 不为NULL, 等于 $_POST['user']
// 否则等于'nobody'
$username = $_GET['user'] ?? $_POST['user'] ?? 'nobody';
```

### 3.5. 太空船操作符

###### 太空船操作符用于比较两个表达式。当$a小于、等于或大于$b时它分别返回-1、0或1。

```
echo 1 <=> 1; // 0
echo 1 <=> 2; // -1
echo 2 <=> 1; // 1
```

### 3.6. 通过 define() 定义常量数组

#### Array 类型的常量现在可以通过 define() 来定义。在 PHP5.6 中仅能通过 const 定义。

```
define('ANIMALS', [
    'dog',
    'cat',
    'bird'
]);
```

### 3.7. 类常量可见性

#### 现在起支持设置类常量的可见性。

```
class ConstDemo
{
    const PUBLIC_CONST_A = 1;
    public const PUBLIC_CONST_B = 2;
    protected const PROTECTED_CONST = 3;
    private const PRIVATE_CONST = 4;
}
```

### 3.8. 多异常捕获处理

###### 一个catch语句块现在可以通过管道字符(|)来实现多个异常的捕获。 这对于需要同时处理来自不同类的不同异常时很有用。

```
try {
    // some code
} catch (FirstException | SecondException $e) {
    // handle first and second exceptions
}
```

------

> 部分素材来源：
> http://www.php7.site/page/about-php7.html 
> https://w3techs.com/ 
> http://php.net/ http://hansionxu.blog.163.com/blog/static/24169810920158704014772/