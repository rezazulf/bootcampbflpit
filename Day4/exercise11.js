function bilPrima(angka) {
    if (angka <= 1) {
        return false;
    }
    for (let i = 2; i <= Math.sqrt(angka); i++) {
        if (angka % i === 0) {
            return false;
        }
    }
    return true;
}

console.log(bilPrima(3));   // Output: true
console.log(bilPrima(7));   // Output: true
console.log(bilPrima(6));   // Output: false
console.log(bilPrima(23));  // Output: true
console.log(bilPrima(33));  // Output: false
