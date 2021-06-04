
target=app
#获取当前目录下的所有.c文件
src=$(wildcard ./*.c)
#把所有.c替换成.o
obj=$(patsubst ./%.c,./%.o,$(src))
$(target):$(obj)
        gcc test.o -o app

%.o:%.c
        gcc -c $< -o $@

#声明伪目标
.PHONY:clean
clean:
        rm $(obj) $(target) -f
