def main():
    word = input("Enter the string ")
    chars = {}
    for char in word: 
        if char in chars:
            print("{} is not an unique string".format(word))
            return
        else:
            chars[char] = 1
    print("{} is an unique string".format(word))


if __name__ == '__main__':
    main()
