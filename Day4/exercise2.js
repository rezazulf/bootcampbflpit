let biodata = {
    nama_depan: 'Ejak',
    nama_belakang: 'Z',
    hobi: ['traveling', 'reading', 'gaming'],
    angka_favorit: 5,
    memakai_kacamata: true,
};

console.log('Biodata awal:', biodata);

console.log('Nama lengkap:', biodata.nama_depan + ' ' + biodata.nama_belakang);

biodata.angka_favorit = 8;
console.log('Angka favorit setelah perubahan:', biodata.angka_favorit);

biodata.hobi.push('sleeping');
console.log('Hobi setelah penambahan:', biodata.hobi);

biodata.lulusan = 'Hacktiv8';
console.log('Biodata setelah penambahan property lulusan:', biodata);

console.log('List hobi:');
for (let i = 0; i < biodata.hobi.length; i++) {
    console.log('-', biodata.hobi[i]);
}

console.log("====================================================");

console.log('List keys:');
for (let key in biodata) {
    console.log('-', key);
}
console.log("====================================================");

console.log('List values:');
for (let key in biodata) {
    console.log('-', biodata[key]);
}

console.log("====================================================");

console.log('Property objek:');
for (let key in biodata) {
    console.log(`${key} : ${biodata[key]}`);
}