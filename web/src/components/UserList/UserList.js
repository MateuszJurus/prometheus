import React, { useState, useEffect } from 'react';
import './UserList.css';

const UserList = () => {

    const [users, setUser] = useState([])

    useEffect(() => {
        async function getUsers() {
            console.log('Attempting to fetch users...');
            try {
                const response = await fetch('http://localhost:8080/user/list');
        
                if (response.ok) {
                    const data = await response.json();
                    console.log('Users list fetched successfuly!');
                    setUser(data)
                } else {
                    console.error('Failed to fetch users');
                }
            }
            catch (error) {
                console.error('Error:', error);
            }
        }
        getUsers()
    }, [])

    // Render user list
    const userList = users.map((user, index) => (
        <tr key={index}>
            <td>{user.ID}</td> 
            <td>{user.username}</td> 
            <td>{user.email}</td> 
            <td>{user.role}</td> 
        </tr>
    ));

    return (
        <table>
            <thead>
                <tr>
                    <td>ID</td>
                    <td>Username</td>
                    <td>Email</td>
                    <td>Role</td>
                </tr>
            </thead>
            <tbody>
            { userList }
            </tbody>
        </table>
    );
}

export default UserList;