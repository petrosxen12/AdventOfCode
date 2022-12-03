if __name__ == "__main__":
    # 1. Name each sum of calories
    # 2. Sum each section
    # 3. Find highest number of sums
    # - Sum the lines until you find new line character
    # - Get results of sum and order them in descending order

    elf_calorie_carriers = open("input.txt")

    sums = []
    total_calories = 0

    for calorie in elf_calorie_carriers.readlines():
        
        if calorie == '\n':
            sums.append(total_calories)
            # print(total_calories)
            total_calories = 0

        else:
            total_calories +=  int(calorie)

    sums.sort(reverse=True)

    print(sums[0:3])
    top_three_elves=sum(sums[0:3],start=0)
    print(top_three_elves)