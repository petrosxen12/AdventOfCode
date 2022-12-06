def overlap_debug(func):
    def check_overlap(*args, **kwargs):
        value = func(*args, **kwargs)
        if value == 1:
            print(kwargs)
            print("=== Overlap detected")
        if value == 0:
            print(f"==No overlap {kwargs}")

        return value

    return check_overlap

@overlap_debug
def has_overlap(first_range, second_range):
    fr = [x for x  in set([int(x) for x in first_range.split("-")])]
    sr = [x for x in set([int(x) for x in second_range.split("-")])]

    fr.sort()
    sr.sort()

    # print(fr)
    # print(sr)
    # print(fr)
    # print(sr)
    if len(fr) == 1 and len(sr) == 1:
        if sr[0] == fr[0]:
            return 1 
        else:
            return 0

    if len(fr) == 1:
        if fr[0]>=sr[0] and fr[0]<=sr[1]:
            return 1
        else:
            return 0

    if len(sr) == 1:
        if sr[0]>=fr[0] and sr[0]<=fr[1]:
            return 1
        else:
            return 0

    if (fr[0] <= sr[0] ) and (fr[1]>= sr[0]) and (sr[1]>=fr[1]):
        return 1

    if (fr[0] >= sr[0] ) and (fr[1]<= sr[0]) and (sr[1]<=fr[1]):
        return 1   

    return 0

def full_containment(first_range, second_range):
    fr = [int(x) for x in first_range.split("-")]
    sr = [int(x) for x in second_range.split("-")]


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

            total_containments += has_overlap(first_range=sections_cleaning[0], second_range=sections_cleaning[1])
            # print("")
    print(total_containments)