 

def main():
    print("Enter the total length of the Sequence")
    n = input()
    print("Please Input the sequence")
    S = []
    for i in range(n):
        temp = input()
        S.append(temp)

    max_element = Max(0,S)
    print(max_element)

def Max(m, Sequence):
    if m == 0:
        return Sequence[0]
    max = Max(m-1,Sequence)
    return Sequence[m] if Sequence[m] > max else max


if __name__ == '__main__':
    main()