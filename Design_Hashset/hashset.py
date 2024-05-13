class MyHashSet:

    def __init__(self):
        self.payload = [False] * 1000001

    def add(self, key: int) -> None:
        self.payload[key] = True

    def remove(self, key: int) -> None:
        self.payload[key] = False

    def contains(self, key: int) -> bool:
        return self.payload[key]


# Your MyHashSet object will be instantiated and called as such:
# obj = MyHashSet()
# obj.add(key)
# obj.remove(key)
# param_3 = obj.contains(key)
