import sys 

def main():
    input = sys.argv[1]
    result = factorial(int(input))
    print(result)

def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)


if __name__ == '__main__':
    main()
