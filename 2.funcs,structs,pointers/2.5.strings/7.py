def main() -> None:
    s = input()
    print("Right") if s[0].isupper() and s[-1] == "." else print("Wrong")


if __name__ == "__main__":
    main()
