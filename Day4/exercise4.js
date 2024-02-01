function oddEven(angka) {
    if (typeof angka !== 'number' || isNaN(angka)) {
        console.log('invalid data!')
    }
    else if (angka % 2 == 0) {
        console.log("Genap");
    }
    else {
        console.log("Ganjil");
    }
}

oddEven(3)
oddEven(4)
oddEven(5)
oddEven(6)
oddEven('a')
oddEven(10)