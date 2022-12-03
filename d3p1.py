f = open("input.txt")

lines = f.readlines()

r = 0

for line in lines:
    line = line.strip()
    s1, s2 = line[:int(len(line)/2)], line[int(len(line)/2):]
    repeted: set[int] = set()
    for s in s1:
        if s in s2:
            if s >= 'a':
                v = ord(s) - 96
            else:
                v = ord(s) - 65 + 27
            repeted.add(v)

    for el in repeted:
        r += el
print(r)

