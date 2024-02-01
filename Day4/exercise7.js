function sortLettersAscending(inputText) {
    return inputText.split('').sort().join('');
}

const originalText = 'ini contoh kata!!';
const sortedText = sortLettersAscending(originalText);
console.log('Original Text:', originalText);
console.log('Sorted Text:', sortedText);
