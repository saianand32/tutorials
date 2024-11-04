class HashMapFinal:
    class Entity:
        # Inner class to represent key-value pairs
        def __init__(self, key, value):
            self.key = key
            self.value = value

    def __init__(self):
        # Initialize the hash map with a list of empty lists and set initial size and load factor
        self.list = [[] for _ in range(10)]
        self.size = 0
        self.lf = 0.5

    def put(self, key, value):
        # Calculate the hash index for the key
        hash_index = abs(hash(key) % len(self.list))
        entities = self.list[hash_index]

        # Check if the key already exists; if so, update its value
        for entity in entities:
            if entity.key == key:
                entity.value = value
                return

        # Check if rehashing is needed
        if self.size / len(self.list) > self.lf:
            self.rehash()

        # Add new key-value pair
        entities.append(self.Entity(key, value))
        self.size += 1

    def rehash(self):
        # Double the size of the hash map and rehash all existing entries
        print("We are now rehashing!")
        old = self.list
        self.list = [[] for _ in range(len(old) * 2)]
        self.size = 0

        # Reinsert all old entries into the new list
        for entities in old:
            for entity in entities:
                self.put(entity.key, entity.value)

    def get(self, key):
        # Retrieve the value associated with the key
        hash_index = abs(hash(key) % len(self.list))
        entities = self.list[hash_index]
        for entity in entities:
            if entity.key == key:
                return entity.value
        return None

    def remove(self, key):
        # Remove the entry associated with the key
        hash_index = abs(hash(key) % len(self.list))
        entities = self.list[hash_index]
        target = None

        # Find the target entity to remove
        for entity in entities:
            if entity.key == key:
                target = entity
                break

        # Remove the entity if found
        if target:
            entities.remove(target)
            self.size -= 1

    def contains_key(self, key):
        # Check if the key exists in the map
        return self.get(key) is not None

    def __str__(self):
        # Create a string representation of the hash map
        builder = []
        builder.append("{")
        for entities in self.list:
            for entity in entities:
                builder.append(f"{entity.key} = {entity.value} , ")
        builder.append("}")
        return ''.join(builder)

# Example usage
hash_map = HashMapFinal()
hash_map.put("key1", "value1")
hash_map.put("key2", "value2")
print(hash_map)  # Display the current state of the hash map
print(hash_map.get("key1"))  # Retrieve value for 'key1'
hash_map.remove("key1")  # Remove entry for 'key1'
print(hash_map)  # Display the state after removal
