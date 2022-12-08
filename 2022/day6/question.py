"""
your subroutine needs to identify the first position where the four most recently received characters were all different. 
Specifically, it needs to report the number of characters from the beginning of the buffer to the end of the first such four-character marker.
"""
DISTINCT_CHARACTER_SELECTION = 14

def is_unique(window):
    # print(window)
    if len(set(window)) < DISTINCT_CHARACTER_SELECTION:
        return False

    return True

if __name__ == "__main__":
   
    with open("input.txt", encoding="utf-8") as f:
        data_stream = f.readline()
        # data_stream = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
        search_window = DISTINCT_CHARACTER_SELECTION
        # print(data_stream)
        for index, character in enumerate(data_stream):
            chars_until_uniqueness = search_window+index
            if is_unique(data_stream[index:search_window+index]):
                print(chars_until_uniqueness)
                break
