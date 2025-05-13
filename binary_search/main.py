def binary_search(list_items, item):
	down = 0
	high = len(list_items) - 1

	while down <= high:
		middle = (down + high) // 2
		shot = list_items[middle]
		if shot == item:
			return middle
		if shot > item:
			high = middle - 1
		else:
			down = middle + 1
	return None

my_list = [1,3,5,7,9,11,13,15,17,19,21,23,25,27,30]
list_128 = list(range(1,129))
list_256 = list(range(1,257))

print(binary_search(my_list, 23)) # => 11
print(binary_search(my_list, -1)) # => None
print(binary_search(list_128, 96)) # => show index 95
print(binary_search(list_256, 255)) # => show indix 254
