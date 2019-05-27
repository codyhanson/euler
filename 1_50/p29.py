# coding=utf-8
"""
Distinct powers
Problem 29 
Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:
2^2=4, 23=8, 24=16, 25=32
3^2=9, 33=27, 34=81, 35=243
4^2=16, 43=64, 44=256, 45=1024
5^2=25, 53=125, 54=625, 55=3125
If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:
4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125
How many distinct terms are in the sequence generated by ab for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
"""

terms = {}

for a in range(2, 101):
    for b in range(2, 101):
        term = a**b
        terms[term] = True

print(len(terms.keys()))
