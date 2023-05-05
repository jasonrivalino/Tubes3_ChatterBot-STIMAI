import axios from 'axios';

const API_URL = 'http://localhost:8080';

// Fungsi untuk mengambil list data riwayat dari API link "http://localhost:8080/histories/:id"
export function getHistories(id) {
    return axios.get(`${API_URL}/histories/${id}`).then((res) => res.data);
    }

    