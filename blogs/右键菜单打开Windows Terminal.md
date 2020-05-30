# 右键菜单打开Windows Terminal

1. 进入任意文件夹，未选中任何内容时，单击右键弹出右键菜单显示打开Windows Terminal

1.1. reg1.bat

新建reg1.bat，内容如下，然后双击执行

Windows Registry Editor Version 5.00 [HKEY_CLASSES_ROOT\Directory\Background\shell\wt] @="Windows Terminal Here" "Icon"="%USERPROFILE%\\AppData\\Local\\Microsoft\\WindowsApps\\wt_32.ico" [HKEY_CLASSES_ROOT\Directory\Background\shell\wt\command] @="%USERPROFILE%\\AppData\\Local\\Microsoft\\WindowsApps\\wt.exe"

2. 选中任意文件夹，单击右键弹出右键菜单显示打开Windows Terminal

2.1. reg2.bat

新建reg2.bat，内容如下，然后双击执行

Windows Registry Editor Version 5.00 [HKEY_CLASSES_ROOT\Directory\shell\wt] @="Windows Terminal Here" "Icon"="%USERPROFILE%\\AppData\\Local\\Microsoft\\WindowsApps\\wt_32.ico" [HKEY_CLASSES_ROOT\Directory\shell\wt\command] @="%USERPROFILE%\\AppData\\Local\\Microsoft\\WindowsApps\\windows_terminal_here.bat"

2.2. windows_terminal_here.bat

新建windows_terminal_here.bat，建文件放置到%USERPROFILE%\AppData\Local\Microsoft\WindowsApps\文件夹下

%USERPROFILE% 具体对应的文件夹可打开CMD.exe，然后输入 echo %USERPROFILE%，按回车键查看输出内容

@echo off REM 此批处理用于右键菜单选择文件夹打开Windows Terminal start "wt" /D %1 "%USERPROFILE%\\AppData\\Local\\Microsoft\\WindowsApps\\wt.exe"

点击右键查看效果：

![img](https://static.ilibing.com/images/blogs/7c3150be8bcc384bbe8b875f25d94aa2.png)

![img](https://static.ilibing.com/images/blogs/f9068620e744e619a276f0531ac279e8.png)

![img](https://static.ilibing.com/images/blogs/e7830f056d0ce8ff924b09ef65d65e0c.png)

![img](https://static.ilibing.com/images/blogs/e8131199dab6d51f2383eab6182c7554.png)

![img](https://static.ilibing.com/images/blogs/c8456ad996f0db08e5784ba7bbb8d9bb.png)