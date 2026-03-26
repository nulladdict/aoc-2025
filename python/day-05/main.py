import asyncio
from operator import itemgetter
import aiofiles


async def main():
    async with aiofiles.open("input.txt") as f:
        contents = await f.read()
    ranges, ids = contents.split("\n\n")
    ranges = parse_ranges(ranges)
    ids = [int(i) for i in ids.splitlines()]
    print(f"{part1(ranges, ids)=}")
    print(f"{part2(ranges)=}")


type Ranges = list[tuple[int, int]]


def parse_ranges(lines: str) -> Ranges:
    results: list[tuple[int, int]] = []
    for line in lines.splitlines():
        s, e = line.split("-")
        results.append((int(s), int(e)))
    return results


def part1(ranges: Ranges, ids: list[int]) -> int:
    count = 0
    for id in ids:
        for s, e in ranges:
            if id >= s and id <= e:
                count += 1
                break
    return count


def part2(ranges: Ranges) -> int:
    ranges.sort(key=itemgetter(0))
    last = -1
    count = 0
    for s, e in ranges:
        start = max(last, s)
        end = max(last, e + 1)
        diff = end - start
        count += diff
        last = end
    return count


asyncio.run(main())
