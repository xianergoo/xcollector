import React, { useState } from 'react';

function DeviceForm({ onCreateDevice }) {
    const [ipAddress, setIPAddress] = useState('');
    const [group, setGroup] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();
        const newDevice = { ip_address: ipAddress, group: parseInt(group, 10) };
        await onCreateDevice(newDevice);
        setIPAddress('');
        setGroup('');
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                IP Address:
                <input type="text" value={ipAddress} onChange={(e) => setIPAddress(e.target.value)} />
            </label>
            <label>
                Group:
                <input type="number"  value={group} onChange={(e) => setGroup(e.target.value)} />
            </label>
            <button type="submit">Add Device</button>
        </form>
    );
}

export default DeviceForm;