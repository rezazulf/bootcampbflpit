let userProfile = {
    nama: 'Muhammad Reza Zulfikarsyah',
    role: 'Tukang Tidur',
    asal: 'Surabaya',
    availability: 'Full Time',
    email: 'rezazulfff@gmail.com',
    phone: 628888888888,
    age: 23,
    address: 'Jl. Jalan 1',
};

function editData() {
    let nama = document.getElementById("inputNama").value;
    let role = document.getElementById("inputRole").value;
    let asal = document.getElementById("inputAsal").value;
    let availability = document.getElementById("inputAvailability").value;
    let email = document.getElementById("inputEmail").value;
    let phone = document.getElementById("inputPhone").value;
    let age = document.getElementById("inputAge").value;
    let address = document.getElementById("inputAddress").value;

    document.getElementById("nama").innerHTML = nama;
    document.getElementById("role").innerHTML = role;
    document.getElementById("asal").innerHTML = asal;
    document.getElementById("availability").innerHTML = availability;
    document.getElementById("email").innerHTML = email;
    document.getElementById("phone").innerHTML = phone;
    document.getElementById("age").innerHTML = age;
    document.getElementById("address").innerHTML = address;
    console.log(nama, role, asal, availability, email, phone, age, address);
}

const form = document.getElementById("form");
form.addEventListener("submit", (event) => {
    $('#form').modal('toggle'); //or  $('#IDModal').modal('hide');
    event.preventDefault();
    editData()
});