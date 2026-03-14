import asyncio
import aiofiles


async def main():
    async with aiofiles.open("input.txt", "r") as f:
        contents = await f.read()
    banks = contents.splitlines()
    print(f"{part1(banks)=}")
    print(f"{part2(banks)=}")


def part1(banks: list[str]) -> int:
    sum = 0
    for bank in banks:
        joltage = 0
        for i in range(0, len(bank) - 1):
            for j in range(i + 1, len(bank)):
                joltage = max(joltage, int(bank[i] + bank[j]))
        sum += joltage
    return sum


def part2(banks: list[str]) -> int:
    sum = 0
    for bank in banks:
        joltage = 0
        start = 0
        for power in range(12, 0, -1):
            end = len(bank) - (power - 1)
            target = bank[start:end]
            (idx, digit) = max_digit(target)
            start += idx + 1
            joltage += digit * (10 ** (power - 1))
        sum += joltage
    return sum


def max_digit(target: str) -> tuple[int, int]:
    max = -1
    idx = 0
    for i, diget in enumerate(target):
        d = int(diget)
        if d > max:
            max = d
            idx = i
    return idx, max


asyncio.run(main())
