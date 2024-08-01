import React, { useState, useEffect } from 'react';
import { fetchDevices, createDevice, deleteDevice } from './services/apiService'; // 导入API函数
import DeviceList from './components/DeviceList';
import DeviceForm from './components/DeviceForm';

function App() {
    const [devices, setDevices] = useState([]);

    useEffect(() => {
        getDevices(); // 调用API函数
    }, []);

    const getDevices = async () => {
        try {
            const response = await fetchDevices();
            if (Array.isArray(response.data)) {
                setDevices(response.data);
            } else {
                console.error('Invalid response format:', response);
            }
        } catch (error) {
            console.error('Error fetching devices:', error);
            // 如果需要，可以设置一个默认的设备列表或错误状态
            setDevices([]);
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