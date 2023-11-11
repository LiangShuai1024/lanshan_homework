add=0
lst = list(input("请输入两以空格分开的整数：").split(" "))
for i in range(len(lst)):
    lst[i] = int(lst[i])
for a in range(0,len(lst)):
    add = add+lst[a]
print("它们的和为：",add)
