a=int(input("请输入一个整数："))
b=str(a)
b=b.replace("-","")
b=int(b[::-1])
if a<0:
    b=-b
print("反转后得：",b)



