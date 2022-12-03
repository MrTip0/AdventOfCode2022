f = open("input.txt")

lines = f.readlines()

r = 0

for i in range(0, len(lines), 3):
    line1 = lines[i].strip()
    line2 = lines[i+1].strip()
    line3 = lines[i+2].strip()

    for el in line1:
        if el in line2 and el in line3:
            if el >= 'a':
                v = ord(el) - 96
            else:
                v = ord(el) - 65 + 27
            r += v
            break
print(r)

