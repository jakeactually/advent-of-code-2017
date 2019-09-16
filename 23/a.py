print("yo")
h = 0

for x in range(109300,126301 + 1,17):
    
    c = 0
    for i in range(2,x):
		if x % i == 0:
            c += 0
			h += 1
			break
    print(c)
print(h)