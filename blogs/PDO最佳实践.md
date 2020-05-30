# PDO最佳实践



## 1. PDO简介

PHP 数据对象 （PDO） 扩展为PHP访问数据库定义了一个轻量级的一致接口。实现 PDO 接口的每个数据库驱动可以公开具体数据库的特性作为标准扩展功能。 注意利用 PDO 扩展自身并不能实现任何数据库功能；必须使用一个 具体数据库的 PDO 驱动 来访问数据库服务。

PDO 提供了一个 数据访问 抽象层，这意味着，不管使用哪种数据库，都可以用相同的函数（方法）来查询和获取数据。 PDO 不提供 数据库 抽象层；它不会重写 SQL，也不会模拟缺失的特性。如果需要的话，应该使用一个成熟的抽象层。

从 PHP 5.1 开始附带了 PDO，在 PHP 5.0 中是作为一个 PECL 扩展使用。 PDO 需要PHP 5 核心的新 OO 特性，因此不能在较早版本的 PHP 上运行。



## 2. 连接数据库



### 2.1. 建立连接

```php
$pdo = new PDO('mysql:host=localhost;port=3306;dbname=test', $username, $password, $options);
```

### 2.2. 推荐连接选项

> 参照：https://www.php.net/manual/zh/pdo.setattribute.php

```php
$options = [
    // 强制列名为指定的大小写 
    PDO::ATTR_CASE => PDO::CASE_NATURAL,
    // 错误报告
    PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
    // 转换 NULL 和空字符串
    PDO::ATTR_ORACLE_NULLS => PDO::NULL_NATURAL,
    // 提取的时候将数值转换为字符串
    PDO::ATTR_STRINGIFY_FETCHES => false,
    // 是否启用本地预处理语句的模拟
    PDO::ATTR_EMULATE_PREPARES => false,
];
$pdo = new PDO('mysql:host=localhost;port=3306;dbname=test', $username, $password, $options);
```

### 2.3. 连接错误异常捕获

```php
try {
    $pdo = new PDO('mysql:host=localhost;port=3306;dbname=test', $username, $password, $options);
} catch (PDOException $e) {
    echo 'Failed to connect database: ' . $e->getMessage();
    exit;
}
```

### 2.4. 关闭连接

```php
$pdo = null;
```

### 2.5. 持久化连接（不推荐）

很多 web 应用程序通过使用到数据库服务的持久连接获得好处。持久连接在脚本结束后不会被关闭，且被缓存，当另一个使用相同凭证的脚本连接请求时被重用。持久连接缓存可以避免每次脚本需要与数据库回话时建立一个新连接的开销，从而让 web 应用程序更快。

```php
// 如果想使用持久连接，必须在传递给 PDO 构造函数的驱动选项数组中设置 PDO::ATTR_PERSISTENT
$options = [
    PDO::ATTR_PERSISTENT => true
];
$pdo = new PDO('mysql:host=localhost;port=3306;dbname=test', $username, $password, $options);
```



## 3. 执行SQL



### 3.1. 事务

```php
// 开启一个事务
$pdo->beginTransaction();

// 提交一个事务
$pdo->commit();

// 回滚一个事务
$pdo->rollBack();
```

### 3.2. 执行预处理语句

```php
// 要执行的SQL语句
$sql = 'select * from user where id = ? limit 1';

// 绑定参数
$args = [1];

// 执行$sql返回一个预处理语句
$stmt = $pdo->prepare($sql);

// 执行预处理语句
if ($stmt instanceof PDOStatement) {
    $stmt->execute($args);
}

// 取出一个包含查询数据行的数组（适用于select语句）
if ($stmt instanceof PDOStatement) {
    $users = $stmt->fetchAll(PDO::FETCH_ASSOC);
}
```



## 4. 错误处理



PDO 提供了三种不同的错误处理模式，以满足不同风格的应用开发：

- **PDO::ERRMODE_SILENT**

此为默认模式。 PDO 将只简单地设置错误码，可使用 PDO::errorCode() 和 PDO::errorInfo() 方法来检查语句和数据库对象。如果错误是由于对语句对象的调用而产生的，那么可以调用那个对象的 PDOStatement::errorCode() 或 PDOStatement::errorInfo() 方法。如果错误是由于调用数据库对象而产生的，那么可以在数据库对象上调用上述两个方法。

- **PDO::ERRMODE_WARNING**

除设置错误码之外，PDO 还将发出一条传统的 E_WARNING 信息。如果只是想看看发生了什么问题且不中断应用程序的流程，那么此设置在调试/测试期间非常有用。

- **PDO::ERRMODE_EXCEPTION**

除设置错误码之外，PDO 还将抛出一个 PDOException 异常类并设置它的属性来反射错误码和错误信息。此设置在调试期间也非常有用，因为它会有效地放大脚本中产生错误的点，从而可以非常快速地指出代码中有问题的潜在区域（记住：如果异常导致脚本终止，则事务被自动回滚）。

异常模式另一个非常有用的是，相比传统 PHP 风格的警告，可以更清晰地构建自己的错误处理，而且比起静默模式和显式地检查每种数据库调用的返回值，异常模式需要的代码/嵌套更少。

**推荐使用：PDO::ERRMODE_EXCEPTION**

设置方法如下：

```php
$options = [
    PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION
];
$pdo = new PDO('mysql:host=localhost;port=3306;dbname=test', $username, $password, $options);
```