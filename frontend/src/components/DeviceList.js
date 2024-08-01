import React from 'react';

function DeviceList({ devices, onDeleteDevice }) {
    return (
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>IP Address</th>
                    <th>Group</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                {devices.map((device) => (
                    <tr key={device.id}>
                        <td>{device.id}</td>
                        <td>{device.ip_address}</td>
                        <td>{device.group}</td>
                        <td>
                            <button onClick={() => onDeleteDevice(device.id)}>Delete</button>
                        </td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
}

export default DeviceList;