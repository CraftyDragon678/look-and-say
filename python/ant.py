n = int(input())

s = '1'

for i in range(n):
  ch = s[0]
  count = 1
  next = ''
  for j in range(1, len(s)):
    if s[j] == ch:
      count += 1
    else:
      next += f'{ch}{count}'
      ch = s[j]
      count = 1
  next += f'{ch}{count}'
  s = next

print(s)