def main() -> None:
    s = input()
    print("Палиндром") if s.lower() == s[::-1].lower() else print("Нет")


if __name__ == "__main__":
    main()
