function perkalianAngka(arr) {
    if (arr.length < 3) {
        return "Angka harus lebih dari 3";
    }
    // Menghapus duplikat
    const uniqueArr = Array.from(new Set(arr));
    // Mengurutkan secara menurun
    const sortedArr = uniqueArr.sort((a, b) => b - a);
    // Mengambil tiga angka terbesar
    const threeLargest = sortedArr.slice(0, 3);
    // Mengalikan tiga angka terbesar
    const result = threeLargest.reduce((acc, num) => acc * num, 1);
    return result;
}
let angka = [2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 10];

try {
    const result = perkalianAngka(angka);
    console.log(result);
} catch (error) {
    console.error(error.message);
}
