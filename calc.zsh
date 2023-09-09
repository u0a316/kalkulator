#!/usr/bin/env zsh

function add {
    local sum=$(( $1 + $2 ))
    echo "Result of addition: $sum"
}

function multiply {
    local product=$(( $1 * $2 ))
    echo "Result of multiplication: $product"
}

# Main program
echo "Zsh Calculator"

# Pastikan ada dua argumen yang diberikan
if (( $# != 2 )); then
    echo "Usage: ./calc.zsh <number1> <number2>"
    exit 1
fi

# Panggil fungsi add dan multiply
add $1 $2
multiply $1 $2

