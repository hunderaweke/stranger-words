
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
function WordCard(word: Word) {
    return (
        <>
            <div className="bg-cyan-700 text-white w-80 p-10 rounded-lg">
                <h1 className="font-sans font-bold text-xl">{word.title}</h1>
                <div className="font-[300] text-sm flex flex-col gap-3">
                    <p>{word.body}</p>
                    <p className="font-[400] text-gray-300 opacity-[0.8]">Author:{word.author.name}</p>
                    <p className="font-[400] text-gray-300 opacity-[0.8]">Email: {word.author.email}</p>
                </div>
            </div>
        </>
    )
}

export default WordCard