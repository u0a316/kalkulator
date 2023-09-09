<?php
// Main program
echo "PHP Calculator\n";

// Pastikan ada dua argumen yang diberikan
if ($argc != 3) {
    echo "Usage: php calc.php <number1> <number2>\n";
    exit(1);
}

// Simpan angka dari argumen
$num1 = (float)$argv[1];
$num2 = (float)$argv[2];

// Hitung hasil penambahan dan perkalian
$sum = $num1 + $num2;
$product = $num1 * $num2;

// Tampilkan hasil
echo "Result of addition: $num1 + $num2 = $sum\n";
echo "Result of multiplication: $num1 * $num2 = $product\n";
?>

