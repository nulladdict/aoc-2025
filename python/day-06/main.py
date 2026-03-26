import asyncio
from math import prod
import aiofiles


async def main():
    async with aiofiles.open("input.txt") as f:
        contents = await f.read()
    print(f"{part1(contents)=}")
    print(f"{part2(contents)=}")


def calc(operation: str, values: list[int]) -> int:
    match operation:
        case "*":
            return prod(values)
        case "+":
            return sum(values)
        case _:
            raise Exception("unreachable")


def part1(contents: str) -> int:
    acc = 0
    *lines, operations = contents.splitlines()
    numbers = [list(map(int, line.split())) for line in lines]
    operations = operations.split()
    for i, op in enumerate(operations):
        values = [num[i] for num in numbers]
        acc += calc(op, values)
    return acc


def parse_number_column(lines: list[str], at: int) -> int:
    base = 0
    for line in lines:
        char = line[at]
        if char == " " and base != 0:
            return base
        digit = 0 if char == " " else int(char)
        base = 10 * base + digit
    return base


def part2(contents: str) -> int:
    acc = 0
    *lines, operations = contents.splitlines()
    operations = operations.split()
    values: list[int] = []
    for i in range(len(lines[0]) - 1, -1, -1):
        num = parse_number_column(lines, i)
        if num != 0:
            values.append(num)
        else:
            op = operations.pop()
            acc += calc(op, values)
            values.clear()
    op = operations.pop()
    acc += calc(op, values)
    return acc


asyncio.run(main())
