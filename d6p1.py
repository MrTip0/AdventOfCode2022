f = open("input.txt", "r")

r = 0

s = f.read()

k = 4

for i, v in enumerate(s):
    if i < k:
        continue
    not_found = False
    
    used:set[str] = set()
    for j in range(i, i - k, -1):
        if s[j] in used:
            not_found = True
            break
        else:
            used.add(s[j])
    
    if not not_found:
        r = i
        break
    


print(r + 1)

