# n = gets.chomp.to_i
# arr = gets.chomp.split(" ").map(&:to_i)
A,B,C,D,E,F = gets.chomp.split(" ").map(&:to_i)
i = 998244353
a = (A*B*C) - (D*E*F)
puts a%i
