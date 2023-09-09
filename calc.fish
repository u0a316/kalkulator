#!/usr/bin/env fish

# Main program
echo "Fish Calculator"

# Pastikan ada dua argumen yang diberikan
if count $argv != 2
    echo "Usage: ./calc.fish <number1> <number2>"
    exit 1
end

set num1 $argv[1]
set num2 $argv[2]

# Hitung hasil penambahan dan perkalian
set sum (math $num1 + $num2)
set product (math $num1 * $num2)

echo "Result of addition: $sum"
echo "Result of multiplication: $product"

