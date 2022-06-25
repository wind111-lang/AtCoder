N, K, Q = list(map(int, input().split()))

A = list(map(int, input().split()))
L = list(map(int, input().split()))

List = []

for i in A:
    for j in L:
        #print(i,j)
        if j == N or List.count(j-1) > 0:
            break
        elif j-1 == A.index(i):
            List.append(i+1)
            break

for k in List:
  print(k)
