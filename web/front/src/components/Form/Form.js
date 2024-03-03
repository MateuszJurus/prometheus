import React, { useState } from 'react';

const Form = () => {
    const [name, setName] = useState('');
    const [result, setResult] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();
      
        try {
          const response = await fetch('http://localhost:8080/name', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name }),
          });
      
          if (response.ok) {
            const data = await response.json();
            setResult(data);
          } else {
            console.error('Error submitting name');
          }
        } catch (error) {
          console.error('Error:', error);
        }
      };

    const handleNameChange = (event) => {
        setName(event.target.value);
    };

    return(
        <div>
            <form className="form" onSubmit={handleSubmit}>
                <input
                    type="text"
                    name="name"
                    value={name}
                    onChange={handleNameChange}
                    placeholder="Enter your name"
                />
                <input type="submit" value="Submit" />
            </form>
            <div>
                {result}
            </div>
        </div>
    )
}

export default Form;