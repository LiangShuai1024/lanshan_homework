n = int(input())
#print(n)
award = [[]for i in range(n)]
for i in range(n):
    award[i] = input().split(' ')
#print(award)
total = []
for h in range(n):
    name = award[h][0]
    final = int(award[h][1])
    talk = int(award[h][2])
    cadre = award[h][3]
    west = award[h][4]
    theis = int(award[h][5])
    #Dr=np.sign(award[h,2]-80)
    dr = 8000
    ws = 4000
    ach = 2000
    w = 1000
    c = 850
    if theis < 1 or final <= 80:
        dr = 0
    if final <= 85 or talk <= 80:
        ws = 0
    if final <= 90:
        ach = 0
    if west == "N" or final <= 85:
        w = 0
    if cadre == "N" or talk <= 80:
        c = 0
    t = dr+ws+ach+w+c
    total.append(t)
print(total)
m = max(total)
nm = award[total.index(m)][0]
s = sum(total)
print(nm, "\n", m, "\n", s)

