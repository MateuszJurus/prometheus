import React, { useState, useEffect } from 'react';

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
                    console.error('Failed to add user');
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
            {/* Example: Assuming user object has `username` */}
            <td>{user.username}</td> 
        </tr>
    ));

    return (
        <table>
            { userList }
        </table>
    );
}

export default UserList;