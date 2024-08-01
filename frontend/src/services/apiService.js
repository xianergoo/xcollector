import axios from 'axios';

const apiUrl = 'http://localhost:8080'; // 替换为您的API服务地址

export function fetchDevices() {
    return axios.get(`${apiUrl}/devices`);
}

export function createDevice(device) {
    console.log(`${apiUrl}/devices`, device);
    return axios.post(`${apiUrl}/devices`, device);
}

export function deleteDevice(id) {
    return axios.delete(`${apiUrl}/devices/${id}`);
}