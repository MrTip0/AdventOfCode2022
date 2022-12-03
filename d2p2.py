f = open("input.txt")

lines = f.readlines()
r = 0


for line in lines:
    line = line.strip()
    n1, n2 = 0, 0 
    s1, s2 = line.split(" ")[0], line.split(" ")[1]
    match s1:
        case "A":
            n1 = 0
        case "B":
            n1 = 1
        case "C":
            n1 = 2

    match s2:
        case "X":
            n2 = (n1 + 2) % 3
        case "Y":
            n2 = n1
        case "Z":
            n2 = (n1 + 1) % 3

    r += n2 + 1
    if (n2 + 2)%3 == n1:
        r += 6
    elif n2 == n1:
        r += 3
print(r)
