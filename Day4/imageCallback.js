function loadImage(imageUrl, successCallback, errorCallback) {
    // Membuat elemen img baru
    const imgElement = document.createElement('img');

    // Menetapkan atribut src dengan URL yang diberikan
    imgElement.src = imageUrl;

    // Menangani kejadian saat gambar berhasil dimuat
    imgElement.onload = function () {
        // Memanggil callback sukses dengan elemen gambar
        successCallback(imgElement);
    };

    // Menangani kejadian saat terjadi kesalahan pada gambar
    imgElement.onerror = function () {
        // Memanggil callback error jika gambar gagal dimuat
        errorCallback('Failed to load image');
    };
}

// Contoh pemanggilan fungsi loadImage dengan callback
loadImage(
    'https://avatars.githubusercontent.com/u/70256496?v=4',
    function (imgElement) {
        console.log('Image loaded successfully:', imgElement);
    },
    function (error) {
        console.error('Error:', error);
    }
);
