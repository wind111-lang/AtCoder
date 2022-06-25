from itertools import count
import string

N, X = list(map(int, input().split()))

omoji = [chr(i) for i in range(65, 91)]

A = []

for i in omoji:
    for j in range(N):
        A += i

print(A[X-1])
