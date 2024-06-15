import React, { useState } from 'react';
import axios from 'axios';

const SubscriptionForm = () => {
    const [email, setEmail] = useState('');
    const [plan, setPlan] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.post('/api/subscriptions', { email, plan });
            alert('Subscription created!');
        } catch (err) {
            console.error(err);
            alert('Error creating subscription');
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input type='email' value={email} onChange={(e) => setEmail(e.target.value)} placeholder='Email' required />
            <select value={plan} onChange={(e) => setPlan(e.target.value)}>
                <option value='basic'>Basic</option>
                <option value='premium'>Premium</option>
            </select>
            <button type='submit'>Subscribe</button>
        </form>
    );
};

export default SubscriptionForm;