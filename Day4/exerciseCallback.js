function fetchUserData(username, callback) {
    // URL API GitHub
    const apiUrl = `https://api.github.com/users/${username}`;
  
    // Lakukan HTTP GET request menggunakan fetch
    fetch(apiUrl)
      .then(response => {
        // Periksa apakah response OK (status code 200)
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        // Parse response JSON
        return response.json();
      })
      .then(data => {
        // Panggil callback dengan data yang diambil dari API
        callback(null, data);
      })
      .catch(error => {
        // Panggil callback dengan error jika terjadi masalah
        callback(error, null);
      });
  }
  
  // Contoh penggunaan:
  const username = "rezazulf";
  fetchUserData(username, (error, data) => {
    if (error) {
      console.error(`Error: ${error.message}`);
    } else {
      console.log("Data dari GitHub:", data);
      // Lakukan operasi lain dengan data yang diperoleh
    }
  });
  