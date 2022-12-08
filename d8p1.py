matrix:list[list[int]] = []
visible:list[list[bool]] = []
f = open("input.txt", "r")

lines = f.readlines()

for line in lines:
    line = line.strip()
    list = []
    bool = []
    for c in line:
        bool.append(False)
        list.append(int(c))
    matrix.append(list)
    visible.append(bool)
    
for i, line in enumerate(matrix):
    max = -1
    for j, n in enumerate(line):
        if n > max:
            max = n
            visible[i][j] = True

    max = -1
    for j in range(len(line)-1, -1, -1):
        if line[j] > max:
            max = line[j]
            visible[i][j] = True

l = len(matrix[0])
ll = len(matrix)

for j in range(l):
    max = -1
    for i in range(ll):
        if matrix[i][j] > max:
            max = matrix[i][j]
            visible[i][j] = True

    max = -1
    for i in range(ll-1, -1, -1):
        if matrix[i][j] > max:
            max = matrix[i][j]
            visible[i][j] = True

r = 0
for line in visible:
    for v in line:
        if v:
            r += 1
print(r)