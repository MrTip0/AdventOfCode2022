matrix:list[list[int]] = []
f = open("input.txt", "r")

lines = f.readlines()
r = 0

for line in lines:
    line = line.strip()
    if line == "":
        continue
    list = []
    for c in line:
        list.append(int(c))
    matrix.append(list)

l = len(matrix[0])
ll = len(matrix)

for i, line in enumerate(matrix):
    for j, el in enumerate(line):
        a, b, c, d = 0, 0, 0, 0
        v = el
            
        for h in range(i-1, -1, -1):
            a+=1
            if matrix[h][j] >= v:
                break
            
        for h in range(i+1, ll):
            b+=1
            if matrix[h][j] >= v:
                break

        for h in range(j-1, -1, -1):
            c += 1
            if line[h] >= v:
                break

        for h in range(j+1, l):
            d += 1
            if line[h] >= v:
                break
        v = a * b * c * d
        if v > r:
            r = v
print(r)