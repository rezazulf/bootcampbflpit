var angka = [10, 9, 10, 9, 8, 5, 5, 9, 6, 6, 'a', 'c', 'd', 'e', 'c', 'd', 'f', 'c', 'c'];

function count(arr) {
    var n = arr.length;
    var freq = [];
    for (var i = 0; i < n; i++) {
        freq.push(0);
    }
    for (var i = 0; i < n; i++) {
        for (var j = 0; j < n; j++) {
            if (arr[i] === arr[j]) {
                freq[i]++;

            }
        }
    }
    var angkaTemp = 0;
    var angkasTemp = 0;
    for (i = 0; i < n; i++) {
        if (freq[i] > angkaTemp) {
            angkasTemp = i;
            angkaTemp = freq[i];
        }
    }
    // return (angka[angkasTemp],angkaTemp);
    console.log(arr[angkasTemp], angkaTemp);
}

count(angka);