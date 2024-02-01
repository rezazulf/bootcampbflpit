function compareNumbers(firstNumber, secondNumber) {
    if (firstNumber < secondNumber) {
        return true;
    } else if (firstNumber > secondNumber) {
        return false;
    } else {
        return -1;
    }
}

console.log(compareNumbers(5, 10));
console.log(compareNumbers(15, 5));
console.log(compareNumbers(8, 8));
