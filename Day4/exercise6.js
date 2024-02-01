function reverseString(inputString) {
    return inputString.split('').reduce((reversed, char) => char + reversed, '');
}

const originalString = 'Hello, World!';
const reversedString = reverseString(originalString);
console.log('Original String:', originalString);
console.log('Reversed String:', reversedString);
