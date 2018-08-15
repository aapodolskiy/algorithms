print('input a sequence of positive integer numbers divided by spaces:')

str_arr = input().strip().split(' ')
arr = [int(num) for num in str_arr]

print('here is your sequence:')
print(arr)

def insertionSort(a):
  for j in range(1, len(a)):
    k = a[j]
    i = j - 1
    while a[i] > k and i>=0:
      a[i+1] = a[i]
      i -= 1
    a[i+1] = k
  return a

arr = insertionSort(arr)

print('here is your sequence sorted:')

print(arr)

