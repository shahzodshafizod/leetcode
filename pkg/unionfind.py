class UnionFind:
    def __init__(self, size: int):
        self.root = [-1] * size
        self.rank = [0] * size

    def Find(self, x: int) -> int:
        if self.root[x] == -1:
            return x
        self.root[x] = self.Find(self.root[x])
        return self.root[x]

    def Union(self, x: int, y: int) -> bool:
        px = self.Find(x)
        py = self.Find(y)
        if px == py:
            return False
        if self.rank[px] < self.rank[py]:
            self.root[px] = py
            self.rank[py] += self.rank[px]
        else:
            self.root[py] = px
            self.rank[px] += self.rank[py]
        return True
