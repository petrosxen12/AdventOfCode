def full_containment(first_range, second_range):

    fr = [int(x) for x in first_range.split("-")]
    sr = [int(x) for x in second_range.split("-")]

    # print(fr)
    # print(sr)

    if fr[0]>=sr[0] and fr[1]<=sr[1]:
        return 1
    if sr[0]>=fr[0] and sr[1]<=fr[1]:
        return 1

    return 0

if __name__ == "__main__":
    total_containments = 0

    with open("input.txt") as f:

        for pairs in f.readlines():
            sections_cleaning = pairs.strip().split(",")

            total_containments += full_containment(first_range=sections_cleaning[0], second_range=sections_cleaning[1])

    print(total_containments)