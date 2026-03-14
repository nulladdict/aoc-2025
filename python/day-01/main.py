import asyncio
import aiofiles


async def main():
    async with aiofiles.open("input.txt", "r") as f:
        data = await f.read()
    rotations = str(data).splitlines()
    offsets = parse(rotations)
    print(part1(offsets))
    print(part2(offsets))


def parse(rotations: list[str]) -> list[int]:
    return [(1 if r[0] == "R" else -1) * int(r[1:]) for r in rotations]


def part1(offsets: list[int]) -> int:
    dial = 50
    count = 0
    for offset in offsets:
        dial += offset
        dial %= 100
        if dial == 0:
            count += 1
    return count


def part2(offsets: list[int]) -> int:
    dial = 50
    count = 0
    for offset in offsets:
        dial += offset
        count += abs(dial) // 100
        if dial == 0 or (dial < 0 and dial != offset):
            count += 1
        dial %= 100
    return count


asyncio.run(main())
