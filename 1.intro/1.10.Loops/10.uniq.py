def python(a: str, b: str) -> None:
    to_set = lambda i: set([char for char in i])
    a, b = to_set(a), to_set(b)
    inter = a.intersection(b)
    print("Python", " ".join(list(inter)), end="\n")


def golang(a: str, b: str) -> None:
    result = ""
    for char in a:
        if char in b and char not in result:
            result += f"{char} "
    print("Go", result, end="\n")


def main():
    a, b = input(), input()
    python(a, b)
    golang(a, b)


if __name__ == "__main__":
    main()
