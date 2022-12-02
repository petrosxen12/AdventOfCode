if __name__ == "__main__":
    max = 0
    sum = 0  
    with open("input.txt") as f:
        for l in f.readlines():
            if l == "\n":
                if sum > max:
                    max = sum
                sum = 0
                continue
            else:
                sum+=int(l)

        print(max)