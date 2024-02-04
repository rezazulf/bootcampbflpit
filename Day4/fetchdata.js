function fetchUserData(username) {
    return new Promise((resolve, reject) => {
        const url = `https://api.github.com/users/${username}`;
        
        // Gunakan Fetch API untuk melakukan HTTP GET request
        fetch(url)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                // Panggil callback dengan data yang diambil dari API
                resolve(data);
            })
            .catch(error => {
                // Tangani error jika terjadi
                reject(error);
            });
    });
}

// Contoh penggunaan
const username = 'rezazulf';

fetchUserData(username)
    .then(data => {
        // Callback: Lakukan sesuatu dengan data yang diambil dari API
        console.log(data);
    })
    .catch(error => {
        // Callback: Tangani error jika terjadi
        console.error(error);
    });
