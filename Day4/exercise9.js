
function checkDistance(num) {
    for (let i = 0; i < num.length; i++) {
        if (num[i] === "a" && num[i + 4] === "b") {
            return true;
        } else if (num[i] === "b" && num[i + 4] === "a") {
            return true;
        }
    }
    return false;
}

// TEST CASES
console.log(checkDistance('lane borrowed')); // true
console.log(checkDistance('i am sick')); // false
console.log(checkDistance('you are boring')); // true
console.log(checkDistance('barbarian')); // true
console.log(checkDistance('bacon and meat')); // false
