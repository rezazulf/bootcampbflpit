function bilPrimaAntara(angkaPertama, angkaKedua) {
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

    let angkaPrimaArray = [];

    for (let i = angkaPertama; i <= angkaKedua; i++) {
        if (bilPrima(i)) {
            angkaPrimaArray.push(i);
        }
    }

    return angkaPrimaArray;
}

console.log(bilPrimaAntara(1, 5));
console.log(bilPrimaAntara(5, 10));
console.log(bilPrimaAntara(10, 20));