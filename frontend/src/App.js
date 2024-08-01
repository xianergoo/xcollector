import React, { useState, useEffect } from 'react';
import { fetchDevices, createDevice, deleteDevice } from './services/apiService'; // 导入API函数
import DeviceList from './components/DeviceList';
import DeviceForm from './components/DeviceForm';

function App() {
    const [devices, setDevices] = useState([]);

    useEffect(() => {
      handleGetDevices(); // 调用API函数
    }, []);

    const handleGetDevices = async () => {
        try {
            const response = await fetchDevices(); // 使用导入的API函数
            setDevices(response.data);
        } catch (error) {
            console.error('Error fetching devices:', error);
        }
    };

    const handleCreateDevice = async (newDevice) => {
        try {
            await createDevice(newDevice); // 使用导入的API函数
            fetchDevices(); // 重新获取设备列表
        } catch (error) {
            console.error('Error creating device:', error);
        }
    };

    const handleDeleteDevice = async (deviceId) => {
        try {
            await deleteDevice(deviceId); // 使用导入的API函数
            fetchDevices(); // 重新获取设备列表
        } catch (error) {
            console.error('Error deleting device:', error);
        }
    };

    return (
        <div className="App">
            <h1>Device Management</h1>
            <DeviceForm onCreateDevice={handleCreateDevice} />
            <DeviceList devices={devices} onDeleteDevice={handleDeleteDevice} />
        </div>
    );
}

export default App;