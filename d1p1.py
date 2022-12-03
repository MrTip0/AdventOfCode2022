f = open("input.txt", "r")

massimo = 0
attuale = 0

lines = f.readlines()

for line in lines:
    line = line.strip()
    if line == "":
        if attuale > massimo:
            massimo = attuale
        attuale = 0

    else:
        tmp = int(line)
        attuale += tmp

print(massimo)


