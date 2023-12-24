import { useState } from 'react';
import axios from 'axios';

function PostForm() {
    const [formData, setFormData] = useState({
        title: '',
        body: '',
        authorName: '',
        authorEmail: '',
    });

    const handleInputChange = (e: { target: { name: any; value: any; }; }) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };

    const handleSubmit = (e: { preventDefault: () => void; }) => {
        e.preventDefault();

        const postData = {
            title: formData.title,
            body: formData.body,
            author: {
                name: formData.authorName,
                email: formData.authorEmail,
            },
        };

        axios.post('your_backend_endpoint', postData)
            .then((response) => {
                console.log('Data posted successfully:', response.data);
                setFormData({
                    title: '',
                    body: '',
                    authorName: '',
                    authorEmail: '',
                });
            })
            .catch((error) => {
                console.error('Error posting data:', error);
            });
    };

    return (
        <form className='bg-sky-700 flex p-5 rounded-md flex-col align-center gap-3 w-[35rem]' onSubmit={handleSubmit}>
            <div className='flex flex-col'>
                <label htmlFor="title">Title:</label>
                <input
                    type="text"
                    id="title"
                    name="title"
                    value={formData.title}
                    onChange={handleInputChange}
                />
            </div>
            <div className='flex flex-col'>
                <label htmlFor="body">Body:</label>
                <textarea
                    id="body"
                    name="body"
                    value={formData.body}
                    onChange={handleInputChange}
                ></textarea>
            </div>
            <div className='flex gap-3 px-3'>
                <div className='flex flex-col w-1/2'>
                    <label htmlFor="authorName">Name:</label>
                    <input
                        type="text"
                        id="authorName"
                        name="authorName"
                        value={formData.authorName}
                        onChange={handleInputChange}
                    />
                </div>
                <div className='flex flex-col w-1/2'>
                    <label htmlFor="authorEmail">Email:</label>
                    <input
                        type="email"
                        id="authorEmail"
                        name="authorEmail"
                        value={formData.authorEmail}
                        onChange={handleInputChange}
                    />
                </div>
            </div>
            <button type="submit" className='bg-emerald-500'>Submit</button>
        </form>
    );
};

export default PostForm;
