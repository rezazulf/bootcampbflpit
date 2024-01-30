var umur = [30, 32, 24, 26, 19, 17, 81];

// function sort(arr) {
//     let panjangArray = arr.length;
//     for (let i = 0; i < panjangArray; i++) {
//         for (let j = 0; j < panjangArray - i - 1; j++) {
//             if (arr[j] > arr[j + 1]) {
//                 let temp = arr[j];
//                 arr[j] = arr[j + 1];
//                 arr[j + 1] = temp;
//             }
//         }
//     }
//     return arr;
// }

// console.log(umur);

function sorting(arr) {
    let len = arr.length;
    for (let i = 0; i < len; i++) {
        for (let j = 0; j < len - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                let temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
    for (let i = 0; i < len; i++) {
        console.log(arr[i]);
    }
}

sorting(umur);
console.log("\n");


var Angka = [1, 7, 2, 8, 3, 4, 5, 0, 9];

for (let i = 0; i < Angka.length; i++) {

    for (j = 0; j < i; j++) {
        if (Angka[i] < Angka[j]) {
            var x = Angka[i];
            Angka[i] = Angka[j];
            Angka[j] = x;
        }
    }
}

for (let i = 0; i < Angka.length; i++) {
    console.log(Angka[i]);
}




