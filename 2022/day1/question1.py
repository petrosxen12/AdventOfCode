if __name__ == "__main__":
    max = 0
    total_cal = 0  

    sums = []

    with open("input.txt") as f:
        for l in f.readlines():
            if l == "\n":
                sums.append(total_cal)

                # if total_cal > max:
                #     max = total_cal
                total_cal = 0
                # continue
            else:
                total_cal+=int(l)
            
        sums.sort(reverse=True)

        # print(max)
        
        top_three_sums = sum(sums[0:3],start=0)
        print(top_three_sums)