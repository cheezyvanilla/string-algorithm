def rreverse(s):
    if s == "":
        return s
    else:
        print(s[1:])
        return rreverse(s[1:]) + s[0]

a = rreverse("jangkrik")
print("hasil: ", a)