f = open("input.txt", "r")

m1, m2, m3 = 0, 0, 0
attuale = 0

lines = f.readlines()

for line in lines:
    line = line.strip()
    if line == "":
        if attuale > m1:
            m1, m2, m3 = attuale, m1, m2
        elif attuale > m2:
            m2, m3 = attuale, m2
        elif attuale > m3:
            m3 = attuale
        attuale = 0

    else:
        tmp = int(line)
        attuale += tmp

print(m1 + m2 + m3)


