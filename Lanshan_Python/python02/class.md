# Python基础语法(2)

## 函数

### 1.1 函数的概念

>什么是函数？

**函数**:是一个 **`被命名的`** 、独立的、**`完成特定功能的代码段`**，其可能给调用它的程序一个 **`返回值`**。

通俗的说函数就是把一段可以`实现某种功能`的代码封装起来, 想要使用这个功能就可以直接调用函数。

**`被命名的`**：在Python中，大部分函数都是有名函数。

**`完成特定功能的代码段`**：函数的功能要专一，专门为了完成某个功能而定义。

**`返回值`**：当函数执行完毕后，其可能会返回一个值给函数的调用处。

函数的主要作用：
- ① 模块化编程
- ② 代码重用

函数中的几个重要概念：

- 函数
- 参数
- 返回值

函数语法：
```text
def 函数名([参数, ..]):
    代码1
    代码2
    ...
    [return 具体的返回值]
```
### 1.2 函数定义与调用

Python函数需要使用`def`关键字来定义。

函数在使用时，特点：

- （1）先定义，后调用；
- （2）不调用，不执行；
- （3）调用一次，执行一次。


调用函数：
```text
函数名()
```

`实例：`
（1）编写一个show()函数，并在函数中输出：bug虐我千百遍, 我待bug如初恋。

（2）调用函数，观察执行结果。
```python
  # 1.定义
def show():
    print("Hello World!")

# 2.调用
show()  # 1.写对函数名;   2.()

# 扩展
def show():
    # 代码
    pass
```

### 1.2 函数的参数
当在定义函数时，设定了参数，则可称该函数为：有参函数。反之，没有参数的函数，称为：无参函数。

定义有参数的函数，语法:

```text
def 函数名(参数1,参数2,...):  # 形参
	代码1
    代码2
    ...
```

调用函数，语法：
```text
函数名(参数值1,参数值2,...)  # 实参
```
**说明：**
- （1）形参是指形式参数，表示在定义函数时的参数；

- （2）实参是指实际参数，表示在调用函数时传递的参数值，具有实际意义。

**实例：**
（1）定义一个函数，用于求解两个数之和；

（2）接着，调用函数并执行程序，观察效果。
```python
# 1.定义函数
def get_sum(a,b):
    sum = a + b
    print(f"两个数之和为:{sum}")

# 2.调用函数
get_sum(10,20)
```
**总结：**
- （1）当定义了有参数的函数时，调用函数也需要传递参数值；

- （2）注意：当给有参函数传递参数时，要关注参数的三要素：个数、类型、顺序。

### 1.3 函数的返回值

函数的返回值指的是：当函数完成一件事情后，最后要返回给函数的结果。

返回值语法：
```text
def 函数名([参数1, 参数2, ...]):
	代码1
	代码2
	...
	return 值
```
- （1）若要给函数返回结果，需要使用return关键字；
- （2）return关键字的作用：把结果返回给函数；结束函数；
- （3）当函数没有返回值时，默认返回None。

**实例：**
```python
# 1.求解积
def get_multiply(a, b):
    ret = a * b
    return ret

print(get_multiply(10,20))   # print()
print('-'*50)

# 2.求差、和
def get_sub_sum(a, b):
    ret1 = a -b
    ret2 = a + b
    return ret1,ret2   # 元组
    
result = get_sub_sum(30,14)
print(result)
print('-'*50)

# 3.输出：(输出没有返回值)
def show():
    print("蓝山工作室")
print(show())
```

### 1.4 说明文档
说明文档指的是：在定义函数的第一行，加上多行注释。这样的操作，可以让函数更加规范化。

当添加了说明文档后，在调用函数时，就能查看到函数的作用提示。

当需要看函数的说明文档的具体内容时，语法：
```text
help(函数名)
```
### 2. 函数的嵌套调用
### 2.1 嵌套调用及执行流程
函数的嵌套调用指的是：在一个函数中，调用了另一个函数。

嵌套调用语法：
```text
def 函数1():
	代码
	...

def 函数2():
	代码
	# 调用函数1
	函数1()
	...
	# 在函数2中，调用了函数1。
```

```python
# 1.定义func()
def func():
    print("Hello World!\n")

# 2.定义test()、调用func()
def test():
    print("111")
    func()
    print("000")
    
# 3.调用test()
test() 
```

```python
# 1.定义函数
def line():
    print("-"*40)

# 2.N条横线
def print_lines(num):
    i = 0
    while i < num:
        # 嵌套调用
        line()
        i += 1

print_lines(10)
```


```python
import math


print(math.floor(32.9))
```

```python
from datetime import datetime

print(datetime.now())
```
#### 自己编写函数

```python
def add(a: int, b: int) -> int:
    return a + b


a = int(input("a="))
b = int(input("b="))
print("a+b={}".format(add(a, b)))
```

#### 内置函数

[Python 内置函数](https://docs.python.org/zh-cn/3/library/functions.html)

## 模块


### 1，什么是模块
在 Python 中，模块是一种组织 Python 代码的方法。 **模块可以包含定义（例如类、函数和变量）和可执行代码。** 

如果当我们要使用的功能在模块中时，我们就可以导入模块，使用模块中的相应功能。

也就是说：模块可以帮助你将代码组织得更有逻辑，并允许你将代码划分为更小的、更可管理的部分。

例如，我们想实现一些与时间相关的功能，我们就可以直接导入python中现有的time模块，然后调用time模块的功能帮我们实现。

### 2.模块的导入和使用
模块的导入一般写在文件的开头（这一点和C语言中，要包含对应的头文件类似，要写在开头）
导入的语法如下：
```text
 import 模块名 (可以一次导入多个模块，模块名之间用逗号隔开)
 from 模块名 import 类、变量、方法等
 from 模块名 import *
 import 模块名 as 别名
 from 模块名 import 功能名 as 别名
```
```python
import time  # 导入time模块
time.sleep(10)  # 通过 . 操作符使用功能

from time import sleep
from time import sleep  # 导入time模块的sleep功能
sleep(10)  # 直接使用功能

from time import sleep as al  # 导入time模块的sleep功能,给sleep设置别名：sl
al(10)  # 通过别名al来使用sleep功能
```
### 三，自定义模块
自定义模块其实就是创建一个以.py结尾的python文件，然后我们可以在这个文件里面写功能

（注意：自定义模块的命名要符合标识符的命名规则）

当我们需要调用这个文件里面的功能时，只需要导入这个模块就可以（模块名就是文件名）

例如，我们创建一个自定义模块:

使用它：
```python
import test  # 导入我们的自定义模块
test.ptint1()  # 使用里面的功能
```

2，同名功能的选择
注意：当导入多个不同的模块，但模块内有同名功能，当调用这个同名功能时，*调用的是后面导入的模块的功能*

3，*和 __all__

在Python中，使用 from module_name import * 表示导入模块中的所有内容。

这会导入模块中的所有变量、函数和类，我们可以在当前命名空间中直接使用它们。

（注意：这种方法可以快速地导入模块中的所有内容，但可能导致命名冲突，因此通常不推荐使用）

__all__:

__all__ 是模块中的一个特殊变量，它定义了模块中应该被导入的内容。


在Python中，__main__ 是一个特殊的模块，它主要用于执行模块中的main()函数。

当你直接运行一个Python文件时，这个文件中的代码会进入__main__模块。

换句话说，__main__模块是Python解释器直接运行的入口。

在Python中，一个模块的__name__属性会自动设置为’__main__'，

如果你直接运行一个模块，你可以通过这个属性来判断当前的模块是否是__main__模块。

当我们运行这条代码所写在的文件的时候：__name__ 就会直接被设置成 __main__

当我们运行别的文件的时候，__name__ 是不会被设置成 __main__ 的


#### 导入模块中的函数

[Python 标准库](https://docs.python.org/zh-cn/3/library/index.html)

## 文件处理
### 1 引言
在 Python 编程中，文件操作是一项基础且重要的技能，无论是数据分析、网络编程还是自动化脚本，都离不开对文件的读写操作。本文旨在全面介绍 Python 中文件处理的基本方法，包括文件的打开、读取、写入和关闭操作

2 文件打开 - open 函数
在 Python 中，使用内置的 open 函数来打开文件。open 函数的基本语法如下：
```python
file_object = open(file_name, mode)
```
- file_name：文件的路径和名称。
- mode：打开文件的模式。常用模式包括 'r'（只读），'w'（只写），'a'（追加），'r+'（读写），'b'（二进制模式）。
当使用 Python 的 open 函数打开文件时，mode 参数定义了文件应该以什么模式打开。不同的模式决定了文件如何被读取或写入。以下是 mode 参数的不同选项及其描述：

| 模式      |                        	描述                        |
|---------|:-------------------------------------------------:|
| `'r'`   |             只读模式。如果文件不存在，会发生异常。这是默认模式             |
| `'w'`   |          只写模式。如果文件存在，将被覆盖。如果文件不存在，创建新文件           |
| `'a'`   |     追加模式。如果文件存在，写入的数据将被追加到文件末尾。如果文件不存在，创建新文件      |
| `'r+'`  |           读写模式。可以读取和写入文件。如果文件不存在，会发生异常            |
| `'w+'`  |          读写模式。如果文件存在，将被覆盖。如果文件不存在，创建新文件           |
| `'a+'`  |     读写模式。如果文件存在，写入的数据将被追加到文件末尾。如果文件不存在，创建新文件      |
| `'x'`   |       独占写模式。如果文件已存在，操作将失败。只有在创建新文件时才使用这种模式        |
| `'b'`   |二进制模式。用于非文本文件如图像或二进制数据。可以与其他模式结合使用(如 `'rb'` 或 `'wb'`) |
| `'t'`   |文本模式（默认）。用于文本文件处理。可以与其他模式结合使用(如 `'rt'` 或 `'wt'`)|

可以根据需要选择适合的模式来打开文件。

例如，如果需要读取文件的内容，应该使用 `'r'` 模式。

如果想写入内容到一个新文件，并且不希望覆盖已有的文件，应该使用 `'x'` 模式。

如果打算处理二进制数据，如图片或者视频文件，应该选择 `'b'` 模式。

### 3 文件读取
### 3.1 读取整个文件
使用 `read()` 方法可以一次性读取整个文件内容
```python
with open('example.txt', 'r') as file:
    content = file.read()
    print(content)
```
### 3.2 按行读取
使用 `readline()` 或 `readlines()` 方法可以按行读取文件内容
```python
# 使用 readline()
with open('example.txt', 'r') as file:
    line = file.readline()
    while line:
        print(line, end='')
        line = file.readline()

# 使用 readlines()
with open('example.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        print(line, end='')
```

### 4 文件写入
### 4.1 写入文本
使用 `write()` 方法可以将文本写入文件：
```python
with open('example.txt', 'w') as file:
    file.write("Hello, Python!\n")
```
### 4.2 写入多行
使用 `writelines()` 方法可以一次性写入多行：
```python
lines = ["First line\n", "Second line\n"]
with open('example.txt', 'w') as file:
    file.writelines(lines)
```
### 4.3 文件关闭
在 Python 中，使用 `with` 语句打开文件是最佳实践，因为它会确保文件在操作完成后自动关闭。

如果使用传统的 `open()` 和 `close()` 方法，务必不要忘记调用 `close()` 来关闭文件。
```python
file = open('example.txt', 'r')
# 进行文件操作
file.close()
```
### 异常基础
在Python中，异常是一种特殊的对象，用于处理程序运行过程中发生的错误情况。

当Python解释器检测到一个错误，它会抛出（raise）一个异常。

如果这个异常没有被捕获（catch），程序就会终止并显示一个错误消息，即跟踪回溯（traceback）。

通过使用**try…except**语句块，我们可以捕获异常并采取适当措施。

代码示例：基本异常处理
```python
try:
    result = 10 / 0
except ZeroDivisionError:
    print("除数不能为0")
```
**核心特性和优势:**
- **增强程序健壮性：** 通过预测和处理潜在错误，确保程序不会意外终止。
- **提高代码可读性：** 明确地指出哪些部分代码可能会出错，并给出处理方案。
- **模块化错误处理：** 允许在不同层级捕获和处理异常，使得错误处理逻辑更加集中和模块化。

### 三、技术细节
**异常层次结构**

Python中的异常类形成了一个层次结构，位于BaseException之下，

常见的如Exception、ValueError等。了解这一结构有助于精准地捕获特定类型的异常。

**自定义异常**

开发者可以自定义异常类，继承自内置的异常类，以更好地表达特定业务逻辑中的错误情况。
```python
class MyCustomError(Exception):
    pass

def some_function():
    raise MyCustomError("发生了一个自定义错误")

try:
    some_function()
except MyCustomError as e:
    print(e)

```
**异常链**

当在异常处理块中再次抛出异常时，可以使用raise from语法保留原始异常信息，形成异常链，便于调试。

### 应用场景：文件操作
文件读取时，可能遇到文件不存在、权限不足等多种错误，合理使用异常处理可以优雅地应对这些问题。
```python
try:
    with open("example.txt", 'r') as file:
        content = file.read()
except FileNotFoundError:
    print("文件未找到，请检查路径是否正确。")
except PermissionError:
    print("没有足够的权限访问文件。")
else:
    print(content)
finally:
    print("文件操作完成，无论是否成功都会执行此段代码。")
```

### 五、优化与改进
**避免过度捕获**

仅捕获那些你确实能够处理的异常，避免使用过于宽泛的except Exception，这会捕获所有异常，可能隐藏真正的错误。

**使用上下文管理器和装饰器**

利用上下文管理器和装饰器可以进一步封装异常处理逻辑，使代码更加整洁
```python
from contextlib import contextmanager

def handle_file(path):
    try:
        f = open(path, 'r')
        yield f
    except FileNotFoundError:
        print(f"{path} not found.")
    finally:
        f.close()

with handle_file("example.txt") as file:
    content = file.read()
```

### 六、常见问题
**问题1：忽视异常信息**

解决方案：始终打印或记录异常的具体信息，以便于问题定位。

**问题2：过度使用except:捕获所有异常**

解决方案：尽量精确捕获异常类型，必要时才使用宽泛捕获，并提供合理的错误处理逻辑。
