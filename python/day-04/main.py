import asyncio
from collections.abc import Generator
import aiofiles


async def main():
    async with aiofiles.open("input.txt", "r") as f:
        contents = await f.read()
    grid = contents.splitlines()
    rolls = parse(grid)
    print(f"{part1(rolls)=}")
    print(f"{part2(rolls)=}")


def parse(grid: list[str]) -> set[complex]:
    rolls = set[complex]()
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == "@":
                rolls.add(complex(x, y))
    return rolls


def part1(rolls: set[complex]) -> int:
    count = 0
    for roll in rolls:
        length = sum(1 for _ in neigbours(rolls, roll))
        if length < 4:
            count += 1
    return count


def neigbours(rolls: set[complex], p: complex) -> Generator[complex, None, None]:
    deltas = [
        1 + 1j,
        1 + 0j,
        1 - 1j,
        0 + 1j,
        # 0 + 0j,
        0 - 1j,
        -1 + 1j,
        -1 + 0j,
        -1 - 1j,
    ]
    for d in deltas:
        neigbour = p + d
        if neigbour in rolls:
            yield neigbour


def part2(rolls: set[complex]) -> int:
    rolls = rolls.copy()
    count = 0
    while True:
        remove = set[complex]()
        for roll in rolls:
            length = sum(1 for _ in neigbours(rolls, roll))
            if length < 4:
                remove.add(roll)
        if len(remove) == 0:
            return count
        rolls.difference_update(remove)
        count += len(remove)


asyncio.run(main())
