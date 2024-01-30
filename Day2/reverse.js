// function reverse(str) {
//     let words = str.split(" ");
//     let reversedWords = [];
//     for (let i = 0; i < words.length; i++) {
//         let reversedWord = "";
//         for (let j = words[i].length - 1; j >= 0; j--) {
//             reversedWord += words[i][j];
//         }
//         reversedWords.push(reversedWord);
//     }
//     let reversedString = "";
//     for (let i = 0; i < reversedWords.length; i++) {
//         reversedString += reversedWords[i];
//         if (i < reversedWords.length - 1) {
//             reversedString += " ";
//         }
//     }
//     return reversedString;
// }

// console.log(reverse("HELLO"));

let str = 'wadidaw';

let newString = '';
for (let i = str.length - 1; i >= 0; i--) {
    newString += str[i];
}

// function isPalindrome(s) {
//     return s === s.split("").reverse().join("") ? "Palindrom" : "Not Palindrom";
// }

if (str == newString){
    console.log("Kata Palindrom");
}else{
    console.log("Bukan Palindrom")
}

console.log(newString);
// console.log(isPalindrome(str));