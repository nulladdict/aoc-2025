import asyncio
import aiofiles


async def main():
    async with aiofiles.open("input.txt", "r") as f:
        data = await f.read()
    ranges = parse(data)
    print(f"{part1(ranges)=}")
    print(f"{part2(ranges)=}")


type Range = tuple[int, int]


def parse(data: str) -> list[Range]:
    ranges: list[tuple[int, int]] = []
    for pair in data.split(","):
        lower, upper = pair.split("-")
        ranges.append((int(lower), int(upper)))
    return ranges


def part1(ranges: list[Range]) -> int:
    sum = 0
    for lower, upper in ranges:
        for id in range(lower, upper + 1):
            target = str(id)
            length = len(target)
            if length % 2 == 1:
                continue
            mid = length // 2
            if target[:mid] == target[mid:]:
                sum += id
    return sum


def part2(ranges: list[Range]) -> int:
    sum = 0
    for lower, upper in ranges:
        for id in range(lower, upper + 1):
            target = str(id)
            max = len(target) // 2
            for factor in range(1, max + 1):
                if check(target, factor):
                    sum += id
                    break
    return sum


def check(target: str, factor: int) -> bool:
    if len(target) % factor != 0:
        return False
    pattern = target[:factor]
    for i in range(0, len(target), factor):
        if target[i : i + factor] != pattern:
            return False
    return True


asyncio.run(main())
