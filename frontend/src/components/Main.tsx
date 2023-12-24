import axios from 'axios'
import { useEffect, useState } from 'react'
import WordCard from './WordCard'
import Form from './PostForm'
interface Word {
    id: number,
    title: string,
    body: string,
    author: Author,
}
interface Author {
    id: number,
    name: string,
    email: string,
}
function Main() {
    const [data, setData] = useState<Word[]>()
    useEffect(() => {
        axios.get("http://localhost:8080/api/words/").then((res) => {
            setData(res.data)
        })
    }, [])
    return (
        <main >
            <Form />
            <div className='flex gap-20 p-20'>
                {/* {data && data.map((word) => (
                    <WordCard key={word.id} id={word.id} title={word.title} body={word.body} author={word.author} />
                ))} */}

                <WordCard id={1} title='Breakpoints and media queries' body='You can also use variant modifiers to target media queries like responsive breakpoints, dark mode, prefers-reduced-motion, and more. For example, use md:font-bold to apply the font-bold utility at only medium screen sizes and above' author={{ id: 1, name: "Hundera", email: "test@test.com" }} />
                <WordCard id={1} title='Breakpoints and media queries' body='You can also use variant modifiers to target media queries like responsive breakpoints, dark mode, prefers-reduced-motion, and more. For example, use md:font-bold to apply the font-bold utility at only medium screen sizes and above' author={{ id: 1, name: "Hundera", email: "test@test.com" }} />
                <WordCard id={1} title='Breakpoints and media queries' body='You can also use variant modifiers to target media queries like responsive breakpoints, dark mode, prefers-reduced-motion, and more. For example, use md:font-bold to apply the font-bold utility at only medium screen sizes and above' author={{ id: 1, name: "Hundera", email: "test@test.com" }} />
            </div>
        </main>
    )
}

export default Main